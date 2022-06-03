package fields

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

type Uint64s struct {
	Val   []uint64
	Valid bool
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

	v := reflect.Indirect(reflect.ValueOf(value))
	switch t.Kind() {
	case reflect.Slice:
		for i := 0; i < t.Len(); i++ {
			var d uint64
			err = convertAssign(&d, v.Index(i).Interface())
			if err != nil {
				continue
			}
			n.Val = append(n.Val, d)
		}
		n.Valid = err == nil
		return err
	case reflect.String:
		b := StringToBytes(v.String())
		err = msgpack.Unmarshal(b, &n.Val)
		n.Valid = err == nil
		return err
	}

	return errors.New("scan slice err")
}

func (n Uint64s) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	b, err := msgpack.Marshal(n.Val)
	n.Valid = err == nil
	return BytesToString(b), err
}

func (n Uint64s) ValidateValuer() any {
	if !n.Valid {
		return nil
	}
	return n.Val
}

func (n *Uint64s) NilValue() *Uint64s {
	if !n.Valid {
		return nil
	}
	return n
}

func (n *Uint64s) Ptr() *[]uint64 {
	if !n.Valid {
		return nil
	}
	return &n.Val
}

func (n Uint64s) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return nullString, nil
	}
	return json.Marshal(n.Val)
}

func (n *Uint64s) UnmarshalJSON(b []byte) error {
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

func (n Uint64s) MarshalMsgpack() ([]byte, error) {
	if !n.Valid {
		return nullString, nil
	}
	return msgpack.Marshal(n.Val)
}

func (n *Uint64s) UnmarshalMsgpack(b []byte) error {
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

var (
	_ ValueScanner = (*Uint64s)(nil)
)
