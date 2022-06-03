package fields

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/vmihailenco/msgpack/v5"
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

func (n TimeDate) ValidateValuer() any {
	if !n.Valid {
		return nil
	}
	return n.Val
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
	return n.Val.Format(timeDateFormat)
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
	case int, int8, int16, int32, uint, uint8, uint16, uint32, uint64:
		n.Val, err = parseDateTime(asString(v), Location)
		n.Valid = (err == nil)
		return err
	}

	n.Valid = false
	return errors.New("unknown type")
}

func (n TimeDate) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return nullString, nil
	}
	return json.Marshal(n.String())
}

func (n *TimeDate) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, nullString) {
		return n.Scan(nil)
	}
	var s any
	if err := json.Unmarshal(b, &s); err != nil {
		n.Valid = false
		return err
	}
	return n.Scan(s)
}

func (n TimeDate) MarshalMsgpack() ([]byte, error) {
	if !n.Valid {
		return nullString, nil
	}
	return msgpack.Marshal(n.String())
}

func (n *TimeDate) UnmarshalMsgpack(b []byte) error {
	if bytes.Equal(b, nullString) {
		return n.Scan(nil)
	}
	var s any
	if err := msgpack.Unmarshal(b, &s); err != nil {
		n.Valid = false
		return err
	}
	return n.Scan(s)
}

var _ ValueScanner = (*TimeDate)(nil)
