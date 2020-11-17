package types

import (
	"testing"
)

func Test_UInt8ToNative(t *testing.T) {
	b := NewUInt8()
	b = 1
	if b.ToNative() != 1 {
		t.Error("UInt8ToNative Failed!")
	} else {
		t.Log("UInt8ToNative OK")
	}
}

func Test_UInt8ToInt64(t *testing.T) {
	b := NewUInt8()
	b = 1
	if b.ToInt64() != 1 {
		t.Error("UInt8ToInt64 Failed!")
	} else {
		t.Log("UInt8ToInt64 OK")
	}
}

func Test_UInt8ToInt32(t *testing.T) {
	b := NewUInt8()
	b = 1
	if b.ToInt32() != 1 {
		t.Error("UInt8ToInt32 Failed!")
	} else {
		t.Log("UInt8ToInt32 OK")
	}
}

func Test_UInt8ToInt16(t *testing.T) {
	b := NewUInt8()
	b = 1
	if b.ToInt16() != 1 {
		t.Error("UInt8ToInt16 Failed!")
	} else {
		t.Log("UInt8ToInt16 OK")
	}
}

func Test_UInt8ToInt8(t *testing.T) {
	b := NewUInt8()
	b = 1
	if b.ToInt8() != 1 {
		t.Error("UInt8ToInt8 Failed!")
	} else {
		t.Log("UInt8ToInt8 OK")
	}
}

func Test_UInt8ToUInt64(t *testing.T) {
	b := NewUInt8()
	b = 1
	if b.ToUInt64() != 1 {
		t.Error("UInt8ToUInt64 Failed!")
	} else {
		t.Log("UInt8ToUInt64 OK")
	}
}

func Test_UInt8ToUInt32(t *testing.T) {
	b := NewUInt8()
	b++
	if b.ToUInt32() != 1 {
		t.Error("UInt8ToUInt32 Failed!")
	} else {
		t.Log("UInt8ToUInt32 OK")
	}
}

func Test_UInt8ToUInt16(t *testing.T) {
	b := NewUInt8()
	b++
	if b.ToUInt16() != 1 {
		t.Error("UInt8ToUInt16 Failed!")
	} else {
		t.Log("UInt8ToUInt16 OK")
	}
}

func Test_UInt8ToUInt8(t *testing.T) {
	b := NewUInt8()
	b++
	if b.ToUInt8() != 1 {
		t.Error("UInt8ToUInt8 Failed!")
	} else {
		t.Log("UInt8ToUInt8 OK")
	}
}

func Test_UInt8ToFloat64(t *testing.T) {
	b := NewUInt8()
	b++
	if b.ToFloat64() != 1 {
		t.Error("UInt8ToFloat64 Failed!")
	} else {
		t.Log("UInt8ToFloat64 OK")
	}
}

func Test_UInt8ToFloat32(t *testing.T) {
	b := NewUInt8()
	b++
	if b.ToFloat32() != 1 {
		t.Error("UInt8ToFloat32 Failed!")
	} else {
		t.Log("UInt8ToFloat32 OK")
	}
}

func Test_UInt8ToInt(t *testing.T) {
	b := NewUInt8()
	b++
	if b.ToInt() != 1 {
		t.Error("UInt8ToInt Failed!")
	} else {
		t.Log("UInt8ToInt OK")
	}
}

func Test_UInt8ToUInt(t *testing.T) {
	b := NewUInt8()
	b++
	if b.ToUInt() != 1 {
		t.Error("UInt8ToUInt Failed!")
	} else {
		t.Log("UInt8ToUInt OK")
	}
}

func Test_UInt8ToFloat(t *testing.T) {
	b := NewUInt8()
	b++
	if b.ToFloat() != 1 {
		t.Error("UInt8ToFloat Failed!")
	} else {
		t.Log("UInt8ToFloat OK")
	}
}

func Test_UInt8ToBool(t *testing.T) {
	b := NewUInt8()
	if b.ToBool() != false {
		t.Error("UInt8ToBool Failed!")
	} else {
		t.Log("UInt8ToBool OK")
	}
	b = 123
	if b.ToBool() != true {
		t.Error("UInt8ToBool Failed!")
	} else {
		t.Log("UInt8ToBool OK")
	}
}

func Test_UInt8ToJSON(t *testing.T) {
	b := NewUInt8()
	if b.ToJSON().ToString() != "0" {
		t.Error("UInt8ToJSON Failed!")
	} else {
		t.Log("UInt8ToJSON OK")
	}
}

func Test_UInt8ToString(t *testing.T) {
	b := NewUInt8()
	if b.ToString() != "0" {
		t.Error("UInt8ToJSON Failed!")
	} else {
		t.Log("UInt8ToJSON OK")
	}
}

func Test_UInt8ToJSONString(t *testing.T) {
	b := NewUInt8()
	if s, _ := b.ToJSONString(); s != "0" {
		t.Error("ToJSONString Failed!")
	} else {
		t.Log("ToJSONString OK")
	}
}
