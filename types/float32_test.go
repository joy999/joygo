package types

import (
	"testing"
)

func Test_Float32ToNative(t *testing.T) {
	b := NewFloat32()
	b++
	if b.ToNative() != 1 {
		t.Error("Float32ToNative Failed!")
	} else {
		t.Log("Float32ToNative OK")
	}
	b = 1.2
	if b.ToNative() != 1.2 {
		t.Error("Float32ToNative Failed!")
	} else {
		t.Log("Float32ToNative OK")
	}
}

func Test_Float32ToInt64(t *testing.T) {
	b := NewFloat32()
	b++
	if b.ToInt64() != 1 {
		t.Error("Float32ToInt64 Failed!")
	} else {
		t.Log("Float32ToInt64 OK")
	}
	b = 1.2
	if b.ToInt64() != 1 {
		t.Error("Float32ToInt64 Failed!")
	} else {
		t.Log("Float32ToInt64 OK")
	}
}

func Test_Float32ToInt32(t *testing.T) {
	b := NewFloat32()
	b++
	if b.ToInt32() != 1 {
		t.Error("Float32ToInt32 Failed!")
	} else {
		t.Log("Float32ToInt32 OK")
	}
	b = 1.2
	if b.ToInt32() != 1 {
		t.Error("Float32ToInt32 Failed!")
	} else {
		t.Log("Float32ToInt32 OK")
	}
}

func Test_Float32ToInt16(t *testing.T) {
	b := NewFloat32()
	b++
	if b.ToInt16() != 1 {
		t.Error("Float32ToInt16 Failed!")
	} else {
		t.Log("Float32ToInt16 OK")
	}
	b = 1.2
	if b.ToInt16() != 1 {
		t.Error("Float32ToInt16 Failed!")
	} else {
		t.Log("Float32ToInt16 OK")
	}
}

func Test_Float32ToInt8(t *testing.T) {
	b := NewFloat32()
	b++
	if b.ToInt8() != 1 {
		t.Error("Float32ToInt8 Failed!")
	} else {
		t.Log("Float32ToInt8 OK")
	}
	b = 1.2
	if b.ToInt8() != 1 {
		t.Error("Float32ToInt8 Failed!")
	} else {
		t.Log("Float32ToInt8 OK")
	}
}

func Test_Float32ToUInt64(t *testing.T) {
	b := NewFloat32()
	b++
	if b.ToUInt64() != 1 {
		t.Error("Float32ToUInt64 Failed!")
	} else {
		t.Log("Float32ToUInt64 OK")
	}
	b = 1.2
	if b.ToUInt64() != 1 {
		t.Error("Float32ToUInt64 Failed!")
	} else {
		t.Log("Float32ToUInt64 OK")
	}
	b = -1
	if b.ToUInt64() != 0 {
		t.Error("Float64ToUInt64 Failed!")
	} else {
		t.Log("Float64ToUInt64 OK")
	}
	b = -1.2
	if b.ToUInt64() != 0 {
		t.Error("Float64ToUInt64 Failed!")
	} else {
		t.Log("Float64ToUInt64 OK")
	}
}

func Test_Float32ToUInt32(t *testing.T) {
	b := NewFloat32()
	b++
	if b.ToUInt32() != 1 {
		t.Error("Float32ToUInt32 Failed!")
	} else {
		t.Log("Float32ToUInt32 OK")
	}
	b = 1.2
	if b.ToUInt32() != 1 {
		t.Error("Float32ToUInt32 Failed!")
	} else {
		t.Log("Float32ToUInt32 OK")
	}
	b = -1
	if b.ToUInt32() != 0 {
		t.Error("Float64ToUInt32 Failed!")
	} else {
		t.Log("Float64ToUInt32 OK")
	}
	b = -1.2
	if b.ToUInt32() != 0 {
		t.Error("Float64ToUInt32 Failed!")
	} else {
		t.Log("Float64ToUInt32 OK")
	}
}

func Test_Float32ToUInt16(t *testing.T) {
	b := NewFloat32()
	b++
	if b.ToUInt16() != 1 {
		t.Error("Float32ToUInt16 Failed!")
	} else {
		t.Log("Float32ToUInt16 OK")
	}
	b = 1.2
	if b.ToUInt16() != 1 {
		t.Error("Float32ToUInt16 Failed!")
	} else {
		t.Log("Float32ToUInt16 OK")
	}
	b = -1.2
	if b.ToUInt16() != 0 {
		t.Error("Float64ToUInt16 Failed!")
	} else {
		t.Log("Float64ToUInt16 OK")
	}
}

func Test_Float32ToUInt8(t *testing.T) {
	b := NewFloat32()
	b++
	if b.ToUInt8() != 1 {
		t.Error("Float32ToUInt8 Failed!")
	} else {
		t.Log("Float32ToUInt8 OK")
	}
	b = 1.2
	if b.ToUInt8() != 1 {
		t.Error("Float32ToUInt8 Failed!")
	} else {
		t.Log("Float32ToUInt8 OK")
	}
	b = -1.2
	if b.ToUInt8() != 0 {
		t.Error("Float64ToUInt8 Failed!")
	} else {
		t.Log("Float64ToUInt8 OK")
	}
}

