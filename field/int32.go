package field

import (
	"database/sql/driver"
	"encoding/json"
	"reflect"
)

type Int32 struct {
	Val   int32
	Valid bool // Valid is true if Int32 is not NULL
}

// Scan 扫描写入值 interface.
func (n *Int32) Scan(value interface{}) (err error) {
	if value == nil {
		n.Val, n.Valid = 0, false
		return err
	}
	err = convertAssign(&n.Val, value)
	n.Valid = err == nil
	return err
}

func (n Int32) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Val, nil
}

func (n Int32) MarshalJSON() ([]byte, error) {
	if n.Valid {
		return json.Marshal(n.Val)
	}
	return nullString, nil
}

func (n *Int32) UnmarshalJSON(b []byte) error {
	var s interface{}
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	return n.Scan(s)
}

var (
	_ json.Marshaler   = (*Int32)(nil)
	_ json.Unmarshaler = (*Int32)(nil)
	_ driver.Valuer    = (*Int32)(nil)

	TypeInt32 = reflect.TypeOf(Int32{})
)
