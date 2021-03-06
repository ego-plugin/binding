// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import (
	"github.com/ego-plugin/binding/fields"
	"io/ioutil"
	"net/http"

	"github.com/golang/protobuf/proto"
)

type protobufBinding struct{}

func (protobufBinding) Name() string {
	return "protobuf"
}

func (b protobufBinding) Bind(req *http.Request, obj interface{}, lang string) error {
	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}
	// 写入默认值
	if err = fields.SetDefaultValue(obj); err != nil {
		return err
	}
	return b.BindBody(buf, obj, lang)
}

func (protobufBinding) BindBody(body []byte, obj interface{}, lang string) error {
	if err := proto.Unmarshal(body, obj.(proto.Message)); err != nil {
		return err
	}
	// Here it's same to return validate(obj), but util now we can't add
	// `binding:""` to the struct which automatically generate by gen-proto
	return nil
	// return validate(obj)
}
