package fields

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

type Float64 struct {
	Val   float64
	Valid bool
}

func (n *Float64) Scan(value interface{}) (err error) {
	if value == nil {
		n.Val, n.Valid = 0, false
		return err
	}
	err = convertAssign(&n.Val, value)
	n.Valid = err == nil
	return err
}

func (n Float64) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Val, nil
}

func (n *Float64) NilValue() *Float64 {
	if !n.Valid {
		return nil
	}
	return n
}

func (n *Float64) Ptr() *float64 {
	if !n.Valid {
		return nil
	}
	return &n.Val
}

func (n Float64) MarshalJSON() ([]byte, error) {
	if n.Valid {
		return json.Marshal(n.Val)
	}
	return nullString, nil
}

func (n *Float64) UnmarshalJSON(b []byte) error {
	// scan for null
	if bytes.Equal(b, nullString) {
		return n.Scan(nil)
	}
	var s interface{}
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	return n.Scan(s)
}

func (n Float64) MarshalMsgpack() ([]byte, error) {
	if n.Valid {
		return msgpack.Marshal(n.Val)
	}
	return nullString, nil
}

func (n *Float64) UnmarshalMsgpack(b []byte) error {
	// scan for null
	if bytes.Equal(b, nullString) {
		return n.Scan(nil)
	}
	var s interface{}
	if err := msgpack.Unmarshal(b, &s); err != nil {
		return err
	}
	return n.Scan(s)
}

var (
	_           ValueScanner = (*Float64)(nil)
	Float64Type              = reflect.TypeOf(Float64{})
)
