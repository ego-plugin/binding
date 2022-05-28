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
	if n.Valid {
		return json.Marshal(n.Val)
	}
	return nullString, nil
}

func (n *Uint64) UnmarshalJSON(b []byte) error {
	// scan for null
	if bytes.Equal(b, nullString) {
		return n.Scan(nil)
	}
	if err := json.Unmarshal(b, &n.Val); err != nil {
		n.Valid = false
		return err
	}
	n.Valid = true
	return nil
}

func (n Uint64) MarshalMsgpack() ([]byte, error) {
	if n.Valid {
		return msgpack.Marshal(n.Val)
	}
	return nullString, nil
}

func (n *Uint64) UnmarshalMsgpack(b []byte) error {
	// scan for null
	if bytes.Equal(b, nullString) {
		return n.Scan(nil)
	}
	if err := msgpack.Unmarshal(b, &n.Val); err != nil {
		n.Valid = false
		return err
	}
	n.Valid = true
	return nil
}

var (
	_          ValueScanner = (*Uint64)(nil)
	Uint64Type              = reflect.TypeOf(Uint64{})
)
