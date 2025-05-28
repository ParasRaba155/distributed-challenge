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
	default:
		return fmt.Sprintf("invalid(%d)", m)
	}
}

func MessageTypeFromString(input string) Type {
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
	*m = MessageTypeFromString(s)
	if *m == INVALID {
		return fmt.Errorf("received message type: %q with body: %v", s, data)
	}
	return nil
}
