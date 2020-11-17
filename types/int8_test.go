package types

import (
	"testing"
)

func Test_Int8ToNative(t *testing.T) {
	b := NewInt8()
	b = 1
	if b.ToNative() != 1 {
		t.Error("Int8ToNative Failed!")
	} else {
		t.Log("Int8ToNative OK")
	}
	b = -1
	if b.ToNative() != -1 {
		t.Error("Int8ToNative Failed!")
	} else {
		t.Log("Int8ToNative OK")
	}
}

func Test_Int8ToInt64(t *testing.T) {
	b := NewInt8()
	b = 1
	if b.ToInt64() != 1 {
		t.Error("Int8ToInt64 Failed!")
	} else {
		t.Log("Int8ToInt64 OK")
	}
	b = -1
	if b.ToInt64() != -1 {
		t.Error("Int8ToInt64 Failed!")
	} else {
		t.Log("Int8ToInt64 OK")
	}
}

func Test_Int8ToInt32(t *testing.T) {
	b := NewInt8()
	b = 1
	if b.ToInt32() != 1 {
		t.Error("Int8ToInt32 Failed!")
	} else {
		t.Log("Int8ToInt32 OK")
	}
	b = -1
	if b.ToInt32() != -1 {
		t.Error("Int8ToInt32 Failed!")
	} else {
		t.Log("Int8ToInt32 OK")
	}
}

func Test_Int8ToInt16(t *testing.T) {
	b := NewInt8()
	b = 1
	if b.ToInt16() != 1 {
		t.Error("Int8ToInt16 Failed!")
	} else {
		t.Log("Int8ToInt16 OK")
	}
	b = -1
	if b.ToInt16() != -1 {
		t.Error("Int8ToInt16 Failed!")
	} else {
		t.Log("Int8ToInt16 OK")
	}
}

func Test_Int8ToInt8(t *testing.T) {
	b := NewInt8()
	b = 1
	if b.ToInt8() != 1 {
		t.Error("Int8ToInt8 Failed!")
	} else {
		t.Log("Int8ToInt8 OK")
	}
	b = -1
	if b.ToInt8() != -1 {
		t.Error("Int8ToInt8 Failed!")
	} else {
		t.Log("Int8ToInt8 OK")
	}
}

func Test_Int8ToUInt64(t *testing.T) {
	b := NewInt8()
	b = 1
	if b.ToUInt64() != 1 {
		t.Error("Int8ToUInt64 Failed!")
	} else {
		t.Log("Int8ToUInt64 OK")
	}
	b = -1
	if b.ToUInt64() != 0 {
		t.Errorf("Int8ToUInt64 Failed! %v", b.ToUInt64())
	} else {
		t.Log("Int8ToUInt64 OK")
	}
}

func Test_Int8ToUInt32(t *testing.T) {
	b := NewInt8()
	b++
	if b.ToUInt32() != 1 {
		t.Error("Int8ToUInt32 Failed!")
	} else {
		t.Log("Int8ToUInt32 OK")
	}
	b = -1
	if b.ToUInt32() != 0 {
		t.Error("Int8ToUInt32 Failed!")
	} else {
		t.Log("Int8ToUInt32 OK")
	}
}

func Test_Int8ToUInt16(t *testing.T) {
	b := NewInt8()
	b++
	if b.ToUInt16() != 1 {
		t.Error("Int8ToUInt16 Failed!")
	} else {
		t.Log("Int8ToUInt16 OK")
	}
	b = -1
	if b.ToUInt16() != 0 {
		t.Error("Int8ToUInt16 Failed!")
	} else {
		t.Log("Int8ToUInt16 OK")
	}
}

func Test_Int8ToUInt8(t *testing.T) {
	b := NewInt8()
	b++
	if b.ToUInt8() != 1 {
		t.Error("Int8ToUInt8 Failed!")
	} else {
		t.Log("Int8ToUInt8 OK")
	}
	b = -1
	if b.ToUInt8() != 0 {
		t.Error("Int8ToUInt8 Failed!")
	} else {
		t.Log("Int8ToUInt8 OK")
	}
}

func Test_Int8ToFloat64(t *testing.T) {
	b := NewInt8()
	b++
	if b.ToFloat64() != 1 {
		t.Error("Int8ToFloat64 Failed!")
	} else {
		t.Log("Int8ToFloat64 OK")
	}
	b = -1
	if b.ToFloat64() != -1 {
		t.Error("Int8ToFloat64 Failed!")
	} else {
		t.Log("Int8ToFloat64 OK")
	}
}

func Test_Int8ToFloat32(t *testing.T) {
	b := NewInt8()
	b++
	if b.ToFloat32() != 1 {
		t.Error("Int8ToFloat32 Failed!")
	} else {
		t.Log("Int8ToFloat32 OK")
	}
	b = -1
	if b.ToFloat32() != -1 {
		t.Error("Int8ToFloat32 Failed!")
	} else {
		t.Log("Int8ToFloat32 OK")
	}
}

func Test_Int8ToInt(t *testing.T) {
	b := NewInt8()
	b++
	if b.ToInt() != 1 {
		t.Error("Int8ToInt Failed!")
	} else {
		t.Log("Int8ToInt OK")
	}
	b = -1
	if b.ToInt() != -1 {
		t.Error("Int8ToInt Failed!")
	} else {
		t.Log("Int8ToInt OK")
	}
}

func Test_Int8ToUInt(t *testing.T) {
	b := NewInt8()
	b++
	if b.ToUInt() != 1 {
		t.Error("Int8ToUInt Failed!")
	} else {
		t.Log("Int8ToUInt OK")
	}
	b = -1
	if b.ToUInt() != 0 {
		t.Error("Int8ToUInt Failed!")
	} else {
		t.Log("Int8ToUInt OK")
	}
}

func Test_Int8ToFloat(t *testing.T) {
	b := NewInt8()
	b++
	if b.ToFloat() != 1 {
		t.Error("Int8ToFloat Failed!")
	} else {
		t.Log("Int8ToFloat OK")
	}
	b = -1
	if b.ToFloat() != -1 {
		t.Error("Int8ToFloat Failed!")
	} else {
		t.Log("Int8ToFloat OK")
	}
}

func Test_Int8ToBool(t *testing.T) {
	b := NewInt8()
	if b.ToBool() != false {
		t.Error("Int8ToBool Failed!")
	} else {
		t.Log("Int8ToBool OK")
	}
	b = 1
	if b.ToBool() != true {
		t.Error("Int8ToBool Failed!")
	} else {
		t.Log("Int8ToBool OK")
	}
}

func Test_Int8ToJSON(t *testing.T) {
	b := NewInt8()
	if b.ToJSON().ToString() != "0" {
		t.Error("Int8ToJSON Failed!")
	} else {
		t.Log("Int8ToJSON OK")
	}
	b = 1
	if b.ToJSON().ToString() != "1" {
		t.Errorf("Int8ToJSON Failed! %v", b.ToJSON().ToString())
	} else {
		t.Log("Int8ToJSON OK")
	}
}

func Test_Int8ToString(t *testing.T) {
	b := NewInt8()
	if b.ToString() != "0" {
		t.Error("Int8ToJSON Failed!")
	} else {
		t.Log("Int8ToJSON OK")
	}
	b = -1
	if b.ToString() != "-1" {
		t.Errorf("Int8ToJSON Failed! %v", b.ToString())
	} else {
		t.Log("Int8ToJSON OK")
	}
}

func Test_Int8ToJSONString(t *testing.T) {
	b := NewInt8()
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
