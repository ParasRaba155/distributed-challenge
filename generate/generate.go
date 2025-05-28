package generate

// https://fly.io/dist-sys/2/

import (
	"fmt"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"

	"fly.io/distributed-challenge/message"
)

type response struct {
	Type      message.Type `json:"type"`
	MsgID     int          `json:"msg_id"`
	InReplyTo int          `json:"in_reply_to"`
	ID        string       `json:"id"`
}

type request struct {
	Type  message.Type `json:"type"`
	MsgID int          `json:"msg_id"`
}

func newResponse(req request) response {
	return response{
		Type:      message.ECHO_OK,
		MsgID:     req.MsgID,
		InReplyTo: req.MsgID,
	}
}

type generateHandler struct {
	node *maelstrom.Node
}

func NewGenerateHanlder(node *maelstrom.Node) generateHandler {
	if node == nil {
		panic("nil node in NewGenerateHanlder")
	}
	return generateHandler{node: node}
}

func (gh generateHandler) Handle(req maelstrom.Message) error {
	return fmt.Errorf("not implemented")
}
