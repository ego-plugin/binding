package binding

import (
	"github.com/ego-plugin/binding/fields"
	"reflect"
	"sync"
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

// FormBindType
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

// Get 取得数据
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
	RegisterFormType.Bind(fields.StringType, func(inputValue []string) (reflect.Value, error) {
		v := fields.String{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(fields.BoolType, func(inputValue []string) (reflect.Value, error) {
		v := fields.Bool{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(fields.TimeDateType, func(inputValue []string) (reflect.Value, error) {
		v := fields.TimeDate{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(fields.TimeNumberType, func(inputValue []string) (reflect.Value, error) {
		v := fields.TimeNumber{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(fields.DecimalType, func(inputValue []string) (reflect.Value, error) {
		v := fields.Decimal{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(fields.TimeType, func(inputValue []string) (reflect.Value, error) {
		v := fields.Time{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(fields.Float32Type, func(inputValue []string) (reflect.Value, error) {
		v := fields.Float32{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(fields.Float64Type, func(inputValue []string) (reflect.Value, error) {
		v := fields.Float64{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(fields.IntType, func(inputValue []string) (reflect.Value, error) {
		v := fields.Int{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(fields.Int32Type, func(inputValue []string) (reflect.Value, error) {
		v := fields.Int32{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(fields.Int64Type, func(inputValue []string) (reflect.Value, error) {
		v := fields.Int64{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(fields.UintType, func(inputValue []string) (reflect.Value, error) {
		v := fields.Uint{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(fields.Uint32Type, func(inputValue []string) (reflect.Value, error) {
		v := fields.Uint32{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(fields.Uint64Type, func(inputValue []string) (reflect.Value, error) {
		v := fields.Uint{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})
}
