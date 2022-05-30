// Copyright 2018 Gin Core Team.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import (
	"bytes"
	"github.com/ego-plugin/binding/fields"
	"io"
	"net/http"

	"gopkg.in/yaml.v2"
)

type yamlBinding struct{}

func (yamlBinding) Name() string {
	return "yaml"
}

func (yamlBinding) Bind(req *http.Request, obj interface{}, lang string) error {
	// 写入默认值
	if err := fields.SetDefaultValue(obj); err != nil {
		return err
	}
	return decodeYAML(req.Body, obj, lang)
}

func (yamlBinding) BindBody(body []byte, obj interface{}, lang string) error {
	// 写入默认值
	if err := fields.SetDefaultValue(obj); err != nil {
		return err
	}
	return decodeYAML(bytes.NewReader(body), obj, lang)
}

func decodeYAML(r io.Reader, obj interface{}, lang string) error {
	decoder := yaml.NewDecoder(r)
	if err := decoder.Decode(obj); err != nil {
		return err
	}
	return validate(obj, lang)
}
