package broadcast

import (
	"errors"
	"fmt"

	"fly.io/distributed-challenge/message"
)

var errValidatonError = errors.New("invalid")

type broadcastRequest struct {
	message.BaseRequest
	Message int `json:"message"`
}

func (b broadcastRequest) validate() error {
	if b.Type != message.BROADCAST {
		return fmt.Errorf(
			"%w: broadcast request should have %s type got %s",
			errValidatonError,
			message.BROADCAST,
			b.Type,
		)
	}
	return nil
}

type broadcastResponse struct {
	message.BaseResponse
}

func brodcastResponseFromReq(req broadcastRequest) broadcastResponse {
	return broadcastResponse{
		BaseResponse: message.BaseResponse{
			Type:      message.BROADCAST_OK,
			MsgID:     req.MsgID,
			InReplyTo: req.MsgID,
		},
	}
}

type readRequest struct {
	message.BaseRequest
}

func (b readRequest) validate() error {
	if b.Type != message.READ {
		return fmt.Errorf(
			"%w: broadcast read request should have %s type got %s",
			errValidatonError,
			message.READ,
			b.Type,
		)
	}
	return nil
}

type readResponse struct {
	message.BaseResponse
	Messages []int `json:"messages"`
}

func brodcastReadResponseFromReq(req readRequest, msgs []int) readResponse {
	return readResponse{
		BaseResponse: message.BaseResponse{
			Type:      message.READ_OK,
			MsgID:     req.MsgID,
			InReplyTo: req.MsgID,
		},
		Messages: msgs,
	}
}

type topologyRequest struct {
	message.BaseRequest
	Topology map[string][]string `json:"topology"`
}

func (b topologyRequest) validate() error {
	if b.Type != message.TOPOLOGY {
		return fmt.Errorf(
			"%w: broadcast topology request should have %s type got %s",
			errValidatonError,
			message.TOPOLOGY,
			b.Type,
		)
	}
	return nil
}

type topologyResponse struct {
	message.BaseResponse
}

func brodcastTopologyResponseFromReq(req topologyRequest) topologyResponse {
	return topologyResponse{
		BaseResponse: message.BaseResponse{
			Type:      message.TOPOLOGY_OK,
			MsgID:     req.MsgID,
			InReplyTo: req.MsgID,
		},
	}
}