func Test_Float32ToFloat64(t *testing.T) {
	b := NewFloat32()
	b++
	if b.ToFloat64() != 1 {
		t.Error("Float32ToFloat64 Failed!")
	} else {
		t.Log("Float32ToFloat64 OK")
	}
	b = 1.2
	if b.ToFloat64() != 1.2 {
		t.Error("Float32ToFloat64 Failed!")
	} else {
		t.Log("Float32ToFloat64 OK")
	}
	b = -1.2
	if b.ToFloat64() != -1.2 {
		t.Error("Float64ToFloat64 Failed!")
	} else {
		t.Log("Float64ToFloat64 OK")
	}
}

func Test_Float32ToFloat32(t *testing.T) {
	b := NewFloat32()
	b++
	if b.ToFloat32() != 1 {
		t.Error("Float32ToFloat32 Failed!")
	} else {
		t.Log("Float32ToFloat32 OK")
	}
	b = 1.2
	if b.ToFloat32() != 1.2 {
		t.Error("Float32ToFloat32 Failed!")
	} else {
		t.Log("Float32ToFloat32 OK")
	}
	b = -1.2
	if b.ToFloat32() != -1.2 {
		t.Error("Float64ToFloat32 Failed!")
	} else {
		t.Log("Float64ToFloat32 OK")
	}
}

func Test_Float32ToInt(t *testing.T) {
	b := NewFloat32()
	b++
	if b.ToInt() != 1 {
		t.Error("Float32ToInt Failed!")
	} else {
		t.Log("Float32ToInt OK")
	}
	b = 1.2
	if b.ToInt() != 1 {
		t.Error("Float32ToInt Failed!")
	} else {
		t.Log("Float32ToInt OK")
	}
	b = -1.2
	if b.ToInt() != -1 {
		t.Error("Float64ToInt Failed!")
	} else {
		t.Log("Float64ToInt OK")
	}
}

func Test_Float32ToUInt(t *testing.T) {
	b := NewFloat32()
	b++
	if b.ToUInt() != 1 {
		t.Error("Float32ToUInt Failed!")
	} else {
		t.Log("Float32ToUInt OK")
	}
	b = 1.2
	if b.ToUInt() != 1 {
		t.Error("Float32ToUInt Failed!")
	} else {
		t.Log("Float32ToUInt OK")
	}
	b = -1.2
	if b.ToUInt() != 0 {
		t.Error("Float64ToUInt Failed!")
	} else {
		t.Log("Float64ToUInt OK")
	}
}

func Test_Float32ToFloat(t *testing.T) {
	b := NewFloat32()
	b++
	if b.ToFloat() != 1 {
		t.Error("Float32ToFloat Failed!")
	} else {
		t.Log("Float32ToFloat OK")
	}
	b = 1.2
	if b.ToFloat() != 1.2 {
		t.Errorf("Float32ToFloat Failed! %v", b.ToFloat())
	} else {
		t.Log("Float32ToFloat OK")
	}
	b = -1.2
	if b.ToFloat() != -1.2 {
		t.Error("Float64ToFloat Failed!")
	} else {
		t.Log("Float64ToFloat OK")
	}
}

func Test_Float32ToBool(t *testing.T) {
	b := NewFloat32()
	if b.ToBool() != false {
		t.Error("Float32ToBool Failed!")
	} else {
		t.Log("Float32ToBool OK")
	}
	b = 1.2
	if b.ToBool() != true {
		t.Error("Float32ToBool Failed!")
	} else {
		t.Log("Float32ToBool OK")
	}
}

func Test_Float32ToJSON(t *testing.T) {
	b := NewFloat32()
	if b.ToJSON().ToString() != "0" {
		t.Error("Float32ToJSON Failed!")
	} else {
		t.Log("Float32ToJSON OK")
	}
	b = 1.2
	if b.ToJSON().ToString() != "1.2" {
		t.Errorf("Float32ToJSON Failed! %v", b.ToJSON().ToString())
	} else {
		t.Log("Float32ToJSON OK")
	}
}

func Test_Float32ToString(t *testing.T) {
	b := NewFloat32()
	if b.ToString() != "0" {
		t.Error("Float32ToJSON Failed!")
	} else {
		t.Log("Float32ToJSON OK")
	}
	b = 1.2
	if b.ToString() != "1.2" {
		t.Errorf("Float32ToJSON Failed! %v", b.ToString())
	} else {
		t.Log("Float32ToJSON OK")
	}
	b = -1.2
	if b.ToString() != "-1.2" {
		t.Errorf("Float64ToJSON Failed! %v", b.ToString())
	} else {
		t.Log("Float64ToJSON OK")
	}
}

func Test_Float32ToJSONString(t *testing.T) {
	b := NewFloat32()
	if s, _ := b.ToJSONString(); s != "0" {
		t.Error("ToJSONString Failed!")
	} else {
		t.Log("ToJSONString OK")
	}
	b = 1.2
	if s, _ := b.ToJSONString(); s != "1.2" {
		t.Errorf("ToJSONString Failed! %v", s)
	} else {
		t.Log("ToJSONString OK")
	}
	b = -1.2
	if s, _ := b.ToJSONString(); s != "-1.2" {
		t.Errorf("ToJSONString Failed! %v", s)
	} else {
		t.Log("ToJSONString OK")
	}
}
