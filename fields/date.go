package fields

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
	"time"
)

const (
	timeFormat10 = "2006-01-02"
)

type Date struct {
	Val   time.Time
	Valid bool // Valid is true if Time is not NULL
}

func (n *Date) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Val.Format(timeFormat10), nil
}

func (n *Date) NilValue() *Date {
	if !n.Valid {
		return nil
	}
	return n
}

func (n *Date) String() string {
	if !n.Valid {
		return ""
	}
	return n.Val.Format(timeFormat10)
}

func (n *Date) Ptr() *time.Time {
	if !n.Valid {
		return nil
	}
	return &n.Val
}

func (n *Date) Scan(value interface{}) error {
	var err error

	if value == nil {
		n.Val, n.Valid = time.Time{}, false
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		// 如果时区不是上海 转为上海时区
		n.Val = v
		n.Valid = true
		return nil
	case []byte:
		n.Val, err = parseDateTime(string(v), Location)
		n.Valid = (err == nil)
		return err
	case string:
		n.Val, err = parseDateTime(v, Location)
		n.Valid = (err == nil)
		return err
	}

	n.Valid = false
	return errors.New("unknown type")
}

func (n *Date) MarshalJSON() ([]byte, error) {
	if n.Valid {
		return json.Marshal(n.Val.Format(timeFormat10))
	}
	return nullString, nil
}

func (n *Date) UnmarshalJSON(b []byte) error {
	// scan for null
	if bytes.Equal(b, nullString) {
		return n.Scan(nil)
	}
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	return n.Scan(s)
}

func (n Date) MarshalMsgpack() ([]byte, error) {
	if n.Valid {
		return msgpack.Marshal(n.Val.Format(timeFormat10))
	}
	return nullString, nil
}

func (n *Date) UnmarshalMsgpack(b []byte) error {
	// scan for null
	if bytes.Equal(b, nullString) {
		return n.Scan(nil)
	}
	var s string
	if err := msgpack.Unmarshal(b, &s); err != nil {
		return err
	}
	return n.Scan(s)
}

var (
	_        ValueScanner = (*Date)(nil)
	DateType              = reflect.TypeOf(Date{})
)
