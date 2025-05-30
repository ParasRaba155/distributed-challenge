package message

// BaseResponse is the common fields on all the maelstrom response.
type BaseResponse struct {
	Type      Type `json:"type"`
	MsgID     int  `json:"msg_id"`
	InReplyTo int  `json:"in_reply_to"`
}

// BaseRequest is the common fields on all the maelstrom requests.
type BaseRequest struct {
	Type  Type `json:"type"`
	MsgID int  `json:"msg_id"`
}
