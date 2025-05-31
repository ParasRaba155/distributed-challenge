package counter

import (
	"errors"
	"fmt"

	"fly.io/distributed-challenge/message"
)

var errValidatonError = errors.New("invalid")

type addRequest struct {
	message.BaseRequest
	Delta int `json:"delta"`
}

func (b addRequest) validate() error {
	if b.Type != message.ADD {
		return fmt.Errorf(
			"%w: counter add request should have %s type got %s",
			errValidatonError,
			message.ADD,
			b.Type,
		)
	}
	return nil
}

type addResponse struct {
	message.BaseResponse
}

func addResponseFromReq(req addRequest) addResponse {
	return addResponse{
		BaseResponse: message.BaseResponse{
			Type:      message.ADD_OK,
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
			"%w: counter read request should have %s type got %s",
			errValidatonError,
			message.READ,
			b.Type,
		)
	}
	return nil
}

type readResponse struct {
	message.BaseResponse
	Value int `json:"value"`
}

func readResponseFromReq(req readRequest, value int) readResponse {
	return readResponse{
		BaseResponse: message.BaseResponse{
			Type:      message.READ_OK,
			MsgID:     req.MsgID,
			InReplyTo: req.MsgID,
		},
		Value: value,
	}
}
