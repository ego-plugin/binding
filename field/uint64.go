package field

import (
"database/sql/driver"
"encoding/json"
"reflect"
)

type Uint64 struct {
	Val   uint64
	Valid bool // Valid is true if Int64 is not NULL
}

// Scan 扫描写入值 interface.
func (n *Uint64) Scan(value interface{}) (err error) {
	if value == nil {
		n.Val, n.Valid = 0, false
		return err
	}
	err = convertAssign(&n.Val, value)
	n.Valid = err == nil
	return err
}

func (n Uint64) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Val, nil
}

func (n Uint64) MarshalJSON() ([]byte, error) {
	if n.Valid {
		return json.Marshal(n.Val)
	}
	return nullString, nil
}

func (n *Uint64) UnmarshalJSON(b []byte) error {
	var s interface{}
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	return n.Scan(s)
}

var (
	_ json.Marshaler = (*Uint64)(nil)
	_ json.Unmarshaler = (*Uint64)(nil)
	TypeUint64 = reflect.TypeOf(Uint64{})
)
