// Copyright 2017 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import (
	"fmt"
	"github.com/ego-plugin/fields"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
	"sync"
	"time"
)

type defaultValidator struct {
	once     sync.Once
	validate *validator.Validate
}

type sliceValidateError []error

func (err sliceValidateError) Error() string {
	var errMsgs []string
	for i, e := range err {
		if e == nil {
			continue
		}
		errMsgs = append(errMsgs, fmt.Sprintf("[%d]: %s", i, e.Error()))
	}
	return strings.Join(errMsgs, "\n")
}

var _ StructValidator = &defaultValidator{}

// ValidateStruct receives any kind of type, but only performed struct or pointer to struct type.
func (v *defaultValidator) ValidateStruct(obj interface{}) error {
	if obj == nil {
		return nil
	}

	value := reflect.ValueOf(obj)
	switch value.Kind() {
	case reflect.Ptr:
		return v.ValidateStruct(value.Elem().Interface())
	case reflect.Struct:
		return v.validateStruct(obj)
	case reflect.Slice, reflect.Array:
		count := value.Len()
		validateRet := make(sliceValidateError, 0)
		for i := 0; i < count; i++ {
			if err := v.ValidateStruct(value.Index(i).Interface()); err != nil {
				validateRet = append(validateRet, err)
			}
		}
		if len(validateRet) == 0 {
			return nil
		}
		return validateRet
	default:
		return nil
	}
}

// validateStruct receives struct type
func (v *defaultValidator) validateStruct(obj interface{}) error {
	v.lazyinit()
	return v.validate.Struct(obj)
}

// Engine returns the underlying validator engine which powers the default
// Validator instance. This is useful if you want to register custom validations
// or struct level validations. See validator GoDoc for more info -
// https://pkg.go.dev/github.com/go-playground/validator/v10
func (v *defaultValidator) Engine() interface{} {
	v.lazyinit()
	return v.validate
}

func (v *defaultValidator) lazyinit() {
	v.once.Do(func() {
		v.validate = validator.New()
		v.validate.SetTagName("binding")
		v.init()
	})
}

func (v *defaultValidator) GetValidate() *validator.Validate {
	return v.validate
}

func (v *defaultValidator) init() {
	// 注册string数组类型
	v.validate.RegisterCustomTypeFunc(func(value reflect.Value) interface{} {
		if val, ok := value.Interface().(field.FORM[string]); ok {
			return val.Val
		}
		return nil
	}, field.FORM[string]{})

	v.validate.RegisterCustomTypeFunc(func(value reflect.Value) interface{} {
		if val, ok := value.Interface().(field.FORM[bool]); ok {
			return val.Val
		}
		return nil
	}, field.FORM[bool]{})

	v.validate.RegisterCustomTypeFunc(func(value reflect.Value) interface{} {
		if val, ok := value.Interface().(field.Date); ok {
			return val.Val
		}
		return nil
	}, field.Date{})

	v.validate.RegisterCustomTypeFunc(func(value reflect.Value) interface{} {
		if val, ok := value.Interface().(field.Decimal); ok {
			return val.Decimal
		}
		return nil
	}, field.Decimal{})

	v.validate.RegisterCustomTypeFunc(func(value reflect.Value) interface{} {
		if val, ok := value.Interface().(field.FORM[float32]); ok {
			return val.Val
		}
		return nil
	}, field.FORM[float32]{})

	v.validate.RegisterCustomTypeFunc(func(value reflect.Value) interface{} {
		if val, ok := value.Interface().(field.FORM[float64]); ok {
			return val.Val
		}
		return nil
	}, field.FORM[float64]{})

	v.validate.RegisterCustomTypeFunc(func(value reflect.Value) interface{} {
		if val, ok := value.Interface().(field.FORM[int32]); ok {
			return val.Val
		}
		return nil
	}, field.FORM[int32]{})

	v.validate.RegisterCustomTypeFunc(func(value reflect.Value) interface{} {
		if val, ok := value.Interface().(field.FORM[int]); ok {
			return val.Val
		}
		return nil
	}, field.FORM[int]{})

	v.validate.RegisterCustomTypeFunc(func(value reflect.Value) interface{} {
		if val, ok := value.Interface().(field.FORM[int8]); ok {
			return val.Val
		}
		return nil
	}, field.FORM[int8]{})

	v.validate.RegisterCustomTypeFunc(func(value reflect.Value) interface{} {
		if val, ok := value.Interface().(field.FORM[int16]); ok {
			return val.Val
		}
		return nil
	}, field.FORM[int16]{})

	v.validate.RegisterCustomTypeFunc(func(value reflect.Value) interface{} {
		if val, ok := value.Interface().(field.FORM[int32]); ok {
			return val.Val
		}
		return nil
	}, field.FORM[int32]{})

	v.validate.RegisterCustomTypeFunc(func(value reflect.Value) interface{} {
		if val, ok := value.Interface().(field.FORM[int64]); ok {
			return val.Val
		}
		return nil
	}, field.FORM[int64]{})

	v.validate.RegisterCustomTypeFunc(func(value reflect.Value) interface{} {
		if val, ok := value.Interface().(field.FORM[uint32]); ok {
			return val.Val
		}
		return nil
	}, field.FORM[uint32]{})

	v.validate.RegisterCustomTypeFunc(func(value reflect.Value) interface{} {
		if val, ok := value.Interface().(field.FORM[uint]); ok {
			return val.Val
		}
		return nil
	}, field.FORM[uint]{})

	v.validate.RegisterCustomTypeFunc(func(value reflect.Value) interface{} {
		if val, ok := value.Interface().(field.FORM[uint8]); ok {
			return val.Val
		}
		return nil
	}, field.FORM[uint8]{})

	v.validate.RegisterCustomTypeFunc(func(value reflect.Value) interface{} {
		if val, ok := value.Interface().(field.FORM[uint16]); ok {
			return val.Val
		}
		return nil
	}, field.FORM[uint16]{})

	v.validate.RegisterCustomTypeFunc(func(value reflect.Value) interface{} {
		if val, ok := value.Interface().(field.FORM[uint32]); ok {
			return val.Val
		}
		return nil
	}, field.FORM[uint32]{})

	v.validate.RegisterCustomTypeFunc(func(value reflect.Value) interface{} {
		if val, ok := value.Interface().(field.FORM[uint64]); ok {
			return val.Val
		}
		return nil
	}, field.FORM[uint64]{})

	v.validate.RegisterCustomTypeFunc(func(value reflect.Value) interface{} {
		if val, ok := value.Interface().(field.Time); ok {
			return val.Val
		}
		return nil
	}, field.Time{})

	v.validate.RegisterCustomTypeFunc(func(value reflect.Value) interface{} {
		if val, ok := value.Interface().(field.FORM[time.Time]); ok {
			return val.Val
		}
		return nil
	}, field.FORM[time.Time]{})

	v.validate.RegisterCustomTypeFunc(func(value reflect.Value) interface{} {
		if val, ok := value.Interface().(field.JSON); ok {
			return val.Val
		}
		return nil
	}, field.JSON{})
}
