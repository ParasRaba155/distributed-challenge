package main

import (
	"encoding/json"
	"fmt"
)

type Message_Type int

const (
	INVALID Message_Type = iota - 1
	INIT    Message_Type = iota
	ECHO
	ECHO_OK
)

func (m Message_Type) String() string {
	switch m {
	case INIT:
		return "init"
	case ECHO:
		return "echo"
	case ECHO_OK:
		return "echo_ok"
	default:
		return fmt.Sprintf("invalid(%d)", m)
	}
}

func MessageTypeFromString(input string) Message_Type {
	switch input {
	case "init":
		return INIT
	case "echo":
		return ECHO
	case "echo_ok":
		return ECHO_OK
	default:
		return INVALID
	}
}

func (m Message_Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.String())
}

func (m *Message_Type) UnmarshalJSON(data []byte) error {
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
