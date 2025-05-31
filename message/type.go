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
)

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
