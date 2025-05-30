package generate

// https://fly.io/dist-sys/2/

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"time"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
	"github.com/oklog/ulid/v2"

	"fly.io/distributed-challenge/message"
)

type generateResponse struct {
	message.BaseResponse
	GUID string `json:"id"`
}

type generateRequest struct {
	message.BaseRequest
}

func newResponse(req generateRequest, guid string) generateResponse {
	return generateResponse{
		BaseResponse: message.BaseResponse{
			Type:      message.GENERATE_OK,
			MsgID:     req.MsgID,
			InReplyTo: req.MsgID,
		},
		GUID: guid,
	}
}

type generateHandler struct {
	node *maelstrom.Node
}

type randV2Reader struct {
	rng *rand.Rand
}

func (r *randV2Reader) Read(p []byte) (int, error) {
	const byteLength = 256
	for i := range p {
		p[i] = byte(r.rng.IntN(byteLength))
	}
	return len(p), nil
}

func generateULID() (ulid.ULID, error) {
	//nolint:gosec // reason: int64 of time is always positive and will be < uint64 max
	seed := uint64(time.Now().UnixNano())
	//nolint:gosec // reason: acceptable for ULID generation
	rng := rand.New(rand.NewPCG(seed, seed+1))
	entropy := &randV2Reader{rng}
	ms := ulid.Timestamp(time.Now())
	return ulid.New(ms, entropy)
}

func NewGenerateHanlder(node *maelstrom.Node) generateHandler {
	if node == nil {
		panic("nil node in NewGenerateHanlder")
	}
	return generateHandler{node: node}
}

func (gh generateHandler) Handle(req maelstrom.Message) error {
	var generateReq generateRequest
	if err := json.Unmarshal(req.Body, &generateReq); err != nil {
		return fmt.Errorf("Handle: read request: %w", err)
	}
	if generateReq.Type != message.GENERATE {
		return fmt.Errorf("Handle: illegal body: %+v", generateReq)
	}
	guid, err := generateULID()
	if err != nil {
		return fmt.Errorf("Handle: generate guid: %w", err)
	}

	return gh.node.Send(req.Src, newResponse(generateReq, guid.String()))
}
