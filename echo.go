package main

import (
	"encoding/json"
	"fmt"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type EchoResponse struct {
	Type      Message_Type `json:"type"`
	MsgID     int          `json:"msg_id"`
	InReplyTo int          `json:"in_reply_to"`
	Echo      string       `json:"echo"`
}

type EchoRequest struct {
	Type  Message_Type `json:"type"`
	MsgID int          `json:"msg_id"`
	Echo  string       `json:"echo"`
}

func NewEchoResponse(req EchoRequest) EchoResponse {
	return EchoResponse{
		Type:      ECHO_OK,
		MsgID:     req.MsgID,
		InReplyTo: req.MsgID,
		Echo:      req.Echo,
	}
}

func EchoHandler(node *maelstrom.Node) maelstrom.HandlerFunc {
	return func(req maelstrom.Message) error {
		var echoReq EchoRequest
		if err := json.Unmarshal(req.Body, &echoReq); err != nil {
			return fmt.Errorf("EchoHandler: read request: %w", err)
		}
		if echoReq.Type != ECHO {
			return fmt.Errorf("EchoHandler: illegal body: %+v", echoReq)
		}

		return node.Send(req.Src, NewEchoResponse(echoReq))
	}
}
