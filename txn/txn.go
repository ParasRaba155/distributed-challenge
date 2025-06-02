package txn

import (
	"encoding/json"
	"fmt"
	"sync"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

// https://fly.io/dist-sys/6a/

type handler struct {
	node *maelstrom.Node

	mu      sync.RWMutex
	kvStore map[int]int
}

func NewHandler(node *maelstrom.Node) handler {
	if node == nil {
		panic("nil node in NewHandler")
	}
	return handler{
		node:    node,
		kvStore: make(map[int]int),
	}
}

func (h *handler) Handle(req maelstrom.Message) error {
	var request txnRequest
	if err := json.Unmarshal(req.Body, &request); err != nil {
		return fmt.Errorf("Handle: read request: %w", err)
	}

	if err := request.validate(); err != nil {
		return fmt.Errorf("Handle: validate request: %w", err)
	}

	for i, txn := range request.Txns {
		if txn.Op == WriteOp {
			h.mu.Lock()
			h.kvStore[txn.Key] = *txn.Value
			h.mu.Unlock()
			continue
		}
		h.mu.RLock()
		val, exists := h.kvStore[txn.Key]
		if !exists {
			request.Txns[i] = TxnItem{
				Op:    txn.Op,
				Key:   txn.Key,
				Value: nil,
			}
			h.mu.RUnlock()
			continue
		}
		request.Txns[i] = TxnItem{
			Op:    txn.Op,
			Key:   txn.Key,
			Value: &val,
		}
		h.mu.RUnlock()
		continue
	}

	return h.node.Send(
		req.Src,
		sendResponseFromReq(request, request.Txns),
	)
}
