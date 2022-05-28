package fields

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"testing"
)

type forms struct {
	Name     String  `json:"name,omitempty" msgpack:"name,omitempty"`
	Email    *String `json:"email,omitempty" type:"string" msgpack:"email,omitempty"`
	Username string  `json:"username,omitempty" msgpack:"username,omitempty"`
	Password string  `json:"Password,omitempty" msgpack:"Password,omitempty"`
}

func TestStringMarshalJSON(t *testing.T) {
	form := new(forms)
	form.Username = "username"
	if err := form.Name.Scan("名字"); err != nil {
		t.Errorf("Scan: %s", err)
	}
	t.Logf("%v", form)
	b, err := json.Marshal(form)
	if err != nil {
		t.Errorf("json.Marshal err: %s", err)
		return
	}
	t.Logf("%s", b)
}

func TestStringUnmarshalJSON(t *testing.T) {
	form := new(forms)
	err := json.Unmarshal([]byte(`{"name":"my name","email":"", "username":"user", "Password":""}`), form)
	if err != nil {
		t.Errorf("json.Unmarshal err: %s", err)
		return
	}
	t.Logf("%v", form)
}

func TestStringMarshalMsgpack(t *testing.T) {
	form := new(forms)
	form.Username = "user"
	if err := form.Name.Scan("名字"); err != nil {
		t.Errorf("Scan: %s", err)
	}
	t.Logf("%v", form)
	b, err := msgpack.Marshal(form)
	if err != nil {
		t.Errorf("json.Marshal err: %s", err)
		return
	}
	t.Logf("%s", b)
}

func TestStringUnmarshalMsgpack(t *testing.T) {
	form := new(forms)
	form.Username = "user"
	if err := form.Name.Scan("名字"); err != nil {
		t.Errorf("Scan: %s", err)
	}
	b, _ := msgpack.Marshal(form)
	err := msgpack.Unmarshal(b, form)
	if err != nil {
		t.Errorf("json.Unmarshal err: %s", err)
		return
	}
	t.Logf("%v", form)
}
