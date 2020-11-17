package types

import (
	"testing"
)

func Test_Int64ToNative(t *testing.T) {
	b := NewInt64()
	b = 1
	if b.ToNative() != 1 {
		t.Error("Int64ToNative Failed!")
	} else {
		t.Log("Int64ToNative OK")
	}
	b = -1
	if b.ToNative() != -1 {
		t.Error("Int64ToNative Failed!")
	} else {
		t.Log("Int64ToNative OK")
	}
}

func Test_Int64ToInt64(t *testing.T) {
	b := NewInt64()
	b = 1
	if b.ToInt64() != 1 {
		t.Error("Int64ToInt64 Failed!")
	} else {
		t.Log("Int64ToInt64 OK")
	}
	b = -1
	if b.ToInt64() != -1 {
		t.Error("Int64ToInt64 Failed!")
	} else {
		t.Log("Int64ToInt64 OK")
	}
}

func Test_Int64ToInt32(t *testing.T) {
	b := NewInt64()
	b = 1
	if b.ToInt32() != 1 {
		t.Error("Int64ToInt32 Failed!")
	} else {
		t.Log("Int64ToInt32 OK")
	}
	b = -1
	if b.ToInt32() != -1 {
		t.Error("Int64ToInt32 Failed!")
	} else {
		t.Log("Int64ToInt32 OK")
	}
}

func Test_Int64ToInt16(t *testing.T) {
	b := NewInt64()
	b = 1
	if b.ToInt16() != 1 {
		t.Error("Int64ToInt16 Failed!")
	} else {
		t.Log("Int64ToInt16 OK")
	}
	b = -1
	if b.ToInt16() != -1 {
		t.Error("Int64ToInt16 Failed!")
	} else {
		t.Log("Int64ToInt16 OK")
	}
}

func Test_Int64ToInt8(t *testing.T) {
	b := NewInt64()
	b = 1
	if b.ToInt8() != 1 {
		t.Error("Int64ToInt8 Failed!")
	} else {
		t.Log("Int64ToInt8 OK")
	}
	b = -1
	if b.ToInt8() != -1 {
		t.Error("Int64ToInt8 Failed!")
	} else {
		t.Log("Int64ToInt8 OK")
	}
}

func Test_Int64ToUInt64(t *testing.T) {
	b := NewInt64()
	b = 1
	if b.ToUInt64() != 1 {
		t.Error("Int64ToUInt64 Failed!")
	} else {
		t.Log("Int64ToUInt64 OK")
	}
	b = -1
	if b.ToUInt64() != 0 {
		t.Errorf("Int64ToUInt64 Failed! %v", b.ToUInt64())
	} else {
		t.Log("Int64ToUInt64 OK")
	}
}

func Test_Int64ToUInt32(t *testing.T) {
	b := NewInt64()
	b++
	if b.ToUInt32() != 1 {
		t.Error("Int64ToUInt32 Failed!")
	} else {
		t.Log("Int64ToUInt32 OK")
	}
	b = -1
	if b.ToUInt32() != 0 {
		t.Error("Int64ToUInt32 Failed!")
	} else {
		t.Log("Int64ToUInt32 OK")
	}
}

func Test_Int64ToUInt16(t *testing.T) {
	b := NewInt64()
	b++
	if b.ToUInt16() != 1 {
		t.Error("Int64ToUInt16 Failed!")
	} else {
		t.Log("Int64ToUInt16 OK")
	}
	b = -1
	if b.ToUInt16() != 0 {
		t.Error("Int64ToUInt16 Failed!")
	} else {
		t.Log("Int64ToUInt16 OK")
	}
}

func Test_Int64ToUInt8(t *testing.T) {
	b := NewInt64()
	b++
	if b.ToUInt8() != 1 {
		t.Error("Int64ToUInt8 Failed!")
	} else {
		t.Log("Int64ToUInt8 OK")
	}
	b = -1
	if b.ToUInt8() != 0 {
		t.Error("Int64ToUInt8 Failed!")
	} else {
		t.Log("Int64ToUInt8 OK")
	}
}

func Test_Int64ToFloat64(t *testing.T) {
	b := NewInt64()
	b++
	if b.ToFloat64() != 1 {
		t.Error("Int64ToFloat64 Failed!")
	} else {
		t.Log("Int64ToFloat64 OK")
	}
	b = -1
	if b.ToFloat64() != -1 {
		t.Error("Int64ToFloat64 Failed!")
	} else {
		t.Log("Int64ToFloat64 OK")
	}
}

func Test_Int64ToFloat32(t *testing.T) {
	b := NewInt64()
	b++
	if b.ToFloat32() != 1 {
		t.Error("Int64ToFloat32 Failed!")
	} else {
		t.Log("Int64ToFloat32 OK")
	}
	b = -1
	if b.ToFloat32() != -1 {
		t.Error("Int64ToFloat32 Failed!")
	} else {
		t.Log("Int64ToFloat32 OK")
	}
}

func Test_Int64ToInt(t *testing.T) {
	b := NewInt64()
	b++
	if b.ToInt() != 1 {
		t.Error("Int64ToInt Failed!")
	} else {
		t.Log("Int64ToInt OK")
	}
	b = -1
	if b.ToInt() != -1 {
		t.Error("Int64ToInt Failed!")
	} else {
		t.Log("Int64ToInt OK")
	}
}

func Test_Int64ToUInt(t *testing.T) {
	b := NewInt64()
	b++
	if b.ToUInt() != 1 {
		t.Error("Int64ToUInt Failed!")
	} else {
		t.Log("Int64ToUInt OK")
	}
	b = -1
	if b.ToUInt() != 0 {
		t.Error("Int64ToUInt Failed!")
	} else {
		t.Log("Int64ToUInt OK")
	}
}

func Test_Int64ToFloat(t *testing.T) {
	b := NewInt64()
	b++
	if b.ToFloat() != 1 {
		t.Error("Int64ToFloat Failed!")
	} else {
		t.Log("Int64ToFloat OK")
	}
	b = -1
	if b.ToFloat() != -1 {
		t.Error("Int64ToFloat Failed!")
	} else {
		t.Log("Int64ToFloat OK")
	}
}

func Test_Int64ToBool(t *testing.T) {
	b := NewInt64()
	if b.ToBool() != false {
		t.Error("Int64ToBool Failed!")
	} else {
		t.Log("Int64ToBool OK")
	}
	b = 1
	if b.ToBool() != true {
		t.Error("Int64ToBool Failed!")
	} else {
		t.Log("Int64ToBool OK")
	}
}

func Test_Int64ToJSON(t *testing.T) {
	b := NewInt64()
	if b.ToJSON().ToString() != "0" {
		t.Error("Int64ToJSON Failed!")
	} else {
		t.Log("Int64ToJSON OK")
	}
	b = 1
	if b.ToJSON().ToString() != "1" {
		t.Errorf("Int64ToJSON Failed! %v", b.ToJSON().ToString())
	} else {
		t.Log("Int64ToJSON OK")
	}
}

func Test_Int64ToString(t *testing.T) {
	b := NewInt64()
	if b.ToString() != "0" {
		t.Error("Int64ToJSON Failed!")
	} else {
		t.Log("Int64ToJSON OK")
	}
	b = -1
	if b.ToString() != "-1" {
		t.Errorf("Int64ToJSON Failed! %v", b.ToString())
	} else {
		t.Log("Int64ToJSON OK")
	}
}

func Test_Int64ToJSONString(t *testing.T) {
	b := NewInt64()
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
