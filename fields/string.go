package fields

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

type String struct {
	Val   string
	Valid bool
}

// Scan 扫描写入值 interface.
func (n *String) Scan(value interface{}) (err error) {
	if value == nil {
		n.Val, n.Valid = "", false
		return err
	}
	err = convertAssign(&n.Val, value)
	n.Valid = err == nil
	return err
}

func (n *String) NilValue() *String {
	if !n.Valid {
		return nil
	}
	return n
}

func (n *String) Ptr() *string {
	if !n.Valid {
		return nil
	}
	return &n.Val
}

func (n String) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Val, nil
}

func (n String) MarshalJSON() ([]byte, error) {
	if n.Valid {
		return json.Marshal(n.Val)
	}
	return nil, nil
}

func (n *String) UnmarshalJSON(b []byte) error {
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

func (n String) MarshalMsgpack() ([]byte, error) {
	if n.Valid {
		return msgpack.Marshal(n.Val)
	}
	return nil, nil
}

func (n *String) UnmarshalMsgpack(b []byte) error {
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
	_          ValueScanner = (*String)(nil)
	StringType              = reflect.TypeOf(String{})
)
