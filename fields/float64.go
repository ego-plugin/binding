package fields

import (
	"bytes"
	"database/sql/driver"
	"encoding/binary"
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
	bytesBuffer := bytes.NewBuffer([]byte{})
	if err := binary.Write(bytesBuffer, binary.BigEndian, n.Val); err != nil {
		return nil, err
	}
	return bytesBuffer.Bytes(), nil
}

func (n *Float64) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, nullString) {
		return n.Scan(nil)
	}
	bytesBuffer := bytes.NewBuffer(b)
	err := binary.Read(bytesBuffer, binary.BigEndian, n.Val)
	n.Valid = err == nil
	return err
}

func (n Float64) MarshalMsgpack() ([]byte, error) {
	return n.MarshalJSON()
}

func (n *Float64) UnmarshalMsgpack(b []byte) error {
	return n.UnmarshalJSON(b)
}

var (
	_           ValueScanner = (*Float64)(nil)
	Float64Type              = reflect.TypeOf(Float64{})
)
