package types

import (
	"testing"
)

func Test_NewString(t *testing.T) {
	t.Log("Test_NewString")
	s := NewString("abc:123")
	ar := s.Explode(":")
	//t.Log(ar)
	if ar[0] != "abc" || ar[1] != "123" {
		t.Error("Explode ERROR")
	}
	len := s.GetLength()
	//t.Log(len)
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
		t.Log("JSON String: ", js)
	}

	s = "   123 4   "
	if s.TrimSpace() != `123 4` {
		t.Error(`TrimSpace failed`)
	}

	s = "aDc123 "
	if !s.IsSame("aDc123 ") {
		t.Error("IsSame Failed")
	}
	if s.IsSame("aDc12") {
		t.Error("!IsSame Failed")
	}
}

func Test_StringToUnixLocalTimeStamp(t *testing.T) {
	s := NewString("1984-01-14 23:32:23")
	i64 := s.ToUnixLocalTimeStamp("Y-m-d H:i:s")
	if i64 != 442942343 {
		t.Error("ToUnixLocalTimeStamp Failed")
	}
}

func Test_StringCoder(t *testing.T) {
	s := NewString("abc")
	gs := s.UTF8ToGBK()
	us := gs.GBKToUTF8()

	if gs != "abc" || us != "abc" {
		t.Error("Test_StringCoder Failed")
	}

	s = "您好"
	gs = s.UTF8ToGBK()
	us = gs.GBKToUTF8()

	if gs == "您好" {
		t.Errorf("Test_StringCoder Failed %v", gs)
	}

	if us != "您好" {
		t.Errorf("Test_StringCoder Failed %v", us)
	}
}

func Test_StringWriter(t *testing.T) {
	s := NewString("")
	w := s.NewWriter()
	w.Write([]byte("hello,"))
	w.Write([]byte("world!"))

	if s != "hello,world!" {
		t.Errorf("Test_StringCoder Failed %v", s)
	}
}

func Test_StringToJson(t *testing.T) {
	s := NewString("abc")
	if s.ToJson().ToString() != "abc" {
		t.Error("Test_JSONStringToJson Error!")
	}
}

func Test_StringMatchFind(t *testing.T) {
	s := NewString("abc123acc123")
	m := s.MatchFind("([a-z]+)(\\d+)")
	j := JSON{m}
	if js, _ := j.ToJSONString(); js != `["abc123","abc","123"]` {
		t.Errorf("Test_StringMatchFind Failed %v", js)
	}

	if m = s.MatchFind("ddd"); m != nil {
		t.Error("Test_StringMatchFind Error!")
	}
}

func Test_StringMatchAllFind(t *testing.T) {
	s := NewString("abc123acc122")
	m := s.MatchAllFind("([a-z]+)(\\d+)")
	j := JSON{m}
	if js, _ := j.ToJSONString(); js != `[["abc123","abc","123"],["acc122","acc","122"]]` {
		t.Errorf("Test_StringCoder Failed %v", js)
	}

	if m = s.MatchAllFind("ddd"); m != nil {
		t.Error("Test_StringMatchAllFind Error!")
	}
}
