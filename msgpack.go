// Copyright 2017 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

//go:build !nomsgpack
// +build !nomsgpack

package binding

import (
	"bytes"
	"github.com/ego-plugin/binding/fields"
	"io"
	"net/http"

	"github.com/ugorji/go/codec"
)

type msgpackBinding struct{}

func (msgpackBinding) Name() string {
	return "msgpack"
}

func (msgpackBinding) Bind(req *http.Request, obj interface{}, lang string) error {
	// 写入默认值
	if err := fields.SetDefaultValue(obj); err != nil {
		return err
	}
	return decodeMsgPack(req.Body, obj, lang)
}

func (msgpackBinding) BindBody(body []byte, obj interface{}, lang string) error {
	// 写入默认值
	if err := fields.SetDefaultValue(obj); err != nil {
		return err
	}
	return decodeMsgPack(bytes.NewReader(body), obj, lang)
}

func decodeMsgPack(r io.Reader, obj interface{}, lang string) error {
	cdc := new(codec.MsgpackHandle)
	if err := codec.NewDecoder(r, cdc).Decode(&obj); err != nil {
		return err
	}
	return validate(obj, lang)
}
