package types

import (
	"testing"
)

func Test_Int32ToNative(t *testing.T) {
	b := NewInt32()
	b = 1
	if b.ToNative() != 1 {
		t.Error("Int32ToNative Failed!")
	} else {
		t.Log("Int32ToNative OK")
	}
	b = -1
	if b.ToNative() != -1 {
		t.Error("Int32ToNative Failed!")
	} else {
		t.Log("Int32ToNative OK")
	}
}

func Test_Int32ToInt64(t *testing.T) {
	b := NewInt32()
	b = 1
	if b.ToInt64() != 1 {
		t.Error("Int32ToInt64 Failed!")
	} else {
		t.Log("Int32ToInt64 OK")
	}
	b = -1
	if b.ToInt64() != -1 {
		t.Error("Int32ToInt64 Failed!")
	} else {
		t.Log("Int32ToInt64 OK")
	}
}

func Test_Int32ToInt32(t *testing.T) {
	b := NewInt32()
	b = 1
	if b.ToInt32() != 1 {
		t.Error("Int32ToInt32 Failed!")
	} else {
		t.Log("Int32ToInt32 OK")
	}
	b = -1
	if b.ToInt32() != -1 {
		t.Error("Int32ToInt32 Failed!")
	} else {
		t.Log("Int32ToInt32 OK")
	}
}

func Test_Int32ToInt16(t *testing.T) {
	b := NewInt32()
	b = 1
	if b.ToInt16() != 1 {
		t.Error("Int32ToInt16 Failed!")
	} else {
		t.Log("Int32ToInt16 OK")
	}
	b = -1
	if b.ToInt16() != -1 {
		t.Error("Int32ToInt16 Failed!")
	} else {
		t.Log("Int32ToInt16 OK")
	}
}

func Test_Int32ToInt8(t *testing.T) {
	b := NewInt32()
	b = 1
	if b.ToInt8() != 1 {
		t.Error("Int32ToInt8 Failed!")
	} else {
		t.Log("Int32ToInt8 OK")
	}
	b = -1
	if b.ToInt8() != -1 {
		t.Error("Int32ToInt8 Failed!")
	} else {
		t.Log("Int32ToInt8 OK")
	}
}

func Test_Int32ToUInt64(t *testing.T) {
	b := NewInt32()
	b = 1
	if b.ToUInt64() != 1 {
		t.Error("Int32ToUInt64 Failed!")
	} else {
		t.Log("Int32ToUInt64 OK")
	}
	b = -1
	if b.ToUInt64() != 0 {
		t.Errorf("Int32ToUInt64 Failed! %v", b.ToUInt64())
	} else {
		t.Log("Int32ToUInt64 OK")
	}
}

func Test_Int32ToUInt32(t *testing.T) {
	b := NewInt32()
	b++
	if b.ToUInt32() != 1 {
		t.Error("Int32ToUInt32 Failed!")
	} else {
		t.Log("Int32ToUInt32 OK")
	}
	b = -1
	if b.ToUInt32() != 0 {
		t.Error("Int32ToUInt32 Failed!")
	} else {
		t.Log("Int32ToUInt32 OK")
	}
}

func Test_Int32ToUInt16(t *testing.T) {
	b := NewInt32()
	b++
	if b.ToUInt16() != 1 {
		t.Error("Int32ToUInt16 Failed!")
	} else {
		t.Log("Int32ToUInt16 OK")
	}
	b = -1
	if b.ToUInt16() != 0 {
		t.Error("Int32ToUInt16 Failed!")
	} else {
		t.Log("Int32ToUInt16 OK")
	}
}

func Test_Int32ToUInt8(t *testing.T) {
	b := NewInt32()
	b++
	if b.ToUInt8() != 1 {
		t.Error("Int32ToUInt8 Failed!")
	} else {
		t.Log("Int32ToUInt8 OK")
	}
	b = -1
	if b.ToUInt8() != 0 {
		t.Error("Int32ToUInt8 Failed!")
	} else {
		t.Log("Int32ToUInt8 OK")
	}
}

