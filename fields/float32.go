package fields

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
)

type Float32 struct {
	Val   float32
	Valid bool
}

func (n *Float32) Scan(value interface{}) (err error) {
	if value == nil {
		n.Val, n.Valid = 0, false
		return err
	}
	err = convertAssign(&n.Val, value)
	n.Valid = err == nil
	return err
}

func (n Float32) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Val, nil
}

func (n Float32) ValidateValuer() any {
	if !n.Valid {
		return nil
	}
	return n.Val
}

func (n *Float32) NilValue() *Float32 {
	if !n.Valid {
		return nil
	}
	return n
}

func (n *Float32) Ptr() *float32 {
	if !n.Valid {
		return nil
	}
	return &n.Val
}

func (n Float32) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return nullString, nil
	}
	return json.Marshal(n.Val)
}

func (n *Float32) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, nullString) {
		return n.Scan(nil)
	}
	var s any
	if err := json.Unmarshal(b, &s); err != nil {
		n.Valid = false
		return err
	}
	return n.Scan(s)
}

func (n Float32) MarshalMsgpack() ([]byte, error) {
	if !n.Valid {
		return nullString, nil
	}
	return msgpack.Marshal(n.Val)
}

func (n *Float32) UnmarshalMsgpack(b []byte) error {
	if bytes.Equal(b, nullString) {
		return n.Scan(nil)
	}
	var s any
	if err := msgpack.Unmarshal(b, &s); err != nil {
		n.Valid = false
		return err
	}
	return n.Scan(s)
}

var _ ValueScanner = (*Float32)(nil)
