package generate

// https://fly.io/dist-sys/2/

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
	"github.com/oklog/ulid/v2"

	"fly.io/distributed-challenge/message"
)

type response struct {
	Type      message.Type `json:"type"`
	MsgID     int          `json:"msg_id"`
	InReplyTo int          `json:"in_reply_to"`
	GUID      string       `json:"id"`
}

type request struct {
	Type  message.Type `json:"type"`
	MsgID int          `json:"msg_id"`
}

func newResponse(req request, guid string) response {
	return response{
		Type:      message.GENERATE_OK,
		MsgID:     req.MsgID,
		InReplyTo: req.MsgID,
		GUID:      guid,
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
	var generateReq request
	if err := json.Unmarshal(req.Body, &generateReq); err != nil {
		return fmt.Errorf("Handle: read request: %w", err)
	}
	if generateReq.Type != message.GENERATE {
		return fmt.Errorf("Handle: illegal body: %+v", generateReq)
	}
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())
	guid, err := ulid.New(ms, entropy)
	if err != nil {
		return fmt.Errorf("Handle: generate guid: %w", err)
	}

	return gh.node.Send(req.Src, newResponse(generateReq, guid.String()))
}
