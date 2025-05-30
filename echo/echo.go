package echo

// https://fly.io/dist-sys/1/

import (
	"encoding/json"
	"fmt"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"

	"fly.io/distributed-challenge/message"
)

type echoResponse struct {
	message.BaseResponse
	Echo string `json:"echo"`
}

type echoRequest struct {
	message.BaseRequest
	Echo string `json:"echo"`
}

func newResponse(req echoRequest) echoResponse {
	return echoResponse{
		BaseResponse: message.BaseResponse{
			Type:      message.ECHO_OK,
			MsgID:     req.MsgID,
			InReplyTo: req.MsgID,
		},
		Echo: req.Echo,
	}
}

type echoHandler struct {
	node *maelstrom.Node
}

func NewEchoHandler(node *maelstrom.Node) echoHandler {
	if node == nil {
		panic("nil node in NewEchoHandler")
	}
	return echoHandler{node: node}
}

func (e echoHandler) Handle(req maelstrom.Message) error {
	var echoReq echoRequest
	if err := json.Unmarshal(req.Body, &echoReq); err != nil {
		return fmt.Errorf("Handle: read request: %w", err)
	}
	if echoReq.Type != message.ECHO {
		return fmt.Errorf("Handle: illegal body: %+v", echoReq)
	}

	return e.node.Send(req.Src, newResponse(echoReq))
}
