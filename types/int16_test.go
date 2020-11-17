package types

import (
	"testing"
)

func Test_Int16ToNative(t *testing.T) {
	b := NewInt16()
	b = 1
	if b.ToNative() != 1 {
		t.Error("Int16ToNative Failed!")
	} else {
		t.Log("Int16ToNative OK")
	}
	b = -1
	if b.ToNative() != -1 {
		t.Error("Int16ToNative Failed!")
	} else {
		t.Log("Int16ToNative OK")
	}
}

func Test_Int16ToInt64(t *testing.T) {
	b := NewInt16()
	b = 1
	if b.ToInt64() != 1 {
		t.Error("Int16ToInt64 Failed!")
	} else {
		t.Log("Int16ToInt64 OK")
	}
	b = -1
	if b.ToInt64() != -1 {
		t.Error("Int16ToInt64 Failed!")
	} else {
		t.Log("Int16ToInt64 OK")
	}
}

func Test_Int16ToInt32(t *testing.T) {
	b := NewInt16()
	b = 1
	if b.ToInt32() != 1 {
		t.Error("Int16ToInt32 Failed!")
	} else {
		t.Log("Int16ToInt32 OK")
	}
	b = -1
	if b.ToInt32() != -1 {
		t.Error("Int16ToInt32 Failed!")
	} else {
		t.Log("Int16ToInt32 OK")
	}
}

func Test_Int16ToInt16(t *testing.T) {
	b := NewInt16()
	b = 1
	if b.ToInt16() != 1 {
		t.Error("Int16ToInt16 Failed!")
	} else {
		t.Log("Int16ToInt16 OK")
	}
	b = -1
	if b.ToInt16() != -1 {
		t.Error("Int16ToInt16 Failed!")
	} else {
		t.Log("Int16ToInt16 OK")
	}
}

func Test_Int16ToInt8(t *testing.T) {
	b := NewInt16()
	b = 1
	if b.ToInt8() != 1 {
		t.Error("Int16ToInt8 Failed!")
	} else {
		t.Log("Int16ToInt8 OK")
	}
	b = -1
	if b.ToInt8() != -1 {
		t.Error("Int16ToInt8 Failed!")
	} else {
		t.Log("Int16ToInt8 OK")
	}
}

func Test_Int16ToUInt64(t *testing.T) {
	b := NewInt16()
	b = 1
	if b.ToUInt64() != 1 {
		t.Error("Int16ToUInt64 Failed!")
	} else {
		t.Log("Int16ToUInt64 OK")
	}
	b = -1
	if b.ToUInt64() != 0 {
		t.Errorf("Int16ToUInt64 Failed! %v", b.ToUInt64())
	} else {
		t.Log("Int16ToUInt64 OK")
	}
}

func Test_Int16ToUInt32(t *testing.T) {
	b := NewInt16()
	b++
	if b.ToUInt32() != 1 {
		t.Error("Int16ToUInt32 Failed!")
	} else {
		t.Log("Int16ToUInt32 OK")
	}
	b = -1
	if b.ToUInt32() != 0 {
		t.Error("Int16ToUInt32 Failed!")
	} else {
		t.Log("Int16ToUInt32 OK")
	}
}

func Test_Int16ToUInt16(t *testing.T) {
	b := NewInt16()
	b++
	if b.ToUInt16() != 1 {
		t.Error("Int16ToUInt16 Failed!")
	} else {
		t.Log("Int16ToUInt16 OK")
	}
	b = -1
	if b.ToUInt16() != 0 {
		t.Error("Int16ToUInt16 Failed!")
	} else {
		t.Log("Int16ToUInt16 OK")
	}
}

func Test_Int16ToUInt8(t *testing.T) {
	b := NewInt16()
	b++
	if b.ToUInt8() != 1 {
		t.Error("Int16ToUInt8 Failed!")
	} else {
		t.Log("Int16ToUInt8 OK")
	}
	b = -1
	if b.ToUInt8() != 0 {
		t.Error("Int16ToUInt8 Failed!")
	} else {
		t.Log("Int16ToUInt8 OK")
	}
}

func Test_Int16ToFloat64(t *testing.T) {
	b := NewInt16()
	b++
	if b.ToFloat64() != 1 {
		t.Error("Int16ToFloat64 Failed!")
	} else {
		t.Log("Int16ToFloat64 OK")
	}
	b = -1
	if b.ToFloat64() != -1 {
		t.Error("Int16ToFloat64 Failed!")
	} else {
		t.Log("Int16ToFloat64 OK")
	}
}

func Test_Int16ToFloat32(t *testing.T) {
	b := NewInt16()
	b++
	if b.ToFloat32() != 1 {
		t.Error("Int16ToFloat32 Failed!")
	} else {
		t.Log("Int16ToFloat32 OK")
	}
	b = -1
	if b.ToFloat32() != -1 {
		t.Error("Int16ToFloat32 Failed!")
	} else {
		t.Log("Int16ToFloat32 OK")
	}
}

func Test_Int16ToInt(t *testing.T) {
	b := NewInt16()
	b++
	if b.ToInt() != 1 {
		t.Error("Int16ToInt Failed!")
	} else {
		t.Log("Int16ToInt OK")
	}
	b = -1
	if b.ToInt() != -1 {
		t.Error("Int16ToInt Failed!")
	} else {
		t.Log("Int16ToInt OK")
	}
}

func Test_Int16ToUInt(t *testing.T) {
	b := NewInt16()
	b++
	if b.ToUInt() != 1 {
		t.Error("Int16ToUInt Failed!")
	} else {
		t.Log("Int16ToUInt OK")
	}
	b = -1
	if b.ToUInt() != 0 {
		t.Error("Int16ToUInt Failed!")
	} else {
		t.Log("Int16ToUInt OK")
	}
}

func Test_Int16ToFloat(t *testing.T) {
	b := NewInt16()
	b++
	if b.ToFloat() != 1 {
		t.Error("Int16ToFloat Failed!")
	} else {
		t.Log("Int16ToFloat OK")
	}
	b = -1
	if b.ToFloat() != -1 {
		t.Error("Int16ToFloat Failed!")
	} else {
		t.Log("Int16ToFloat OK")
	}
}

func Test_Int16ToBool(t *testing.T) {
	b := NewInt16()
	if b.ToBool() != false {
		t.Error("Int16ToBool Failed!")
	} else {
		t.Log("Int16ToBool OK")
	}
	b = 1
	if b.ToBool() != true {
		t.Error("Int16ToBool Failed!")
	} else {
		t.Log("Int16ToBool OK")
	}
}

func Test_Int16ToJSON(t *testing.T) {
	b := NewInt16()
	if b.ToJSON().ToString() != "0" {
		t.Error("Int16ToJSON Failed!")
	} else {
		t.Log("Int16ToJSON OK")
	}
	b = 1
	if b.ToJSON().ToString() != "1" {
		t.Errorf("Int16ToJSON Failed! %v", b.ToJSON().ToString())
	} else {
		t.Log("Int16ToJSON OK")
	}
}

func Test_Int16ToString(t *testing.T) {
	b := NewInt16()
	if b.ToString() != "0" {
		t.Error("Int16ToJSON Failed!")
	} else {
		t.Log("Int16ToJSON OK")
	}
	b = -1
	if b.ToString() != "-1" {
		t.Errorf("Int16ToJSON Failed! %v", b.ToString())
	} else {
		t.Log("Int16ToJSON OK")
	}
}

func Test_Int16ToJSONString(t *testing.T) {
	b := NewInt16()
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
