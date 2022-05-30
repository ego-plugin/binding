package fields

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"testing"
)

type formuint64 struct {
	Code Uint64 `json:"code,omitempty" msgpack:"name,omitempty"`
	//Num  []Uint64 `json:"num,omitempty" msgpack:"num,omitempty"`
	An uint64 `json:"an,omitempty" msgpack:"an,omitempty"`
}

func TestUint64_MarshalJSON(t *testing.T) {
	form := new(formuint64)
	if err := form.Code.Scan(2); err != nil {
		t.Errorf("scan: %s", err)
	}
	t.Logf("%v", form)
	b, err := json.Marshal(form)
	if err != nil {
		t.Errorf("json.Marshal err: %s", err)
		return
	}
	t.Logf("%s", b)
}

func TestUint64_UnmarshalJSON(t *testing.T) {
	form := new(formuint64)
	err := json.Unmarshal([]byte(`{"code":"1","num":23}`), form)
	if err != nil {
		t.Errorf("json.Unmarshal err: %s", err)
		return
	}
	t.Logf("%v", form)
}

func TestUint64_MarshalMsgpack(t *testing.T) {
	form := new(formuint64)
	form.Code.Scan(12012)

	t.Logf("%v", form)

	b, err := msgpack.Marshal(form)
	if err != nil {
		t.Errorf("Msgpack err: %s", err)
		return
	}

	t.Logf("%s", b)
}

func TestUint64_UnmarshalMsgpack(t *testing.T) {
	form := new(formuint64)
	form.Code.Scan(12012)

	b, _ := msgpack.Marshal(form)
	err := msgpack.Unmarshal(b, form)
	if err != nil {
		t.Errorf("Msgpack.Unmarshal err: %s", err)
		return
	}
	t.Logf("%v", form)
}
