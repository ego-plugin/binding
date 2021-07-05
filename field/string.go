package field

import (
	"database/sql/driver"
	"encoding/json"
	"reflect"
)

type String struct {
	Val   string
	Valid bool
}

// Scan 扫描写入值 interface.
func (n *String) Scan(value interface{}) (err error) {
	if value == nil {
		n.Val, n.Valid = "", false
		return err
	}
	err = convertAssign(&n.Val, value)
	n.Valid = err == nil
	return err
}

func (n String) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Val, nil
}

func (ns String) MarshalJSON() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.Val)
	}
	return nullString, nil
}

func (ns *String) UnmarshalJSON(b []byte) error {
	var s interface{}
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	return ns.Scan(s)
}

var (
	nullString = []byte("null")
	TypeString = reflect.TypeOf(String{})
)
