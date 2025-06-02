//nolint:cyclop // there are lots of switch case on message type, which increases complexity
package message

import (
	"encoding/json"
	"fmt"
)

type Type int

const (
	INVALID Type = iota - 1
	INIT    Type = iota
	ECHO
	ECHO_OK
	GENERATE
	GENERATE_OK
	BROADCAST
	BROADCAST_OK
	READ
	READ_OK
	TOPOLOGY
	TOPOLOGY_OK
	ADD
	ADD_OK
	SEND
	SEND_OK
	POLL
	POLL_OK
	COMMIT_OFFSETS
	COMMIT_OFFSETS_OK
	LIST_COMMITTED_OFFSETS
	LIST_COMMITTED_OFFSETS_OK
	TXN
	TXN_OK
)

//nolint:funlen // there are lots of switch case on message type
func (m Type) String() string {
	switch m {
	case INIT:
		return "init"
	case ECHO:
		return "echo"
	case ECHO_OK:
		return "echo_ok"
	case GENERATE:
		return "generate"
	case GENERATE_OK:
		return "generate_ok"
	case BROADCAST:
		return "broadcast"
	case BROADCAST_OK:
		return "broadcast_ok"
	case READ:
		return "read"
	case READ_OK:
		return "read_ok"
	case TOPOLOGY:
		return "topology"
	case TOPOLOGY_OK:
		return "topology_ok"
	case ADD:
		return "add"
	case ADD_OK:
		return "add_ok"
	case SEND:
		return "send"
	case SEND_OK:
		return "send_ok"
	case POLL:
		return "poll"
	case POLL_OK:
		return "poll_ok"
	case COMMIT_OFFSETS:
		return "commit_offsets"
	case COMMIT_OFFSETS_OK:
		return "commit_offsets_ok"
	case LIST_COMMITTED_OFFSETS:
		return "list_committed_offsets"
	case LIST_COMMITTED_OFFSETS_OK:
		return "list_committed_offsets_ok"
	case TXN:
		return "txn"
	case TXN_OK:
		return "txn_ok"
	case INVALID:
		return "invalid"
	default:
		return fmt.Sprintf("invalid(%d)", m)
	}
}

func TypeFromString(input string) Type {
	switch input {
	case "init":
		return INIT
	case "echo":
		return ECHO
	case "echo_ok":
		return ECHO_OK
	case "generate":
		return GENERATE
	case "generate_ok":
		return GENERATE_OK
	case "broadcast":
		return BROADCAST
	case "broadcast_ok":
		return BROADCAST_OK
	case "read":
		return READ
	case "read_ok":
		return READ_OK
	case "topology":
		return TOPOLOGY
	case "topology_ok":
		return TOPOLOGY_OK
	case "add":
		return ADD
	case "add_ok":
		return ADD_OK
	case "send":
		return SEND
	case "send_ok":
		return SEND_OK
	case "poll":
		return POLL
	case "poll_ok":
		return POLL_OK
	case "commit_offsets":
		return COMMIT_OFFSETS
	case "commit_offsets_ok":
		return COMMIT_OFFSETS_OK
	case "list_committed_offsets":
		return LIST_COMMITTED_OFFSETS
	case "list_committed_offsets_ok":
		return LIST_COMMITTED_OFFSETS_OK
	case "txn":
		return TXN
	case "txn_ok":
		return TXN_OK
	default:
		return INVALID
	}
}

func (m Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.String())
}

func (m *Type) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	*m = TypeFromString(s)
	if *m == INVALID {
		return fmt.Errorf("received message type: %q with body: %v", s, data)
	}
	return nil
}
