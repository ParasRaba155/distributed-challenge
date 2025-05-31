package counter

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/Rican7/retry"
	"github.com/Rican7/retry/backoff"
	"github.com/Rican7/retry/strategy"
	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

// https://fly.io/dist-sys/4/

const (
	counterKey        = "counter"
	createIfNotExists = true
	defaultTimeout    = 5 // 5 second
	defaultRetryLimit = 5
	defaultRetryAfter = 5 // 5 Millisecond
)

type counterHandler struct {
	node    *maelstrom.Node
	kvStore *maelstrom.KV
}

func NewCounterHandler(node *maelstrom.Node) counterHandler {
	if node == nil {
		panic("nil node in NewCounterHandler")
	}
	kvStore := maelstrom.NewSeqKV(node)
	if kvStore == nil {
		panic("nil KVStore in NewCoutnerHandler")
	}
	return counterHandler{node: node, kvStore: kvStore}
}

func (ch counterHandler) Handle(req maelstrom.Message) error {
	var addReq addRequest
	if err := json.Unmarshal(req.Body, &addReq); err != nil {
		return fmt.Errorf("counter Handle Add: read request: %w", err)
	}

	if err := addReq.validate(); err != nil {
		return fmt.Errorf("counter Handle Add: validate request: %w", err)
	}

	ctx, cancel := ctxWithTimeout()
	defer cancel()

	value, exists, err := ch.currentCounter(ctx)
	if err != nil {
		return fmt.Errorf("counter Handle Add: check for counter exists: %w", err)
	}
	if !exists {
		return ch.kvStore.Write(ctx, counterKey, addReq.Delta)
	}

	err = ch.compareAndSwapWithRetry(ctx, value, addReq.Delta)
	if err != nil {
		return fmt.Errorf("counter Handle Add: update counter:  %w", err)
	}

	return ch.node.Send(req.Src, addResponseFromReq(addReq))
}

func (ch counterHandler) HandleRead(req maelstrom.Message) error {
	var readReq readRequest
	if err := json.Unmarshal(req.Body, &readReq); err != nil {
		return fmt.Errorf("counter Handle Read: read request: %w", err)
	}

	if err := readReq.validate(); err != nil {
		return fmt.Errorf("counter Handle Read: validate request: %w", err)
	}
	ctx, cancel := ctxWithTimeout()
	defer cancel()

	value, exists, err := ch.currentCounter(ctx)
	if err != nil {
		return fmt.Errorf("counter Handle Read: check for counter exists: %w", err)
	}
	if !exists {
		return maelstrom.NewRPCError(maelstrom.KeyDoesNotExist, "key does not exists")
	}

	return ch.node.Send(req.Src, readResponseFromReq(readReq, value))
}

// currentCounter will check for the current counter value and returns it along with
// if it exists or not
//
//nolint:nonamedreturns // reason: named returns are fine, naked are not
func (ch counterHandler) currentCounter(ctx context.Context) (value int, exists bool, err error) {
	val, err := ch.kvStore.ReadInt(ctx, counterKey)
	if err == nil {
		return val, true, nil
	}

	var rpcError *maelstrom.RPCError
	if !errors.As(err, &rpcError) {
		return 0, false, fmt.Errorf("did not get RPCError: %w", err)
	}

	if rpcError.Code == maelstrom.KeyDoesNotExist {
		return 0, false, nil
	}

	return 0, false, rpcError
}

func (ch counterHandler) compareAndSwapWithRetry(ctx context.Context, value, delta int) error {
	action := func(_ uint) error {
		return ch.kvStore.CompareAndSwap(ctx, counterKey, value, value+delta, createIfNotExists)
	}
	return retry.Retry(
		action, strategy.Limit(defaultRetryLimit),
		strategy.Backoff(backoff.BinaryExponential((defaultRetryAfter * time.Millisecond))),
	)
}

func ctxWithTimeout() (context.Context, context.CancelFunc) {
	bg := context.Background()
	return context.WithTimeout(bg, time.Second*defaultTimeout)
}
