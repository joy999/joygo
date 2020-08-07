package types

import (
	"testing"
)

func Test_JSONMapNew(t *testing.T) {
	m := NewJSONMap()
	m.Set("a", 1)
	m.Set(`b`, JSONMap{`b`: JSON{23.3}})
	t.Log(m)
}

func Test_JSONMapToJSONString(t *testing.T) {
	m := NewJSONMap()
	m.Set("a", 1)
	m.Set(`b`, JSONMap{`b`: 23.3})
	m.Set("a1", String("abc"))
	m[`arr`] = JSONArray{1, 2, 3}

	if js, err := m.ToJSONString(); err != nil || js != `{"a":1,"a1":"abc","arr":[1,2,3],"b":{"b":23.3}}` {
		t.Error(err)
	} else {
		t.Log(js)
	}
}

func Test_JSONMapExists(t *testing.T) {
	m := NewJSONMap()
	m.Set("a", 1)
	m.Set("a1", String("abc"))
	m.Set("b", JSONMap{`b`: 23.3})

	n1 := m.Exists("a")
	n2 := m.Exists("a1")
	n3 := m.Exists("a2")

	if !n1 || !n2 || n3 {
		t.Error("Test_Exists FAILED!", n1, n2, n3)
	} else {
		t.Log("Test_Exists: ", n1, n2, n3)
	}
}
