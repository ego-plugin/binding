package field

import (
	"database/sql/driver"
	"encoding/json"
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

func (n Float32) MarshalJSON() ([]byte, error) {
	if n.Valid {
		return json.Marshal(n.Val)
	}
	return nullString, nil
}

func (n *Float32) UnmarshalJSON(b []byte) error {
	var s interface{}
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	return n.Scan(s)
}

var (
	_ json.Marshaler = (*Float32)(nil)
	_ json.Unmarshaler = (*Float32)(nil)
	TypeFloat32 = reflect.TypeOf(Float32{})
)