package fields

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

type Uint struct {
	Val   uint
	Valid bool // Valid is true if Int32 is not NULL
}

// Scan 扫描写入值 interface.
func (n *Uint) Scan(value interface{}) (err error) {
	if value == nil {
		n.Val, n.Valid = 0, false
		return err
	}
	err = convertAssign(&n.Val, value)
	n.Valid = err == nil
	return err
}

func (n Uint) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Val, nil
}

func (n *Uint) NilValue() *Uint {
	if !n.Valid {
		return nil
	}
	return n
}

func (n *Uint) Ptr() *uint {
	if !n.Valid {
		return nil
	}
	return &n.Val
}

func (n Uint) MarshalJSON() ([]byte, error) {
	if n.Valid {
		return json.Marshal(n.Val)
	}
	return nullString, nil
}

func (n *Uint) UnmarshalJSON(b []byte) error {
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

func (n Uint) MarshalMsgpack() ([]byte, error) {
	if n.Valid {
		return msgpack.Marshal(n.Val)
	}
	return nullString, nil
}

func (n *Uint) UnmarshalMsgpack(b []byte) error {
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
	_        ValueScanner = (*Uint)(nil)
	UintType              = reflect.TypeOf(Uint{})
)
