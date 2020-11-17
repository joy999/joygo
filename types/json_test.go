package types

import (
	"testing"
)

func Test_JsonToJson(t *testing.T) {
	j := NewJSON("abc")
	if s, _ := j.ToJSONString(); s != `"abc"` {
		t.Errorf("Test_JsonToJson Failed! %v", s)
	}
	if ToJson(j).ToString() != `"abc"` {
		t.Errorf("Test_JsonToJson Failed! %v", j)
	}
}

func Test_JsonToNumber(t *testing.T) {
	j := NewJSON("123")
	if j.ToInt() != 123 {
		t.Error("Test_JsonToNumber Failed!")
	}
	if j.ToInt8() != 123 {
		t.Error("Test_JsonToNumber Failed!")
	}
	if j.ToInt16() != 123 {
		t.Error("Test_JsonToNumber Failed!")
	}
	if j.ToInt32() != 123 {
		t.Error("Test_JsonToNumber Failed!")
	}
	if j.ToInt64() != 123 {
		t.Error("Test_JsonToNumber Failed!")
	}
	if j.ToUInt() != 123 {
		t.Errorf("Test_JsonToNumber Failed! %v", j.ToUInt())
	}
	if j.ToUInt8() != 123 {
		t.Error("Test_JsonToNumber Failed!")
	}
	if j.ToUInt16() != 123 {
		t.Error("Test_JsonToNumber Failed!")
	}
	if j.ToUInt32() != 123 {
		t.Error("Test_JsonToNumber Failed!")
	}
	if j.ToUInt64() != 123 {
		t.Error("Test_JsonToNumber Failed!")
	}
	if j.ToFloat() != 123 {
		t.Error("Test_JsonToNumber Failed!")
	}
	if j.ToFloat32() != 123 {
		t.Error("Test_JsonToNumber Failed!")
	}
	if j.ToFloat64() != 123 {
		t.Error("Test_JsonToNumber Failed!")
	}

	j.Assign(-123.123)
	if j.ToInt8() != -123 {
		t.Error("Test_JsonToNumber Failed!")
	}
	if j.ToInt16() != -123 {
		t.Error("Test_JsonToNumber Failed!")
	}
	if j.ToInt32() != -123 {
		t.Error("Test_JsonToNumber Failed!")
	}
	if j.ToInt64() != -123 {
		t.Error("Test_JsonToNumber Failed!")
	}
	if j.ToFloat() != -123.123 {
		t.Error("Test_JsonToNumber Failed!")
	}
	if j.ToFloat32() != -123.123 {
		t.Error("Test_JsonToNumber Failed!")
	}
	if j.ToFloat64() != -123.123 {
		t.Error("Test_JsonToNumber Failed!")
	}
	if j.ToUInt() != 0 {
		t.Error("Test_JsonToNumber Failed!")
	}
	if j.ToUInt8() != 0 {
		t.Error("Test_JsonToNumber Failed!")
	}
	if j.ToUInt8() != 0 {
		t.Error("Test_JsonToNumber Failed!")
	}
	if j.ToUInt16() != 0 {
		t.Error("Test_JsonToNumber Failed!")
	}
	if j.ToUInt32() != 0 {
		t.Error("Test_JsonToNumber Failed!")
	}
	if j.ToUInt64() != 0 {
		t.Error("Test_JsonToNumber Failed!")
	}
}

func Test_JsonToJSONArray(t *testing.T) {
	arr := [...]Any{1, 2, 3}
	j := NewJSON(arr[:])
	if s, _ := j.ToJSONArray().ToJSONString(); s != `[1,2,3]` {
		t.Errorf("Test_JsonToJSONArray Failed! %v", s)
	}
}

func Test_JsonToJSONMap(t *testing.T) {
	arr := [...]Any{1, 2, 3}
	j := NewJSON(arr[:])
	if s, _ := j.ToJSONMap().ToJSONString(); s != `{"0":1,"1":2,"2":3}` {
		t.Errorf("Test_JsonToJSONArray Failed! %v", s)
	}
	if s, _ := j.ToJSONMapPtr().ToJSONString(); s != `{"0":1,"1":2,"2":3}` {
		t.Errorf("Test_JsonToJSONArray Failed! %v", s)
	}
}
