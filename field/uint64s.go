package field

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"reflect"
)

type Uint64s struct {
	Val   []uint64
	Valid bool // Valid is true if Int64s is not NULL
}

// Scan 扫描写入值 interface.
func (n *Uint64s) Scan(value interface{}) (err error) {
	if value == nil {
		n.Val, n.Valid = make([]uint64, 0), false
		return err
	}
	// 取类型
	t := reflect.TypeOf(value)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	// 绑断是不是数组类型
	if t.Kind() != reflect.Slice {
		return errors.New("scan on int64s: no slice")
	}
	// 付值
	v := reflect.Indirect(reflect.ValueOf(value))
	for i := 0; i < t.Len(); i++ {
		var d uint64
		err = convertAssign(&d, v.Index(i))
		if err != nil {
			continue
		}
		n.Val = append(n.Val, d)
	}
	n.Valid = err == nil
	return err
}

func (n Uint64s) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Val, nil
}

func (n Uint64s) MarshalJSON() ([]byte, error) {
	if n.Valid {
		return json.Marshal(n.Val)
	}
	return nullString, nil
}

func (n *Uint64s) UnmarshalJSON(b []byte) error {
	var s interface{}
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	return n.Scan(s)
}

var (
	_ json.Marshaler = (*Uint64s)(nil)
	_ json.Unmarshaler = (*Uint64s)(nil)
	TypeUint64s = reflect.TypeOf(Uint64s{})
)
