package kafka

import (
	"encoding/json"
	"fmt"
	"maps"
	"sync"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

// https://fly.io/dist-sys/5a/

const offsetFactor = 10000

type handler struct {
	node            *maelstrom.Node
	logOffsets      map[string][]logEntry
	commitedOffsets map[string]int

	mu sync.RWMutex
}

func NewHandler(node *maelstrom.Node) *handler {
	if node == nil {
		panic("nil node in NewHandler")
	}
	return &handler{
		node:            node,
		logOffsets:      make(map[string][]logEntry),
		commitedOffsets: make(map[string]int),
	}
}

func (e *handler) Send(req maelstrom.Message) error {
	var sendReq sendRequest
	if err := json.Unmarshal(req.Body, &sendReq); err != nil {
		return fmt.Errorf("Send: read request: %w", err)
	}

	if err := sendReq.validate(); err != nil {
		return fmt.Errorf("Send: validate request: %w", err)
	}
	key := sendReq.Key
	e.mu.Lock()
	defer e.mu.Unlock()
	offsetList, ok := e.logOffsets[key]
	// Since the keys would be "1", "2", ...
	// We would initiate offset for each key with 10000, 20000,...
	if !ok {
		offset := sendReq.mustGetKeyInt() * offsetFactor
		e.logOffsets[key] = []logEntry{{Offset: offset, Value: sendReq.Msg}}
		return e.node.Send(req.Src, sendResponseFromReq(sendReq, offset))
	}
	lastoffset := offsetList[len(offsetList)-1]
	newOffset := lastoffset.Offset + 1
	offsetList = append(offsetList, logEntry{Offset: newOffset, Value: sendReq.Msg})
	e.logOffsets[key] = offsetList

	return e.node.Send(req.Src, sendResponseFromReq(sendReq, newOffset))
}

func (e *handler) Poll(req maelstrom.Message) error {
	var pollReq pollRequest
	if err := json.Unmarshal(req.Body, &pollReq); err != nil {
		return fmt.Errorf("Poll: read request: %w", err)
	}

	if err := pollReq.validate(); err != nil {
		return fmt.Errorf("Poll: validate request: %w", err)
	}

	filteredOffsets := make(map[string][]logEntry)
	e.mu.Lock()
	defer e.mu.Unlock()
	for key, offsetFrom := range pollReq.Offsets {
		logs, ok := e.logOffsets[key]
		if !ok {
			continue
		}
		filteredEntries := make([]logEntry, 0, len(logs))
		for _, log := range logs {
			if log.Offset >= offsetFrom {
				filteredEntries = append(filteredEntries, log)
			}
		}
		filteredOffsets[key] = filteredEntries
	}

	return e.node.Send(req.Src, pollResponseFromReq(pollReq, filteredOffsets))
}

func (e *handler) CommitOffset(req maelstrom.Message) error {
	var commitOffsetReq commitOffsetRequest
	if err := json.Unmarshal(req.Body, &commitOffsetReq); err != nil {
		return fmt.Errorf("CommitOffset: read request: %w", err)
	}

	if err := commitOffsetReq.validate(); err != nil {
		return fmt.Errorf("CommitOffset: validate request: %w", err)
	}

	e.mu.Lock()
	defer e.mu.Unlock()
	maps.Copy(e.commitedOffsets, commitOffsetReq.Offsets)
	return e.node.Send(req.Src, commitOffsetResponseFromReq(commitOffsetReq))
}

func (e *handler) ListCommitedOffset(req maelstrom.Message) error {
	var listCommitOffsetReq listCommitOffsetRequest
	if err := json.Unmarshal(req.Body, &listCommitOffsetReq); err != nil {
		return fmt.Errorf("ListCommitedOffset: read request: %w", err)
	}

	if err := listCommitOffsetReq.validate(); err != nil {
		return fmt.Errorf("ListCommitedOffset: validate request: %w", err)
	}

	return e.node.Send(
		req.Src,
		listCommitOffsetResponseFromReq(listCommitOffsetReq, e.commitedOffsets),
	)
}
