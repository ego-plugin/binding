package fields

import (
	"bytes"
	"database/sql/driver"
	"encoding/binary"
	"errors"
	"reflect"
	"time"
)

const (
	timeDateFormat = "2006-01-02"
)

type TimeDate struct {
	Val   time.Time
	Valid bool // Valid is true if Time is not NULL
}

func (n TimeDate) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Val.Format(timeDateFormat), nil
}

func (n *TimeDate) NilValue() *TimeDate {
	if !n.Valid {
		return nil
	}
	return n
}

func (n *TimeDate) Ptr() *time.Time {
	if !n.Valid {
		return nil
	}
	return &n.Val
}

func (n TimeDate) String() string {
	if !n.Valid {
		return ""
	}
	return n.Val.Format(timeFormat19)
}

func (n TimeDate) Unix() int64 {
	if !n.Valid {
		return 0
	}
	return n.Val.Unix()
}

func (n *TimeDate) Scan(value interface{}) error {
	var err error

	if value == nil {
		n.Val, n.Valid = time.Time{}, false
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		n.Val = v
		n.Valid = true
		return nil
	case []byte:
		n.Val, err = parseDateTime(BytesToString(v), Location)
		n.Valid = (err == nil)
		return err
	case string:
		n.Val, err = parseDateTime(v, Location)
		n.Valid = (err == nil)
		return err
	case int64:
		n.Val = time.Unix(v, 0).In(Location)
		n.Valid = (err == nil)
		return err
	}

	n.Valid = false
	return errors.New("unknown type")
}

func (n TimeDate) MarshalJSON() ([]byte, error) {
	bytesBuffer := bytes.NewBuffer([]byte{})
	if err := binary.Write(bytesBuffer, binary.BigEndian, n.Unix()); err != nil {
		return nil, err
	}
	return bytesBuffer.Bytes(), nil
}

func (n *TimeDate) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, nullString) {
		return n.Scan(nil)
	}
	return n.Scan(b)
}

func (n TimeDate) MarshalMsgpack() ([]byte, error) {
	return n.MarshalJSON()
}

func (n *TimeDate) UnmarshalMsgpack(b []byte) error {
	return n.UnmarshalJSON(b)
}

var (
	_            ValueScanner = (*TimeDate)(nil)
	TimeDateType              = reflect.TypeOf(TimeDate{})
)
