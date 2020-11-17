package types

import (
	"testing"
)

func Test_BoolToNative(t *testing.T) {
	b := NewBool()
	b = true
	if b.ToNative() != true {
		t.Error("Bool to Native true Failed!")
	} else {
		t.Log("Bool to native true OK")
	}
	b = false
	if b.ToNative() != false {
		t.Error("Bool to Native false Failed!")
	} else {
		t.Log("Bool to native false OK")
	}
}

func Test_BoolToInt64(t *testing.T) {
	b := NewBool()
	b = true
	if b.ToInt64() != 1 {
		t.Error("Bool to Int64 Failed!")
	} else {
		t.Log("Bool to Int64 OK")
	}
	b = false
	if b.ToInt64() != 0 {
		t.Error("Bool to Int64 false Failed!")
	} else {
		t.Log("Bool to Int64 false OK")
	}
}

func Test_BoolToUInt(t *testing.T) {
	b := NewBool()
	b = true
	if b.ToUInt() != 1 {
		t.Error("Bool ToUInt true Failed!")
	} else {
		t.Log("Bool ToUInt true OK")
	}
	b = false
	if b.ToUInt() != 0 {
		t.Error("Bool ToUInt false Failed!")
	} else {
		t.Log("Bool ToUInt false OK")
	}
}

func Test_BoolToInt32(t *testing.T) {
	b := NewBool()
	b = true
	if b.ToInt32() != 1 {
		t.Error("Bool ToInt32 Failed!")
	} else {
		t.Log("Bool ToInt32 OK")
	}
	b = false
	if b.ToInt32() != 0 {
		t.Error("Bool ToInt32 false Failed!")
	} else {
		t.Log("Bool ToInt32 false OK")
	}
}

func Test_BoolToInt16(t *testing.T) {
	b := NewBool()
	b = true
	if b.ToInt16() != 1 {
		t.Error("Bool ToInt16 Failed!")
	} else {
		t.Log("Bool ToInt16 OK")
	}
	b = false
	if b.ToInt16() != 0 {
		t.Error("Bool ToInt16 false Failed!")
	} else {
		t.Log("Bool ToInt16 false OK")
	}
}

func Test_BoolToInt8(t *testing.T) {
	b := NewBool()
	b = true
	if b.ToInt8() != 1 {
		t.Error("Bool ToInt8 Failed!")
	} else {
		t.Log("Bool ToInt8 OK")
	}
	b = false
	if b.ToInt8() != 0 {
		t.Error("Bool ToInt8 false Failed!")
	} else {
		t.Log("Bool ToInt8 false OK")
	}
}

func Test_BoolToUInt64(t *testing.T) {
	b := NewBool()
	b = true
	if b.ToUInt64() != 1 {
		t.Error("Bool ToUInt64 Failed!")
	} else {
		t.Log("Bool ToUInt64 OK")
	}
	b = false
	if b.ToUInt64() != 0 {
		t.Error("Bool ToUInt64 false Failed!")
	} else {
		t.Log("Bool ToUInt64 false OK")
	}
}

func Test_BoolToUInt32(t *testing.T) {
	b := NewBool()
	b = true
	if b.ToUInt32() != 1 {
		t.Error("Bool ToUInt32 Failed!")
	} else {
		t.Log("Bool ToUInt32 OK")
	}
	b = false
	if b.ToUInt32() != 0 {
		t.Error("Bool ToUInt32 false Failed!")
	} else {
		t.Log("Bool ToUInt32 false OK")
	}
}

func Test_BoolToUInt16(t *testing.T) {
	b := NewBool()
	b = true
	if b.ToUInt16() != 1 {
		t.Error("Bool ToUInt16 Failed!")
	} else {
		t.Log("Bool ToUInt16 OK")
	}
	b = false
	if b.ToUInt16() != 0 {
		t.Error("Bool ToUInt16 false Failed!")
	} else {
		t.Log("Bool ToUInt16 false OK")
	}
}

