package field

import (
	"database/sql/driver"
	"encoding/json"
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

func (n Float64) MarshalJSON() ([]byte, error) {
	if n.Valid {
		return json.Marshal(n.Val)
	}
	return nullString, nil
}

func (n *Float64) UnmarshalJSON(b []byte) error {
	var s interface{}
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	return n.Scan(s)
}

var (
	_ json.Marshaler   = (*Float64)(nil)
	_ json.Unmarshaler = (*Float64)(nil)
	_ driver.Valuer    = (*Float64)(nil)

	TypeFloat64 = reflect.TypeOf(Float64{})
)
