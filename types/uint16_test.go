package types

import (
	"testing"
)

func Test_UInt16ToNative(t *testing.T) {
	b := NewUInt16()
	b = 1
	if b.ToNative() != 1 {
		t.Error("UInt16ToNative Failed!")
	} else {
		t.Log("UInt16ToNative OK")
	}
}

func Test_UInt16ToInt64(t *testing.T) {
	b := NewUInt16()
	b = 1
	if b.ToInt64() != 1 {
		t.Error("UInt16ToInt64 Failed!")
	} else {
		t.Log("UInt16ToInt64 OK")
	}
}

func Test_UInt16ToInt32(t *testing.T) {
	b := NewUInt16()
	b = 1
	if b.ToInt32() != 1 {
		t.Error("UInt16ToInt32 Failed!")
	} else {
		t.Log("UInt16ToInt32 OK")
	}
}

func Test_UInt16ToInt16(t *testing.T) {
	b := NewUInt16()
	b = 1
	if b.ToInt16() != 1 {
		t.Error("UInt16ToInt16 Failed!")
	} else {
		t.Log("UInt16ToInt16 OK")
	}
}

func Test_UInt16ToInt8(t *testing.T) {
	b := NewUInt16()
	b = 1
	if b.ToInt8() != 1 {
		t.Error("UInt16ToInt8 Failed!")
	} else {
		t.Log("UInt16ToInt8 OK")
	}
}

func Test_UInt16ToUInt64(t *testing.T) {
	b := NewUInt16()
	b = 1
	if b.ToUInt64() != 1 {
		t.Error("UInt16ToUInt64 Failed!")
	} else {
		t.Log("UInt16ToUInt64 OK")
	}
}

func Test_UInt16ToUInt32(t *testing.T) {
	b := NewUInt16()
	b++
	if b.ToUInt32() != 1 {
		t.Error("UInt16ToUInt32 Failed!")
	} else {
		t.Log("UInt16ToUInt32 OK")
	}
}

func Test_UInt16ToUInt16(t *testing.T) {
	b := NewUInt16()
	b++
	if b.ToUInt16() != 1 {
		t.Error("UInt16ToUInt16 Failed!")
	} else {
		t.Log("UInt16ToUInt16 OK")
	}
}

func Test_UInt16ToUInt8(t *testing.T) {
	b := NewUInt16()
	b++
	if b.ToUInt8() != 1 {
		t.Error("UInt16ToUInt8 Failed!")
	} else {
		t.Log("UInt16ToUInt8 OK")
	}
}

func Test_UInt16ToFloat64(t *testing.T) {
	b := NewUInt16()
	b++
	if b.ToFloat64() != 1 {
		t.Error("UInt16ToFloat64 Failed!")
	} else {
		t.Log("UInt16ToFloat64 OK")
	}
}

func Test_UInt16ToFloat32(t *testing.T) {
	b := NewUInt16()
	b++
	if b.ToFloat32() != 1 {
		t.Error("UInt16ToFloat32 Failed!")
	} else {
		t.Log("UInt16ToFloat32 OK")
	}
}

func Test_UInt16ToInt(t *testing.T) {
	b := NewUInt16()
	b++
	if b.ToInt() != 1 {
		t.Error("UInt16ToInt Failed!")
	} else {
		t.Log("UInt16ToInt OK")
	}
}

func Test_UInt16ToUInt(t *testing.T) {
	b := NewUInt16()
	b++
	if b.ToUInt() != 1 {
		t.Error("UInt16ToUInt Failed!")
	} else {
		t.Log("UInt16ToUInt OK")
	}
}

func Test_UInt16ToFloat(t *testing.T) {
	b := NewUInt16()
	b++
	if b.ToFloat() != 1 {
		t.Error("UInt16ToFloat Failed!")
	} else {
		t.Log("UInt16ToFloat OK")
	}
}

func Test_UInt16ToBool(t *testing.T) {
	b := NewUInt16()
	if b.ToBool() != false {
		t.Error("UInt16ToBool Failed!")
	} else {
		t.Log("UInt16ToBool OK")
	}
	b = 222
	if b.ToBool() != true {
		t.Error("UInt16ToBool Failed!")
	} else {
		t.Log("UInt16ToBool OK")
	}
}

func Test_UInt16ToJSON(t *testing.T) {
	b := NewUInt16()
	if b.ToJSON().ToString() != "0" {
		t.Error("UInt16ToJSON Failed!")
	} else {
		t.Log("UInt16ToJSON OK")
	}
}

func Test_UInt16ToString(t *testing.T) {
	b := NewUInt16()
	if b.ToString() != "0" {
		t.Error("UInt16ToJSON Failed!")
	} else {
		t.Log("UInt16ToJSON OK")
	}
}

func Test_UInt16ToJSONString(t *testing.T) {
	b := NewUInt16()
	if s, _ := b.ToJSONString(); s != "0" {
		t.Error("ToJSONString Failed!")
	} else {
		t.Log("ToJSONString OK")
	}
}
