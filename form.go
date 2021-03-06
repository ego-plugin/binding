// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import (
	"net/http"

	"github.com/ego-plugin/binding/fields"
)

const defaultMemory = 32 << 20

type (
	formBinding          struct{}
	formPostBinding      struct{}
	formMultipartBinding struct{}
)

func (formBinding) Name() string {
	return "form"
}

func (formBinding) Bind(req *http.Request, obj interface{}, lang string) error {
	// 检查
	if err := req.ParseForm(); err != nil {
		return err
	}

	if err := req.ParseMultipartForm(defaultMemory); err != nil {
		if err != http.ErrNotMultipart {
			return err
		}
	}

	// 写入默认值
	if err := fields.SetDefaultValue(obj); err != nil {
		return err
	}

	if err := mapForm(obj, req.Form); err != nil {
		return err
	}
	return Validate(obj, lang)
}

func (formPostBinding) Name() string {
	return "form-urlencoded"
}

func (formPostBinding) Bind(req *http.Request, obj interface{}, lang string) error {
	if err := req.ParseForm(); err != nil {
		return err
	}
	// 写入默认值
	if err := fields.SetDefaultValue(obj); err != nil {
		return err
	}
	if err := mapForm(obj, req.PostForm); err != nil {
		return err
	}
	return Validate(obj, lang)
}

func (formMultipartBinding) Name() string {
	return "multipart/form-data"
}

func (formMultipartBinding) Bind(req *http.Request, obj interface{}, lang string) error {
	if err := req.ParseMultipartForm(defaultMemory); err != nil {
		return err
	}
	// 写入默认值
	if err := fields.SetDefaultValue(obj); err != nil {
		return err
	}
	if err := mappingByPtr(obj, (*multipartRequest)(req), "form"); err != nil {
		return err
	}

	return Validate(obj, lang)
}
