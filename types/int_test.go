package types

import (
	"testing"
)

func Test_IntToInt64(t *testing.T) {
	b := NewInt()
	b = 1
	if b.ToInt64() != 1 {
		t.Error("IntToInt64 Failed!")
	} else {
		t.Log("IntToInt64 OK")
	}
	b = -1
	if b.ToInt64() != -1 {
		t.Error("IntToInt64 Failed!")
	} else {
		t.Log("IntToInt64 OK")
	}
}

func Test_IntToInt32(t *testing.T) {
	b := NewInt()
	b = 1
	if b.ToInt32() != 1 {
		t.Error("IntToInt32 Failed!")
	} else {
		t.Log("IntToInt32 OK")
	}
	b = -1
	if b.ToInt32() != -1 {
		t.Error("IntToInt32 Failed!")
	} else {
		t.Log("IntToInt32 OK")
	}
}

func Test_IntToInt16(t *testing.T) {
	b := NewInt()
	b = 1
	if b.ToInt16() != 1 {
		t.Error("IntToInt16 Failed!")
	} else {
		t.Log("IntToInt16 OK")
	}
	b = -1
	if b.ToInt16() != -1 {
		t.Error("IntToInt16 Failed!")
	} else {
		t.Log("IntToInt16 OK")
	}
}

func Test_IntToInt8(t *testing.T) {
	b := NewInt()
	b = 1
	if b.ToInt8() != 1 {
		t.Error("IntToInt8 Failed!")
	} else {
		t.Log("IntToInt8 OK")
	}
	b = -1
	if b.ToInt8() != -1 {
		t.Error("IntToInt8 Failed!")
	} else {
		t.Log("IntToInt8 OK")
	}
}

func Test_IntToUInt64(t *testing.T) {
	b := NewInt()
	b = 1
	if b.ToUInt64() != 1 {
		t.Error("IntToUInt64 Failed!")
	} else {
		t.Log("IntToUInt64 OK")
	}
	b = -1
	if b.ToUInt64() != 0 {
		t.Errorf("IntToUInt64 Failed! %v", b.ToUInt64())
	} else {
		t.Log("IntToUInt64 OK")
	}
}

func Test_IntToUInt32(t *testing.T) {
	b := NewInt()
	b++
	if b.ToUInt32() != 1 {
		t.Error("IntToUInt32 Failed!")
	} else {
		t.Log("IntToUInt32 OK")
	}
	b = -1
	if b.ToUInt32() != 0 {
		t.Error("IntToUInt32 Failed!")
	} else {
		t.Log("IntToUInt32 OK")
	}
}

func Test_IntToUInt16(t *testing.T) {
	b := NewInt()
	b++
	if b.ToUInt16() != 1 {
		t.Error("IntToUInt16 Failed!")
	} else {
		t.Log("IntToUInt16 OK")
	}
	b = -1
	if b.ToUInt16() != 0 {
		t.Error("IntToUInt16 Failed!")
	} else {
		t.Log("IntToUInt16 OK")
	}
}

func Test_IntToUInt8(t *testing.T) {
	b := NewInt()
	b++
	if b.ToUInt8() != 1 {
		t.Error("IntToUInt8 Failed!")
	} else {
		t.Log("IntToUInt8 OK")
	}
	b = -1
	if b.ToUInt8() != 0 {
		t.Error("IntToUInt8 Failed!")
	} else {
		t.Log("IntToUInt8 OK")
	}
}

func Test_IntToFloat64(t *testing.T) {
	b := NewInt()
	b++
	if b.ToFloat64() != 1 {
		t.Error("IntToFloat64 Failed!")
	} else {
		t.Log("IntToFloat64 OK")
	}
	b = -1
	if b.ToFloat64() != -1 {
		t.Error("IntToFloat64 Failed!")
	} else {
		t.Log("IntToFloat64 OK")
	}
}

func Test_IntToFloat32(t *testing.T) {
	b := NewInt()
	b++
	if b.ToFloat32() != 1 {
		t.Error("IntToFloat32 Failed!")
	} else {
		t.Log("IntToFloat32 OK")
	}
	b = -1
	if b.ToFloat32() != -1 {
		t.Error("IntToFloat32 Failed!")
	} else {
		t.Log("IntToFloat32 OK")
	}
}

func Test_IntToInt(t *testing.T) {
	b := NewInt()
	b++
	if b.ToInt() != 1 {
		t.Error("IntToInt Failed!")
	} else {
		t.Log("IntToInt OK")
	}
	b = -1
	if b.ToInt() != -1 {
		t.Error("IntToInt Failed!")
	} else {
		t.Log("IntToInt OK")
	}
}

func Test_IntToUInt(t *testing.T) {
	b := NewInt()
	b++
	if b.ToUInt() != 1 {
		t.Error("IntToUInt Failed!")
	} else {
		t.Log("IntToUInt OK")
	}
	b = -1
	if b.ToUInt() != 0 {
		t.Error("IntToUInt Failed!")
	} else {
		t.Log("IntToUInt OK")
	}
}

func Test_IntToFloat(t *testing.T) {
	b := NewInt()
	b++
	if b.ToFloat() != 1 {
		t.Error("IntToFloat Failed!")
	} else {
		t.Log("IntToFloat OK")
	}
	b = -1
	if b.ToFloat() != -1 {
		t.Error("IntToFloat Failed!")
	} else {
		t.Log("IntToFloat OK")
	}
}

func Test_IntToBool(t *testing.T) {
	b := NewInt()
	if b.ToBool() != false {
		t.Error("IntToBool Failed!")
	} else {
		t.Log("IntToBool OK")
	}
	b = 1
	if b.ToBool() != true {
		t.Error("IntToBool Failed!")
	} else {
		t.Log("IntToBool OK")
	}
}

func Test_IntToJSON(t *testing.T) {
	b := NewInt()
	if b.ToJSON().ToString() != "0" {
		t.Error("IntToJSON Failed!")
	} else {
		t.Log("IntToJSON OK")
	}
	b = 1
	if b.ToJSON().ToString() != "1" {
		t.Errorf("IntToJSON Failed! %v", b.ToJSON().ToString())
	} else {
		t.Log("IntToJSON OK")
	}
}

func Test_IntToString(t *testing.T) {
	b := NewInt()
	if b.ToString() != "0" {
		t.Error("IntToJSON Failed!")
	} else {
		t.Log("IntToJSON OK")
	}
	b = -1
	if b.ToString() != "-1" {
		t.Errorf("IntToJSON Failed! %v", b.ToString())
	} else {
		t.Log("IntToJSON OK")
	}
}

func Test_IntToJSONString(t *testing.T) {
	b := NewInt()
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
