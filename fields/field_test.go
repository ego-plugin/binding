package fields

import (
	"encoding/json"
	"testing"
)

func TestSetDefaultValue(t *testing.T) {
	type test struct {
		Name     String `json:"name" default:"my name"`
		Email    string `json:"email" default:"mmmm@qq.com"`
		Username string `json:"username"`
		Password string `json:"Password" default:"pwd"`
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
