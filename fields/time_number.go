package fields

import (
	"bytes"
	"database/sql/driver"
	"errors"
	"reflect"
	"time"
)

type TimeNumber struct {
	Val   time.Time
	Valid bool // Valid is true if Time is not NULL
}

func (n TimeNumber) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Val.Format(timeDateFormat), nil
}

func (n *TimeNumber) NilValue() *TimeNumber {
	if !n.Valid {
		return nil
	}
	return n
}

func (n *TimeNumber) Ptr() *time.Time {
	if !n.Valid {
		return nil
	}
	return &n.Val
}

func (n TimeNumber) String() string {
	if !n.Valid {
		return ""
	}
	return n.Val.Format(timeDateFormat)
}

func (n TimeNumber) Unix() int64 {
	if !n.Valid {
		return 0
	}
	return n.Val.Unix()
}

func (n *TimeNumber) Scan(value interface{}) error {
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

func (n TimeNumber) MarshalJSON() ([]byte, error) {
	if y := n.Val.Year(); y < 0 || y >= 10000 {
		// RFC 3339 is clear that years are 4 digits exactly.
		// See golang.org/issue/4556#c15 for more discussion.
		return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(timeDateFormat)+2)
	b = append(b, '"')
	b = n.Val.AppendFormat(b, timeDateFormat)
	b = append(b, '"')
	return b, nil
}

func (n *TimeNumber) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, nullString) {
		return n.Scan(nil)
	}
	return n.Scan(b)
}

func (n TimeNumber) MarshalMsgpack() ([]byte, error) {
	return n.MarshalJSON()
}

func (n *TimeNumber) UnmarshalMsgpack(b []byte) error {
	return n.UnmarshalJSON(b)
}

var (
	_              ValueScanner = (*TimeNumber)(nil)
	TimeNumberType              = reflect.TypeOf(TimeNumber{})
)
