package types

import (
	"testing"
)

func Test_JSONStringNew(t *testing.T) {
	s := NewJSONString(`{"a":1,"a1":"abc","arr":[1,2,3],"b":{"b":23.3}}`)
	t.Log(s)
}

func Test_JSONStringDecodeToJSON(t *testing.T) {
	s := NewJSONString(`{"a":1,"a1":"abc","arr":[1,2,3],"b":{"b":23.3}}`)
	json, err := s.DecodeToJSON()

	if err != nil {
		t.Error(err)
		return
	}

	t.Log(json)
	m := json.ToJSONMap()
	if !m.Exists("a") {
		t.Error("Check Decode Map Error!")
	}
	arr, ok := m["arr"]
	if !ok {
		t.Error("get Map[arr] Error!")
	}
	jArr := AnyToJSONArray(arr)
	j0 := jArr.GetAsJSON(0)
	j1 := jArr.GetAsJSON(1)
	j2 := jArr.GetAsJSON(2)
	n0 := j0.ToInt64()
	n1 := j1.ToInt64()
	n2 := j2.ToInt64()
	if n0 != 1 || n1 != 2 || n2 != 3 {
		t.Error("get jArr of Map[arr] Error!", j0, n0, j1, n1, j2, n2)
	}
}
