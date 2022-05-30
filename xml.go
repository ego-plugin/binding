// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import (
	"bytes"
	"encoding/xml"
	"github.com/ego-plugin/binding/fields"
	"io"
	"net/http"
)

type xmlBinding struct{}

func (xmlBinding) Name() string {
	return "xml"
}

func (xmlBinding) Bind(req *http.Request, obj interface{}, lang string) error {
	// 写入默认值
	if err := fields.SetDefaultValue(obj); err != nil {
		return err
	}
	return decodeXML(req.Body, obj, lang)
}

func (xmlBinding) BindBody(body []byte, obj interface{}, lang string) error {
	// 写入默认值
	if err := fields.SetDefaultValue(obj); err != nil {
		return err
	}
	return decodeXML(bytes.NewReader(body), obj, lang)
}
func decodeXML(r io.Reader, obj interface{}, lang string) error {
	decoder := xml.NewDecoder(r)
	if err := decoder.Decode(obj); err != nil {
		return err
	}
	return validate(obj, lang)
}
