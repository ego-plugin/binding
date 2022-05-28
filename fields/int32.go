package fields

import (
	"bytes"
	"database/sql/driver"
	"encoding/binary"
	"reflect"
)

type Int32 struct {
	Val   int32
	Valid bool // Valid is true if Int32 is not NULL
}

// Scan 扫描写入值 interface.
func (n *Int32) Scan(value interface{}) (err error) {
	if value == nil {
		n.Val, n.Valid = 0, false
		return err
	}
	err = convertAssign(&n.Val, value)
	n.Valid = err == nil
	return err
}

func (n Int32) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Val, nil
}

func (n *Int32) NilValue() *Int32 {
	if !n.Valid {
		return nil
	}
	return n
}

func (n *Int32) Ptr() *int32 {
	if !n.Valid {
		return nil
	}
	return &n.Val
}

func (n Int32) MarshalJSON() ([]byte, error) {
	bytesBuffer := bytes.NewBuffer([]byte{})
	if err := binary.Write(bytesBuffer, binary.BigEndian, n.Val); err != nil {
		return nil, err
	}
	return bytesBuffer.Bytes(), nil
}

func (n *Int32) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, nullString) {
		return n.Scan(nil)
	}
	bytesBuffer := bytes.NewBuffer(b)
	err := binary.Read(bytesBuffer, binary.BigEndian, n.Val)
	n.Valid = err == nil
	return err
}

func (n Int32) MarshalMsgpack() ([]byte, error) {
	return n.MarshalJSON()
}

func (n *Int32) UnmarshalMsgpack(b []byte) error {
	return n.UnmarshalJSON(b)
}

var (
	_         ValueScanner = (*Int32)(nil)
	Int32Type              = reflect.TypeOf(Int32{})
)
