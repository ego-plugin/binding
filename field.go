package binding

import (
	"github.com/ego-plugin/fields"
	"reflect"
	"sync"
	"time"
)

var (
	ScannerType      = reflect.TypeOf((*Scanner)(nil)).Elem()
	RegisterFormType = NewBindType()
)

// Scanner is an interface used by Scan.
type Scanner interface {
	Scan(value interface{}) (err error)
}

type bindType struct {
	m    map[reflect.Type]FormBindType
	lock *sync.RWMutex
}

// RegisterFormBindType
// 表单绑定类型注册(表单值) 付值,错误
type FormBindType func(inputValue []string) (reflect.Value, error)

func NewBindType() *bindType {
	return &bindType{
		m:    make(map[reflect.Type]FormBindType, 0),
		lock: new(sync.RWMutex),
	}
}

// Bind 绑定类型
func (p *bindType) Bind(t reflect.Type, fn FormBindType) {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.m[t] = fn
}

// 取得数据
func (p *bindType) Get() map[reflect.Type]FormBindType {
	p.lock.RLock()
	defer p.lock.RUnlock()
	return p.m
}

func (p *bindType) Exists(p2 reflect.Type) bool {
	p.lock.RLock()
	defer p.lock.RUnlock()
	for key := range p.m {
		if key == p2 {
			return true
		}
	}
	return false
}

func setTypesField(value reflect.Value, inputValue []string) error {
	// 给数据库类型付值
	for keyType, vFn := range RegisterFormType.Get() {
		if value.Type() == keyType {
			v, err := vFn(inputValue)
			if err != nil {
				return err
			}
			value.Set(v)
			continue
		}
	}
	return nil
}

func init() {
	RegisterFormType.Bind(field.FormTypeBool, func(inputValue []string) (reflect.Value, error) {
		v := field.FORM[bool]{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(field.DateType, func(inputValue []string) (reflect.Value, error) {
		v := field.Date{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(field.DecimalType, func(inputValue []string) (reflect.Value, error) {
		v := field.Decimal{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(field.FormTypeFloat32, func(inputValue []string) (reflect.Value, error) {
		v := field.FORM[float32]{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(field.FormTypeFloat64, func(inputValue []string) (reflect.Value, error) {
		v := field.FORM[float64]{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(field.FormTypeInt, func(inputValue []string) (reflect.Value, error) {
		v := field.FORM[int]{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(field.FormTypeInt8, func(inputValue []string) (reflect.Value, error) {
		v := field.FORM[int8]{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(field.FormTypeInt16, func(inputValue []string) (reflect.Value, error) {
		v := field.FORM[int16]{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(field.FormTypeInt32, func(inputValue []string) (reflect.Value, error) {
		v := field.FORM[int32]{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(field.FormTypeInt64, func(inputValue []string) (reflect.Value, error) {
		v := field.FORM[int64]{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(field.FormTypeUint, func(inputValue []string) (reflect.Value, error) {
		v := field.FORM[uint]{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(field.FormTypeUint8, func(inputValue []string) (reflect.Value, error) {
		v := field.FORM[uint8]{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(field.FormTypeUint16, func(inputValue []string) (reflect.Value, error) {
		v := field.FORM[uint16]{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(field.FormTypeUint32, func(inputValue []string) (reflect.Value, error) {
		v := field.FORM[uint32]{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(field.FormTypeUint64, func(inputValue []string) (reflect.Value, error) {
		v := field.FORM[uint64]{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(field.TimeType, func(inputValue []string) (reflect.Value, error) {
		v := field.Time{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(field.FormTypeTime, func(inputValue []string) (reflect.Value, error) {
		v := field.FORM[time.Time]{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(field.FormTypeString, func(inputValue []string) (reflect.Value, error) {
		v := field.FORM[string]{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(field.JSONType, func(inputValue []string) (reflect.Value, error) {
		v := field.JSON{}
		err := v.Scan(inputValue)
		return reflect.ValueOf(v), err
	})
}
