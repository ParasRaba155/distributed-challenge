package txn

import (
	"errors"
	"fmt"

	"fly.io/distributed-challenge/message"
)

var errValidatonError = errors.New("invalid")

type txnRequest struct {
	message.BaseRequest
	Txns []TxnItem `json:"txn"`
}

func (b txnRequest) validate() error {
	if b.Type != message.TXN {
		return fmt.Errorf(
			"%w: txn request should have %s type got %s",
			errValidatonError,
			message.TXN,
			b.Type,
		)
	}
	return nil
}

type txnResponse struct {
	message.BaseResponse
	Txns []TxnItem `json:"txn"`
}

func sendResponseFromReq(req txnRequest, txnx []TxnItem) txnResponse {
	return txnResponse{
		BaseResponse: message.BaseResponse{
			Type:      message.TXN_OK,
			MsgID:     req.MsgID,
			InReplyTo: req.MsgID,
		},
		Txns: txnx,
	}
}