func Test_Int32ToFloat64(t *testing.T) {
	b := NewInt32()
	b++
	if b.ToFloat64() != 1 {
		t.Error("Int32ToFloat64 Failed!")
	} else {
		t.Log("Int32ToFloat64 OK")
	}
	b = -1
	if b.ToFloat64() != -1 {
		t.Error("Int32ToFloat64 Failed!")
	} else {
		t.Log("Int32ToFloat64 OK")
	}
}

func Test_Int32ToFloat32(t *testing.T) {
	b := NewInt32()
	b++
	if b.ToFloat32() != 1 {
		t.Error("Int32ToFloat32 Failed!")
	} else {
		t.Log("Int32ToFloat32 OK")
	}
	b = -1
	if b.ToFloat32() != -1 {
		t.Error("Int32ToFloat32 Failed!")
	} else {
		t.Log("Int32ToFloat32 OK")
	}
}

func Test_Int32ToInt(t *testing.T) {
	b := NewInt32()
	b++
	if b.ToInt() != 1 {
		t.Error("Int32ToInt Failed!")
	} else {
		t.Log("Int32ToInt OK")
	}
	b = -1
	if b.ToInt() != -1 {
		t.Error("Int32ToInt Failed!")
	} else {
		t.Log("Int32ToInt OK")
	}
}

func Test_Int32ToUInt(t *testing.T) {
	b := NewInt32()
	b++
	if b.ToUInt() != 1 {
		t.Error("Int32ToUInt Failed!")
	} else {
		t.Log("Int32ToUInt OK")
	}
	b = -1
	if b.ToUInt() != 0 {
		t.Error("Int32ToUInt Failed!")
	} else {
		t.Log("Int32ToUInt OK")
	}
}

func Test_Int32ToFloat(t *testing.T) {
	b := NewInt32()
	b++
	if b.ToFloat() != 1 {
		t.Error("Int32ToFloat Failed!")
	} else {
		t.Log("Int32ToFloat OK")
	}
	b = -1
	if b.ToFloat() != -1 {
		t.Error("Int32ToFloat Failed!")
	} else {
		t.Log("Int32ToFloat OK")
	}
}

func Test_Int32ToBool(t *testing.T) {
	b := NewInt32()
	if b.ToBool() != false {
		t.Error("Int32ToBool Failed!")
	} else {
		t.Log("Int32ToBool OK")
	}
	b = 1
	if b.ToBool() != true {
		t.Error("Int32ToBool Failed!")
	} else {
		t.Log("Int32ToBool OK")
	}
}

func Test_Int32ToJSON(t *testing.T) {
	b := NewInt32()
	if b.ToJSON().ToString() != "0" {
		t.Error("Int32ToJSON Failed!")
	} else {
		t.Log("Int32ToJSON OK")
	}
	b = 1
	if b.ToJSON().ToString() != "1" {
		t.Errorf("Int32ToJSON Failed! %v", b.ToJSON().ToString())
	} else {
		t.Log("Int32ToJSON OK")
	}
}

func Test_Int32ToString(t *testing.T) {
	b := NewInt32()
	if b.ToString() != "0" {
		t.Error("Int32ToJSON Failed!")
	} else {
		t.Log("Int32ToJSON OK")
	}
	b = -1
	if b.ToString() != "-1" {
		t.Errorf("Int32ToJSON Failed! %v", b.ToString())
	} else {
		t.Log("Int32ToJSON OK")
	}
}

func Test_Int32ToJSONString(t *testing.T) {
	b := NewInt32()
	if s, _ := b.ToJSONString(); s != "0" {
		t.Error("ToJSONString Failed!")
	} else {
		t.Log("ToJSONString OK")
	}
	b = -1
	if s, _ := b.ToJSONString(); s != "-1" {
		t.Errorf("ToJSONString Failed! %v", s)
	} else {
		t.Log("ToJSONString OK")
	}
}
