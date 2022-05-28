package fields

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"testing"
)

type formuint64 struct {
	Code Uint64 `json:"code,omitempty" msgpack:"name,omitempty"`
	Num  Uint64 `json:"num,omitempty" type:"string" msgpack:"num,omitempty"`
}

func TestUint64MarshalJSON(t *testing.T) {
	form := new(formuint64)
	form.Code.Scan(1)
	t.Logf("%v", form)
	b, err := json.Marshal(form)
	if err != nil {
		t.Errorf("json.Marshal err: %s", err)
		return
	}
	t.Logf("%s", b)
}

func TestUint64UnmarshalJSON(t *testing.T) {
	form := new(formuint64)
	err := json.Unmarshal([]byte(`{"code":"1","num":23}`), form)
	if err != nil {
		t.Errorf("json.Unmarshal err: %s", err)
		return
	}
	t.Logf("%v", form)
}

func TestUint64MarshalMsgpack(t *testing.T) {
	form := new(formuint64)
	form.Code.Scan(12012)
	form.Num.Scan(nil)

	t.Logf("%v", form)

	b, err := msgpack.Marshal(form)
	if err != nil {
		t.Errorf("Msgpack err: %s", err)
		return
	}

	t.Logf("%s", b)
}

func TestUint64UnmarshalMsgpack(t *testing.T) {
	form := new(formuint64)
	form.Code.Scan(12012)
	form.Num.Scan(nil)

	b, _ := msgpack.Marshal(form)
	err := msgpack.Unmarshal(b, form)
	if err != nil {
		t.Errorf("Msgpack.Unmarshal err: %s", err)
		return
	}
	t.Logf("%v", form)
}
