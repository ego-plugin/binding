package fields

import (
	"bytes"
	"database/sql/driver"
	"encoding/binary"
	"reflect"
)

type Bool struct {
	Val   bool
	Valid bool // Valid is true if Bool is not NULL
}

func (n *Bool) Scan(value interface{}) (err error) {
	if value == nil {
		n.Val, n.Valid = false, false
		return err
	}
	err = convertAssign(&n.Val, value)
	n.Valid = err == nil
	return err
}

func (n Bool) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Val, nil
}

func (n *Bool) NilValue() *Bool {
	if !n.Valid {
		return nil
	}
	return n
}

func (n *Bool) Ptr() *bool {
	if !n.Valid {
		return nil
	}
	return &n.Val
}

func (n Bool) MarshalJSON() ([]byte, error) {
	bytesBuffer := bytes.NewBuffer([]byte{})
	if err := binary.Write(bytesBuffer, binary.BigEndian, n.Val); err != nil {
		return nil, err
	}
	return bytesBuffer.Bytes(), nil
}

func (n *Bool) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, nullString) {
		return n.Scan(nil)
	}
	bytesBuffer := bytes.NewBuffer(b)
	err := binary.Read(bytesBuffer, binary.BigEndian, &n.Val)
	n.Valid = err == nil
	return err
}

func (n Bool) MarshalMsgpack() ([]byte, error) {
	return n.MarshalJSON()
}

func (n *Bool) UnmarshalMsgpack(b []byte) error {
	return n.UnmarshalJSON(b)
}

var (
	_        ValueScanner = (*Bool)(nil)
	BoolType              = reflect.TypeOf(Bool{})
)
