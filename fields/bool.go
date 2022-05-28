package fields

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

type Bool struct {
	Val   bool
	Valid bool // Valid is true if Bool is not NULL
}

func (n *Bool) Scan(value interface{}) (err error) {
	if value == nil {
		n.Val, n.Valid = false, false
		return err
	}
	err = convertAssign(&n.Val, value)
	n.Valid = err == nil
	return err
}

func (n Bool) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Val, nil
}

func (n *Bool) NilValue() *Bool {
	if !n.Valid {
		return nil
	}
	return n
}

func (n *Bool) Ptr() *bool {
	if !n.Valid {
		return nil
	}
	return &n.Val
}

func (n Bool) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return nullString, nil
	}
	return json.Marshal(n.Val)
}

func (n *Bool) UnmarshalJSON(b []byte) error {
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

func (n Bool) MarshalMsgpack() ([]byte, error) {
	if !n.Valid {
		return nullString, nil
	}
	return msgpack.Marshal(n.Val)
}

func (n *Bool) UnmarshalMsgpack(b []byte) error {
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

var (
	_        ValueScanner = (*Bool)(nil)
	BoolType              = reflect.TypeOf(Bool{})
)
