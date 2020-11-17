package types

import (
	"testing"
)

func Test_UInt32ToNative(t *testing.T) {
	b := NewUInt32()
	b = 1
	if b.ToNative() != 1 {
		t.Error("UInt32ToNative Failed!")
	} else {
		t.Log("UInt32ToNative OK")
	}
}

func Test_UInt32ToInt64(t *testing.T) {
	b := NewUInt32()
	b = 1
	if b.ToInt64() != 1 {
		t.Error("UInt32ToInt64 Failed!")
	} else {
		t.Log("UInt32ToInt64 OK")
	}
}

func Test_UInt32ToInt32(t *testing.T) {
	b := NewUInt32()
	b = 1
	if b.ToInt32() != 1 {
		t.Error("UInt32ToInt32 Failed!")
	} else {
		t.Log("UInt32ToInt32 OK")
	}
}

func Test_UInt32ToInt16(t *testing.T) {
	b := NewUInt32()
	b = 1
	if b.ToInt16() != 1 {
		t.Error("UInt32ToInt16 Failed!")
	} else {
		t.Log("UInt32ToInt16 OK")
	}
}

func Test_UInt32ToInt8(t *testing.T) {
	b := NewUInt32()
	b = 1
	if b.ToInt8() != 1 {
		t.Error("UInt32ToInt8 Failed!")
	} else {
		t.Log("UInt32ToInt8 OK")
	}
}

func Test_UInt32ToUInt64(t *testing.T) {
	b := NewUInt32()
	b = 1
	if b.ToUInt64() != 1 {
		t.Error("UInt32ToUInt64 Failed!")
	} else {
		t.Log("UInt32ToUInt64 OK")
	}
}

func Test_UInt32ToUInt32(t *testing.T) {
	b := NewUInt32()
	b++
	if b.ToUInt32() != 1 {
		t.Error("UInt32ToUInt32 Failed!")
	} else {
		t.Log("UInt32ToUInt32 OK")
	}
}

func Test_UInt32ToUInt16(t *testing.T) {
	b := NewUInt32()
	b++
	if b.ToUInt16() != 1 {
		t.Error("UInt32ToUInt16 Failed!")
	} else {
		t.Log("UInt32ToUInt16 OK")
	}
}

func Test_UInt32ToUInt8(t *testing.T) {
	b := NewUInt32()
	b++
	if b.ToUInt8() != 1 {
		t.Error("UInt32ToUInt8 Failed!")
	} else {
		t.Log("UInt32ToUInt8 OK")
	}
}

func Test_UInt32ToFloat64(t *testing.T) {
	b := NewUInt32()
	b++
	if b.ToFloat64() != 1 {
		t.Error("UInt32ToFloat64 Failed!")
	} else {
		t.Log("UInt32ToFloat64 OK")
	}
}

func Test_UInt32ToFloat32(t *testing.T) {
	b := NewUInt32()
	b++
	if b.ToFloat32() != 1 {
		t.Error("UInt32ToFloat32 Failed!")
	} else {
		t.Log("UInt32ToFloat32 OK")
	}
}

func Test_UInt32ToInt(t *testing.T) {
	b := NewUInt32()
	b++
	if b.ToInt() != 1 {
		t.Error("UInt32ToInt Failed!")
	} else {
		t.Log("UInt32ToInt OK")
	}
}

func Test_UInt32ToUInt(t *testing.T) {
	b := NewUInt32()
	b++
	if b.ToUInt() != 1 {
		t.Error("UInt32ToUInt Failed!")
	} else {
		t.Log("UInt32ToUInt OK")
	}
}

func Test_UInt32ToFloat(t *testing.T) {
	b := NewUInt32()
	b++
	if b.ToFloat() != 1 {
		t.Error("UInt32ToFloat Failed!")
	} else {
		t.Log("UInt32ToFloat OK")
	}
}

func Test_UInt32ToBool(t *testing.T) {
	b := NewUInt32()
	if b.ToBool() != false {
		t.Error("UInt32ToBool Failed!")
	} else {
		t.Log("UInt32ToBool OK")
	}
	b = 1
	if b.ToBool() != true {
		t.Error("UInt32ToBool Failed!")
	} else {
		t.Log("UInt32ToBool OK")
	}
}

func Test_UInt32ToJSON(t *testing.T) {
	b := NewUInt32()
	if b.ToJSON().ToString() != "0" {
		t.Error("UInt32ToJSON Failed!")
	} else {
		t.Log("UInt32ToJSON OK")
	}
}

func Test_UInt32ToString(t *testing.T) {
	b := NewUInt32()
	if b.ToString() != "0" {
		t.Error("UInt32ToJSON Failed!")
	} else {
		t.Log("UInt32ToJSON OK")
	}
}

func Test_UInt32ToJSONString(t *testing.T) {
	b := NewUInt32()
	if s, _ := b.ToJSONString(); s != "0" {
		t.Error("ToJSONString Failed!")
	} else {
		t.Log("ToJSONString OK")
	}
}
