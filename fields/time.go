package fields

import (
	"bytes"
	"database/sql/driver"
	"errors"
	"reflect"
	"time"
)

const (
	timeFormat   = "2006-01-02 15:04:05.000000"
	timeFormat19 = "2006-01-02 15:04:05"
)

type Time struct {
	Val   time.Time
	Valid bool // Valid is true if Time is not NULL
}

func (n Time) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Val.Format(timeFormat19), nil
}

func (n *Time) NilValue() *Time {
	if !n.Valid {
		return nil
	}
	return n
}

func (n *Time) Ptr() *time.Time {
	if !n.Valid {
		return nil
	}
	return &n.Val
}

func (n Time) String() string {
	if !n.Valid {
		return ""
	}
	return n.Val.Format(timeFormat19)
}

func (n Time) Unix() int64 {
	if !n.Valid {
		return 0
	}
	return n.Val.Unix()
}

func (n *Time) Scan(value interface{}) error {
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

func (n Time) MarshalJSON() ([]byte, error) {
	if y := n.Val.Year(); y < 0 || y >= 10000 {
		// RFC 3339 is clear that years are 4 digits exactly.
		// See golang.org/issue/4556#c15 for more discussion.
		return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(timeFormat19)+2)
	b = append(b, '"')
	b = n.Val.AppendFormat(b, timeFormat19)
	b = append(b, '"')
	return b, nil
}

func (n *Time) UnmarshalJSON(b []byte) error {
	// scan for null
	if bytes.Equal(b, nullString) {
		return n.Scan(nil)
	}

	return n.Scan(b)
}

func (n Time) MarshalMsgpack() ([]byte, error) {
	return n.MarshalJSON()
}

func (n *Time) UnmarshalMsgpack(b []byte) error {
	return n.UnmarshalJSON(b)
}

var (
	_        ValueScanner = (*Time)(nil)
	TimeType              = reflect.TypeOf(Time{})
)
