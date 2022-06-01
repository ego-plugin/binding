package fields

import (
	"encoding/json"
	"testing"
)

func TestSetDefaultValue(t *testing.T) {
	type testa struct {
		Username String `json:"username" default:"user"`
		Password string `json:"Password" default:"pwd"`
	}

	type test struct {
		Name  String `json:"name" default:"my name"`
		Email string `json:"email" default:"mmmm@qq.com"`
		testa
	}

	v := test{}
	if err := SetDefaultValue(&v); err != nil {
		t.Error(err)
		return
	}

	d, err := json.Marshal(v)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%s", d)
}
