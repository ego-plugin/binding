package fields

import (
	"bytes"
	"database/sql/driver"
	"encoding/binary"
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
	bytesBuffer := bytes.NewBuffer([]byte{})
	if err := binary.Write(bytesBuffer, binary.BigEndian, n.Val); err != nil {
		return nil, err
	}
	return bytesBuffer.Bytes(), nil
}

func (n *String) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, nullString) {
		return n.Scan(nil)
	}
	return n.Scan(b)
}

func (n String) MarshalMsgpack() ([]byte, error) {
	return n.MarshalJSON()
}

func (n *String) UnmarshalMsgpack(b []byte) error {
	return n.UnmarshalJSON(b)
}

var (
	_          ValueScanner = (*String)(nil)
	StringType              = reflect.TypeOf(String{})
)
