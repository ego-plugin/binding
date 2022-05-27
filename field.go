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
	RegisterFormType.Bind(fields.FormTypeBool, func(inputValue []string) (reflect.Value, error) {
		v := fields.FORM[bool]{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(fields.DateType, func(inputValue []string) (reflect.Value, error) {
		v := fields.Date{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(fields.DecimalType, func(inputValue []string) (reflect.Value, error) {
		v := fields.Decimal{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(fields.FormTypeFloat32, func(inputValue []string) (reflect.Value, error) {
		v := fields.FORM[float32]{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(fields.FormTypeFloat64, func(inputValue []string) (reflect.Value, error) {
		v := fields.FORM[float64]{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(fields.FormTypeInt, func(inputValue []string) (reflect.Value, error) {
		v := fields.FORM[int]{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(fields.FormTypeInt8, func(inputValue []string) (reflect.Value, error) {
		v := fields.FORM[int8]{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(fields.FormTypeInt16, func(inputValue []string) (reflect.Value, error) {
		v := fields.FORM[int16]{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(fields.FormTypeInt32, func(inputValue []string) (reflect.Value, error) {
		v := fields.FORM[int32]{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(fields.FormTypeInt64, func(inputValue []string) (reflect.Value, error) {
		v := fields.FORM[int64]{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(fields.FormTypeUint, func(inputValue []string) (reflect.Value, error) {
		v := fields.FORM[uint]{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(fields.FormTypeUint8, func(inputValue []string) (reflect.Value, error) {
		v := fields.FORM[uint8]{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(fields.FormTypeUint16, func(inputValue []string) (reflect.Value, error) {
		v := fields.FORM[uint16]{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(fields.FormTypeUint32, func(inputValue []string) (reflect.Value, error) {
		v := fields.FORM[uint32]{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(fields.FormTypeUint64, func(inputValue []string) (reflect.Value, error) {
		v := fields.FORM[uint64]{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(fields.TimeType, func(inputValue []string) (reflect.Value, error) {
		v := fields.Time{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(fields.FormTypeTime, func(inputValue []string) (reflect.Value, error) {
		v := fields.FORM[time.Time]{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(fields.FormTypeString, func(inputValue []string) (reflect.Value, error) {
		v := fields.FORM[string]{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(fields.JSONType, func(inputValue []string) (reflect.Value, error) {
		v := fields.JSON{}
		err := v.Scan(inputValue)
		return reflect.ValueOf(v), err
	})
}
