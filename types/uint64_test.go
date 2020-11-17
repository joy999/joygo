package types

import (
	"testing"
)

func Test_UInt64ToNative(t *testing.T) {
	b := NewUInt64()
	b = 1
	if b.ToNative() != 1 {
		t.Error("UInt64ToNative Failed!")
	} else {
		t.Log("UInt64ToNative OK")
	}
}

func Test_UInt64ToInt64(t *testing.T) {
	b := NewUInt64()
	b = 1
	if b.ToInt64() != 1 {
		t.Error("UInt64ToInt64 Failed!")
	} else {
		t.Log("UInt64ToInt64 OK")
	}
}

func Test_UInt64ToInt32(t *testing.T) {
	b := NewUInt64()
	b = 1
	if b.ToInt32() != 1 {
		t.Error("UInt64ToInt32 Failed!")
	} else {
		t.Log("UInt64ToInt32 OK")
	}
}

func Test_UInt64ToInt16(t *testing.T) {
	b := NewUInt64()
	b = 1
	if b.ToInt16() != 1 {
		t.Error("UInt64ToInt16 Failed!")
	} else {
		t.Log("UInt64ToInt16 OK")
	}
}

func Test_UInt64ToInt8(t *testing.T) {
	b := NewUInt64()
	b = 1
	if b.ToInt8() != 1 {
		t.Error("UInt64ToInt8 Failed!")
	} else {
		t.Log("UInt64ToInt8 OK")
	}
}

func Test_UInt64ToUInt64(t *testing.T) {
	b := NewUInt64()
	b = 1
	if b.ToUInt64() != 1 {
		t.Error("UInt64ToUInt64 Failed!")
	} else {
		t.Log("UInt64ToUInt64 OK")
	}
}

func Test_UInt64ToUInt32(t *testing.T) {
	b := NewUInt64()
	b++
	if b.ToUInt32() != 1 {
		t.Error("UInt64ToUInt32 Failed!")
	} else {
		t.Log("UInt64ToUInt32 OK")
	}
}

func Test_UInt64ToUInt16(t *testing.T) {
	b := NewUInt64()
	b++
	if b.ToUInt16() != 1 {
		t.Error("UInt64ToUInt16 Failed!")
	} else {
		t.Log("UInt64ToUInt16 OK")
	}
}

func Test_UInt64ToUInt8(t *testing.T) {
	b := NewUInt64()
	b++
	if b.ToUInt8() != 1 {
		t.Error("UInt64ToUInt8 Failed!")
	} else {
		t.Log("UInt64ToUInt8 OK")
	}
}

func Test_UInt64ToFloat64(t *testing.T) {
	b := NewUInt64()
	b++
	if b.ToFloat64() != 1 {
		t.Error("UInt64ToFloat64 Failed!")
	} else {
		t.Log("UInt64ToFloat64 OK")
	}
}

func Test_UInt64ToFloat32(t *testing.T) {
	b := NewUInt64()
	b++
	if b.ToFloat32() != 1 {
		t.Error("UInt64ToFloat32 Failed!")
	} else {
		t.Log("UInt64ToFloat32 OK")
	}
}

func Test_UInt64ToInt(t *testing.T) {
	b := NewUInt64()
	b++
	if b.ToInt() != 1 {
		t.Error("UInt64ToInt Failed!")
	} else {
		t.Log("UInt64ToInt OK")
	}
}

func Test_UInt64ToUInt(t *testing.T) {
	b := NewUInt64()
	b++
	if b.ToUInt() != 1 {
		t.Error("UInt64ToUInt Failed!")
	} else {
		t.Log("UInt64ToUInt OK")
	}
}

func Test_UInt64ToFloat(t *testing.T) {
	b := NewUInt64()
	b++
	if b.ToFloat() != 1 {
		t.Error("UInt64ToFloat Failed!")
	} else {
		t.Log("UInt64ToFloat OK")
	}
}

func Test_UInt64ToBool(t *testing.T) {
	b := NewUInt64()
	if b.ToBool() != false {
		t.Error("UInt64ToBool Failed!")
	} else {
		t.Log("UInt64ToBool OK")
	}
	b = 222
	if b.ToBool() != true {
		t.Error("UInt64ToBool Failed!")
	} else {
		t.Log("UInt64ToBool OK")
	}
}

func Test_UInt64ToJSON(t *testing.T) {
	b := NewUInt64()
	if b.ToJSON().ToString() != "0" {
		t.Error("UInt64ToJSON Failed!")
	} else {
		t.Log("UInt64ToJSON OK")
	}
}

func Test_UInt64ToString(t *testing.T) {
	b := NewUInt64()
	if b.ToString() != "0" {
		t.Error("UInt64ToJSON Failed!")
	} else {
		t.Log("UInt64ToJSON OK")
	}
}

func Test_UInt64ToJSONString(t *testing.T) {
	b := NewUInt64()
	if s, _ := b.ToJSONString(); s != "0" {
		t.Error("ToJSONString Failed!")
	} else {
		t.Log("ToJSONString OK")
	}
}
