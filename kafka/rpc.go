package kafka

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"fly.io/distributed-challenge/message"
)

var errValidatonError = errors.New("invalid")

type sendRequest struct {
	message.BaseRequest
	Key string `json:"key"`
	Msg int    `json:"msg"`
}

func (b sendRequest) validate() error {
	if b.Type != message.SEND {
		return fmt.Errorf(
			"%w: kafka send request should have %s type got %s",
			errValidatonError,
			message.SEND,
			b.Type,
		)
	}
	return nil
}

// mustGetKeyInt will return the string key in int
// NOTE: This is by trial and error, but the keys is always a int in string
// This func will panic if the string is found to not to an integer.
func (b sendRequest) mustGetKeyInt() int {
	val, err := strconv.ParseInt(b.Key, 10, 64)
	if err != nil {
		panic(err)
	}
	return int(val)
}

type sendResponse struct {
	message.BaseResponse
	Offset int `json:"offset"`
}

func sendResponseFromReq(req sendRequest, offset int) sendResponse {
	return sendResponse{
		BaseResponse: message.BaseResponse{
			Type:      message.SEND_OK,
			MsgID:     req.MsgID,
			InReplyTo: req.MsgID,
		},
		Offset: offset,
	}
}

type pollRequest struct {
	message.BaseRequest
	Offsets map[string]int `json:"offsets"`
}

func (b pollRequest) validate() error {
	if b.Type != message.POLL {
		return fmt.Errorf(
			"%w: kafka poll request should have %s type got %s",
			errValidatonError,
			message.POLL,
			b.Type,
		)
	}
	return nil
}

type pollResponse struct {
	message.BaseResponse
	LogOffsets map[string][]logEntry `json:"msgs"`
}

func pollResponseFromReq(req pollRequest, logOffset map[string][]logEntry) pollResponse {
	return pollResponse{
		BaseResponse: message.BaseResponse{
			Type:      message.POLL_OK,
			MsgID:     req.MsgID,
			InReplyTo: req.MsgID,
		},
		LogOffsets: logOffset,
	}
}

type commitOffsetRequest struct {
	message.BaseResponse
	Offsets map[string]int `json:"offsets"`
}

func (b commitOffsetRequest) validate() error {
	if b.Type != message.COMMIT_OFFSETS {
		return fmt.Errorf(
			"%w: kafka commit_offsets request should have %s type got %s",
			errValidatonError,
			message.COMMIT_OFFSETS,
			b.Type,
		)
	}
	return nil
}

type commitOffsetResponse struct {
	message.BaseResponse
}

func commitOffsetResponseFromReq(req commitOffsetRequest) commitOffsetResponse {
	return commitOffsetResponse{
		BaseResponse: message.BaseResponse{
			Type:      message.COMMIT_OFFSETS_OK,
			MsgID:     req.MsgID,
			InReplyTo: req.MsgID,
		},
	}
}

type listCommitOffsetRequest struct {
	message.BaseResponse
	Keys []string `json:"keys"`
}

func (b listCommitOffsetRequest) validate() error {
	if b.Type != message.LIST_COMMITTED_OFFSETS {
		return fmt.Errorf(
			"%w: kafka list commit_offsets request should have %s type got %s",
			errValidatonError,
			message.LIST_COMMITTED_OFFSETS,
			b.Type,
		)
	}
	return nil
}

type listCommitOffsetResponse struct {
	message.BaseResponse
	Offsets map[string]int `json:"offsets"`
}

func listCommitOffsetResponseFromReq(
	req listCommitOffsetRequest,
	offsets map[string]int,
) listCommitOffsetResponse {
	return listCommitOffsetResponse{
		BaseResponse: message.BaseResponse{
			Type:      message.LIST_COMMITTED_OFFSETS_OK,
			MsgID:     req.MsgID,
			InReplyTo: req.MsgID,
		},
		Offsets: offsets,
	}
}

// logEntry is internal representation of each logEntry
//
// NOTE: custom MarshalJSON is defined to convert to appropriate response.
type logEntry struct {
	Offset int
	Value  int
}

func (l logEntry) MarshalJSON() ([]byte, error) {
	arr := [2]int{l.Offset, l.Value}
	return json.Marshal(arr)
}
