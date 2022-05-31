package fields

import "testing"

func TestColumns(t *testing.T) {
	type SystemMenu struct {
		Id       uint64 `json:"id" minimum:"1" db:"-" msgpack:"id" description:"按ID更新"`
		ParentID uint64 `json:"parent_id" msgpack:"parent_id" description:"父ID"`
	}

	t.Log(Columns(new(SystemMenu)))
}

func TestContainsColumns(t *testing.T) {
	t.Logf("result: %v", ContainsColumns([]string{"parent_id", "", "route_path"}, []string{"parent_id", "route_path", "weigh", "children"}))
}
