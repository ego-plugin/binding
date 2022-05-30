// Copyright 2018 Gin Core Team.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import "github.com/ego-plugin/binding/fields"

type uriBinding struct{}

func (uriBinding) Name() string {
	return "uri"
}

func (uriBinding) BindUri(m map[string][]string, obj interface{}, lang string) error {
	// 写入默认值
	if err := fields.SetDefaultValue(obj); err != nil {
		return err
	}

	if err := mapUri(obj, m); err != nil {
		return err
	}
	return validate(obj, lang)
}