func Test_BoolToUInt8(t *testing.T) {
	b := NewBool()
	b = true
	if b.ToUInt8() != 1 {
		t.Error("Bool ToUInt8 Failed!")
	} else {
		t.Log("Bool ToUInt8 OK")
	}
	b = false
	if b.ToUInt8() != 0 {
		t.Error("Bool ToUInt8 false Failed!")
	} else {
		t.Log("Bool ToUInt8 false OK")
	}
}

func Test_BoolToFloat64(t *testing.T) {
	b := NewBool()
	b = true
	if b.ToFloat64() != 1 {
		t.Error("Bool ToFloat64 Failed!")
	} else {
		t.Log("Bool ToFloat64 OK")
	}
	b = false
	if b.ToFloat64() != 0 {
		t.Error("Bool ToFloat64 false Failed!")
	} else {
		t.Log("Bool ToFloat64 false OK")
	}
}

func Test_BoolToFloat32(t *testing.T) {
	b := NewBool()
	b = true
	if b.ToFloat32() != 1 {
		t.Error("Bool ToFloat32 Failed!")
	} else {
		t.Log("Bool ToFloat32 OK")
	}
	b = false
	if b.ToFloat32() != 0 {
		t.Error("Bool ToFloat32 false Failed!")
	} else {
		t.Log("Bool ToFloat32 false OK")
	}
}

func Test_BoolToInt(t *testing.T) {
	b := NewBool()
	b = true
	if b.ToInt() != 1 {
		t.Error("Bool ToInt Failed!")
	} else {
		t.Log("Bool ToInt OK")
	}
	b = false
	if b.ToInt() != 0 {
		t.Error("Bool ToInt false Failed!")
	} else {
		t.Log("Bool ToInt false OK")
	}
}

func Test_BoolToFloat(t *testing.T) {
	b := NewBool()
	b = true
	if b.ToFloat() != 1 {
		t.Error("Bool ToFloat Failed!")
	} else {
		t.Log("Bool ToFloat OK")
	}
	b = false
	if b.ToFloat() != 0 {
		t.Error("Bool ToFloat false Failed!")
	} else {
		t.Log("Bool ToFloat false OK")
	}
}

func Test_BoolToBool(t *testing.T) {
	b := NewBool()
	b = true
	if b.ToBool() != true {
		t.Error("Bool ToBool Failed!")
	} else {
		t.Log("Bool ToBool OK")
	}
	b = false
	if b.ToBool() != false {
		t.Error("Bool ToBool false Failed!")
	} else {
		t.Log("Bool ToBool false OK")
	}
}

func Test_BoolToJSON(t *testing.T) {
	b := NewBool()
	b = true
	if b.ToJSON().ToString() != "true" {
		t.Errorf("Bool ToJSON true Failed! %T : %v", b.ToJSON().ToString(), b.ToJSON().ToString())
	} else {
		t.Log("Bool ToJSON true OK")
	}
	b = false
	if b.ToJSON().ToString() != `false` {
		t.Error("Bool ToJSON false Failed!")
	} else {
		t.Log("Bool ToJSON false OK")
	}
}

func Test_BoolToString(t *testing.T) {
	b := NewBool()
	b = true
	if b.ToString() != "true" {
		t.Errorf("Bool ToString true Failed! %T : %v", b.ToString(), b.ToString())
	} else {
		t.Log("Bool ToString true OK")
	}
	b = false
	if b.ToString() != `false` {
		t.Error("Bool ToString false Failed!")
	} else {
		t.Log("Bool ToJToStringSON false OK")
	}
}

func Test_BoolToJSONString(t *testing.T) {
	b := NewBool()
	b = true
	if s, _ := b.ToJSONString(); s != "true" {
		t.Errorf("Bool ToJSONString true Failed! %T : %v", s, s)
	} else {
		t.Log("Bool ToJSONString true OK")
	}
	b = false
	if s, _ := b.ToJSONString(); s != `false` {
		t.Error("Bool ToJSONString false Failed!")
	} else {
		t.Log("Bool ToJSONString false OK")
	}
	var bb Bool
	if s, _ := bb.ToJSONString(); s != `false` {
		t.Error("Bool ToJSONString false Failed!")
	} else {
		t.Log("Bool ToJSONString false OK")
	}
}
