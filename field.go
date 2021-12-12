package binding


import (
	"github.com/ego-plugin/binding/field"
	"reflect"
	"sync"
)

var (
	ScannerType  = reflect.TypeOf((*Scanner)(nil)).Elem()
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
	RegisterFormType.Bind(field.TypeBool, func(inputValue []string) (reflect.Value, error) {
		v := field.Bool{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(field.TypeDate, func(inputValue []string) (reflect.Value, error) {
		v := field.Date{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(field.TypeDecimal, func(inputValue []string) (reflect.Value, error) {
		v := field.Decimal{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(field.TypeFloat32, func(inputValue []string) (reflect.Value, error) {
		v := field.Float32{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(field.TypeFloat64, func(inputValue []string) (reflect.Value, error) {
		v := field.Float64{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(field.TypeFloat64s, func(inputValue []string) (reflect.Value, error) {
		v := field.Float64s{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(field.TypeInt32, func(inputValue []string) (reflect.Value, error) {
		v := field.Int32{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(field.TypeInt32s, func(inputValue []string) (reflect.Value, error) {
		v := field.Int32s{}
		err := v.Scan(inputValue)
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(field.TypeUint32, func(inputValue []string) (reflect.Value, error) {
		v := field.Uint32{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(field.TypeInt64, func(inputValue []string) (reflect.Value, error) {
		v := field.Int64{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(field.TypeInt64s, func(inputValue []string) (reflect.Value, error) {
		v := field.Int64s{}
		err := v.Scan(inputValue)
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(field.TypeUint64, func(inputValue []string) (reflect.Value, error) {
		v := field.Uint64{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(field.TypeUint64s, func(inputValue []string) (reflect.Value, error) {
		v := field.Uint64s{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(field.TypeTime, func(inputValue []string) (reflect.Value, error) {
		v := field.Time{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(field.TypeString, func(inputValue []string) (reflect.Value, error) {
		v := field.String{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})

	RegisterFormType.Bind(field.TypeStrings, func(inputValue []string) (reflect.Value, error) {
		v := field.Strings{}
		err := v.Scan(inputValue[0])
		return reflect.ValueOf(v), err
	})
}
