package fields

import (
	"bytes"
	"database/sql/driver"
	"encoding/binary"
	"reflect"
)

type Int64 struct {
	Val   int64
	Valid bool // Valid is true if Int64 is not NULL
}

// Scan 扫描写入值 interface.
func (n *Int64) Scan(value interface{}) (err error) {
	if value == nil {
		n.Val, n.Valid = 0, false
		return err
	}
	err = convertAssign(&n.Val, value)
	n.Valid = err == nil
	return err
}

func (n Int64) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Val, nil
}

func (n *Int64) NilValue() *Int64 {
	if !n.Valid {
		return nil
	}
	return n
}

func (n *Int64) Ptr() *int64 {
	if !n.Valid {
		return nil
	}
	return &n.Val
}

func (n Int64) MarshalJSON() ([]byte, error) {
	bytesBuffer := bytes.NewBuffer([]byte{})
	if err := binary.Write(bytesBuffer, binary.BigEndian, n.Val); err != nil {
		return nil, err
	}
	return bytesBuffer.Bytes(), nil
}

func (n *Int64) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, nullString) {
		return n.Scan(nil)
	}
	bytesBuffer := bytes.NewBuffer(b)
	err := binary.Read(bytesBuffer, binary.BigEndian, n.Val)
	n.Valid = err == nil
	return err
}

func (n Int64) MarshalMsgpack() ([]byte, error) {
	return n.MarshalJSON()
}

func (n *Int64) UnmarshalMsgpack(b []byte) error {
	return n.UnmarshalJSON(b)
}

var (
	_         ValueScanner = (*Int64)(nil)
	Int64Type              = reflect.TypeOf(Int64{})
)
