package txn

import (
	"encoding/json"
	"errors"
	"fmt"
)

var errNilWriteOpt = errors.New("write operation must have non-nil value")

const opLength = 3

// Operation represents "r" or "w".
type Operation string

const (
	ReadOp  Operation = "r"
	WriteOp Operation = "w"
)

func (op Operation) IsValid() bool {
	return op == ReadOp || op == WriteOp
}

type TxnItem struct {
	Op    Operation
	Key   int
	Value *int // nil for "r", non-nil for "w" and response "r"
}

func (t TxnItem) MarshalJSON() ([]byte, error) {
	// Validation
	if t.Op == WriteOp && t.Value == nil {
		return nil, errNilWriteOpt
	}
	arr := make([]any, opLength)
	arr[0] = t.Op
	arr[1] = t.Key
	if t.Value != nil {
		arr[2] = *t.Value
	} else {
		arr[2] = nil
	}
	return json.Marshal(arr)
}

func (t *TxnItem) UnmarshalJSON(data []byte) error {
	var arr []json.RawMessage
	if err := json.Unmarshal(data, &arr); err != nil {
		return err
	}
	if len(arr) != opLength {
		return errors.New("txn item must be an array of 3 elements")
	}

	var opStr string
	if err := json.Unmarshal(arr[0], &opStr); err != nil {
		return fmt.Errorf("unmarhsal opStr: %w", err)
	}
	op := Operation(opStr)
	if !op.IsValid() {
		return fmt.Errorf("invalid operation: %s", opStr)
	}
	t.Op = op

	if err := json.Unmarshal(arr[1], &t.Key); err != nil {
		return fmt.Errorf("unmarhsal key: %w", err)
	}

	// Value can be null
	var tmp *int
	if string(arr[2]) != "null" {
		var val int
		if err := json.Unmarshal(arr[2], &val); err != nil {
			return err
		}
		tmp = &val
	}
	t.Value = tmp
	// Validation
	if t.Op == WriteOp && t.Value == nil {
		return errNilWriteOpt
	}
	return nil
}
