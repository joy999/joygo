package types

import (
	"testing"
)

func Test_NewString(t *testing.T) {
	t.Log("Test_NewString")
	s := NewString("abc:123")
	ar := s.Explode(":")
	t.Log(ar)
	if ar[0] != "abc" || ar[1] != "123" {
		t.Error("Explode ERROR")
	}
	len := s.GetLength()
	t.Log(len)
	if len != 7 {
		t.Error("GetLength ERROR")
	}
	if !s.Match("abc") {
		t.Error(`Match abc failed`)
	}
	if !s.Match("123") {
		t.Error(`Match 123 failed`)
	}
	if !s.Match("123$") {
		t.Error(`Match 123$ failed`)
	}
	if !s.Match("^abc") {
		t.Error(`Match ^abc failed`)
	}
	if s.Match("abc$") {
		t.Error(`Match NOT abc$ failed`)
	}
	if s.Match("^123") {
		t.Error(`Match NOT ^123 failed`)
	}
	s = "123"
	i := s.ToInt()
	if i != 123 {
		t.Error(`ToInt failed`)
	}
	u64 := s.ToUInt64()
	if u64 != 123 {
		t.Error(`ToUInt64 failed`)
	}
	s = "-123.123"
	f := s.ToFloat64()
	if f != -123.123 {
		t.Error(`ToFloat64 failed`)
	}
	i = s.ToInt()
	if i != -123 {
		t.Error(`Float ToInt failed`, i)
	}
	u64 = s.ToUInt64()
	if u64 != 0 {
		t.Error(`-123.123 ToUInt64 failed`, i)
	}
	if js, err := s.ToJSONString(); err != nil {
		t.Error(`-123.123 to ToJSONString failed`, err)
		if js != `-123.123` {
			t.Error(`-123.123 to ToJSONString failed2`, js)
		}
	} else {
		t.Log("JSON String: ", js, nil)
	}

	s = "   123 4   "
	if s.TrimSpace() != `123 4` {
		t.Error(`TrimSpace failed`)
	}

}
