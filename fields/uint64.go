package fields

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

type Uint64 struct {
	Val   uint64
	Valid bool // Valid is true if Int64 is not NULL
}

// Scan 扫描写入值 interface.
func (n *Uint64) Scan(value interface{}) (err error) {
	if value == nil {
		n.Val, n.Valid = 0, false
		return err
	}
	err = convertAssign(&n.Val, value)
	n.Valid = err == nil
	return err
}

func (n Uint64) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Val, nil
}

func (n *Uint64) NilValue() *Uint64 {
	if !n.Valid {
		return nil
	}
	return n
}

func (n *Uint64) Ptr() *uint64 {
	if !n.Valid {
		return nil
	}
	return &n.Val
}

func (n Uint64) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return nullString, nil
	}
	return json.Marshal(n.Val)
}

func (n *Uint64) UnmarshalJSON(b []byte) error {
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

func (n Uint64) MarshalMsgpack() ([]byte, error) {
	if !n.Valid {
		return nullString, nil
	}
	return msgpack.Marshal(n.Val)
}

func (n *Uint64) UnmarshalMsgpack(b []byte) error {
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
	_          ValueScanner = (*Uint64)(nil)
	Uint64Type              = reflect.TypeOf(Uint64{})
)
