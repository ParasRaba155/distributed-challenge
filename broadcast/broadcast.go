package broadcast

// https://fly.io/dist-sys/3a/

import (
	"encoding/json"
	"fmt"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type broadcastHandler struct {
	node     *maelstrom.Node
	messages []int
}

func NewBroadcastHandler(node *maelstrom.Node) broadcastHandler {
	if node == nil {
		panic("nil node in NewBroadcastHandler")
	}
	return broadcastHandler{node: node, messages: []int{}}
}

func (bd *broadcastHandler) HandleBroadcast(req maelstrom.Message) error {
	var broadcasReq broadcastRequest
	if err := json.Unmarshal(req.Body, &broadcasReq); err != nil {
		return fmt.Errorf("HandleBroadcast: read request: %w", err)
	}

	if err := broadcasReq.validate(); err != nil {
		return fmt.Errorf("HandleBroadcast: validate request: %w", err)
	}
	bd.messages = append(bd.messages, broadcasReq.Message)
	return bd.node.Send(req.Src, brodcastResponseFromReq(broadcasReq))
}

func (bd *broadcastHandler) HandleTopology(req maelstrom.Message) error {
	var topologyReq topologyRequest
	if err := json.Unmarshal(req.Body, &topologyReq); err != nil {
		return fmt.Errorf("HandleTopology: read request: %w", err)
	}
	if err := topologyReq.validate(); err != nil {
		return fmt.Errorf("HandleTopology: validate request: %w", err)
	}

	return bd.node.Send(req.Src, brodcastTopologyResponseFromReq(topologyReq))
}

func (bd *broadcastHandler) HandleRead(req maelstrom.Message) error {
	var readReq readRequest
	if err := json.Unmarshal(req.Body, &readReq); err != nil {
		return fmt.Errorf("HandleRead: read request: %w", err)
	}
	if err := readReq.validate(); err != nil {
		return fmt.Errorf("HandleRead: validate request: %w", err)
	}

	return bd.node.Send(req.Src, brodcastReadResponseFromReq(readReq, bd.messages))
}
