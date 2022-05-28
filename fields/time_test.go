package fields

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"testing"
)

type formTime struct {
	CreateAt Time `json:"create_at,omitempty" msgpack:"create_at,omitempty"`
	UpdateAt Time `json:"update_at,omitempty" msgpack:"update_at,omitempty"`
}

func TestTime_MarshalJSON(t *testing.T) {
	form := new(formTime)
	form.CreateAt.Scan("12321321")
	form.UpdateAt.Scan("2020-10-12 11:11:11")

	t.Logf("%v", form)
	b, err := json.Marshal(form)
	if err != nil {
		t.Errorf("json.Marshal err: %s", err)
		return
	}
	t.Logf("%s", b)
}

func TestTime_UnmarshalJSON(t *testing.T) {
	form := new(formTime)
	err := json.Unmarshal([]byte(`{"create_at": "9182938322","update_at":"2020-10-12 11:11:11"}`), form)
	if err != nil {
		t.Errorf("json.Unmarshal err: %s", err)
		return
	}
	t.Logf("%v", form)
}

func TestTime_MarshalMsgpack(t *testing.T) {
	form := new(formTime)
	form.CreateAt.Scan("12321321")
	form.UpdateAt.Scan("2020-10-12 11:11:11")

	t.Logf("%v", form)
	b, err := msgpack.Marshal(form)
	if err != nil {
		t.Errorf("json.Marshal err: %s", err)
		return
	}
	t.Logf("%s", b)
}

func TestTime_UnmarshalMsgpack(t *testing.T) {
	form := new(formTime)
	form.CreateAt.Scan("12321321")
	form.UpdateAt.Scan("2020-10-12 11:11:11")

	b, _ := msgpack.Marshal(form)
	err := msgpack.Unmarshal(b, form)
	if err != nil {
		t.Errorf("json.Unmarshal err: %s", err)
		return
	}
	t.Logf("%v", form)
}
