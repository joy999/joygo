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

func Test_JSONMapGet(t *testing.T) {
	m := NewJSONMap()
	m.Set("a", 1)
	m.Set("a1", String("abc"))
	m.Set("b", JSONMap{`b`: 23.3})

	if AnyToInt(m.Get("a")) != 1 {
		t.Error("Test_JSONMapGet Failed!")
	}
	if AnyToString(m.Get("a1")) != "abc" {
		t.Error("Test_JSONMapGet Failed!")
	}
	if m.Get("abc") != nil {
		t.Error("Test_JSONMapGet Failed!")
	}

	if m.GetAsJSON("a").ToInt() != 1 {
		t.Error("Test_JSONMapGetAsJSON Failed!")
	}
	if m.GetAsJSON("a1").ToString() != "abc" {
		t.Error("Test_JSONMapGetAsJSON Failed!")
	}
	if m.GetAsJSON("abc") != nil {
		t.Error("Test_JSONMapGetAsJSON Failed!")
	}
	m.Remove("abc")
	m.Remove("a1")
	m.Remove("b")

	if j, _ := m.ToJSON().ToJSONString(); j != `{"a":1}` {
		t.Error("Test_JSONMapToJSON Failed!")
	}
	m.Clear()
	if j, _ := m.ToJSONString(); j != `{}` {
		t.Error("Test_JSONMapToJSON Failed!")
	}
}
