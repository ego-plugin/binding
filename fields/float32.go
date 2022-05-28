package fields

import (
	"bytes"
	"database/sql/driver"
	"encoding/binary"
	"reflect"
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
	bytesBuffer := bytes.NewBuffer([]byte{})
	if err := binary.Write(bytesBuffer, binary.BigEndian, n.Val); err != nil {
		return nil, err
	}
	return bytesBuffer.Bytes(), nil
}

func (n *Float32) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, nullString) {
		return n.Scan(nil)
	}
	bytesBuffer := bytes.NewBuffer(b)
	err := binary.Read(bytesBuffer, binary.BigEndian, &n.Val)
	n.Valid = err == nil
	return err
}

func (n Float32) MarshalMsgpack() ([]byte, error) {
	return n.MarshalJSON()
}

func (n *Float32) UnmarshalMsgpack(b []byte) error {
	return n.UnmarshalJSON(b)
}

var (
	_           ValueScanner = (*Float32)(nil)
	Float32Type              = reflect.TypeOf(Float32{})
)
