package types

import (
	"testing"
)

func Test_FloatToInt64(t *testing.T) {
	b := NewFloat()
	b++
	if b.ToInt64() != 1 {
		t.Error("FloatToInt64 Failed!")
	} else {
		t.Log("FloatToInt64 OK")
	}
	b = 1.2
	if b.ToInt64() != 1 {
		t.Error("FloatToInt64 Failed!")
	} else {
		t.Log("FloatToInt64 OK")
	}
}

func Test_FloatToInt32(t *testing.T) {
	b := NewFloat()
	b++
	if b.ToInt32() != 1 {
		t.Error("FloatToInt32 Failed!")
	} else {
		t.Log("FloatToInt32 OK")
	}
	b = 1.2
	if b.ToInt32() != 1 {
		t.Error("FloatToInt32 Failed!")
	} else {
		t.Log("FloatToInt32 OK")
	}
}

func Test_FloatToInt16(t *testing.T) {
	b := NewFloat()
	b++
	if b.ToInt16() != 1 {
		t.Error("FloatToInt16 Failed!")
	} else {
		t.Log("FloatToInt16 OK")
	}
	b = 1.2
	if b.ToInt16() != 1 {
		t.Error("FloatToInt16 Failed!")
	} else {
		t.Log("FloatToInt16 OK")
	}
}

func Test_FloatToInt8(t *testing.T) {
	b := NewFloat()
	b++
	if b.ToInt8() != 1 {
		t.Error("FloatToInt8 Failed!")
	} else {
		t.Log("FloatToInt8 OK")
	}
	b = 1.2
	if b.ToInt8() != 1 {
		t.Error("FloatToInt8 Failed!")
	} else {
		t.Log("FloatToInt8 OK")
	}
}

func Test_FloatToUInt64(t *testing.T) {
	b := NewFloat()
	b++
	if b.ToUInt64() != 1 {
		t.Error("FloatToUInt64 Failed!")
	} else {
		t.Log("FloatToUInt64 OK")
	}
	b = 1.2
	if b.ToUInt64() != 1 {
		t.Error("FloatToUInt64 Failed!")
	} else {
		t.Log("FloatToUInt64 OK")
	}
}

func Test_FloatToUInt32(t *testing.T) {
	b := NewFloat()
	b++
	if b.ToUInt32() != 1 {
		t.Error("FloatToUInt32 Failed!")
	} else {
		t.Log("FloatToUInt32 OK")
	}
	b = 1.2
	if b.ToUInt32() != 1 {
		t.Error("FloatToUInt32 Failed!")
	} else {
		t.Log("FloatToUInt32 OK")
	}
}

func Test_FloatToUInt16(t *testing.T) {
	b := NewFloat()
	b++
	if b.ToUInt16() != 1 {
		t.Error("FloatToUInt16 Failed!")
	} else {
		t.Log("FloatToUInt16 OK")
	}
	b = 1.2
	if b.ToUInt16() != 1 {
		t.Error("FloatToUInt16 Failed!")
	} else {
		t.Log("FloatToUInt16 OK")
	}
}

func Test_FloatToUInt8(t *testing.T) {
	b := NewFloat()
	b++
	if b.ToUInt8() != 1 {
		t.Error("FloatToUInt8 Failed!")
	} else {
		t.Log("FloatToUInt8 OK")
	}
	b = 1.2
	if b.ToUInt8() != 1 {
		t.Error("FloatToUInt8 Failed!")
	} else {
		t.Log("FloatToUInt8 OK")
	}
}

func Test_FloatToFloat64(t *testing.T) {
	b := NewFloat()
	b++
	if b.ToFloat64() != 1 {
		t.Error("FloatToFloat64 Failed!")
	} else {
		t.Log("FloatToFloat64 OK")
	}
	b = 1.2
	if b.ToFloat64() != 1.2 {
		t.Error("FloatToFloat64 Failed!")
	} else {
		t.Log("FloatToFloat64 OK")
	}
}

func Test_FloatToFloat32(t *testing.T) {
	b := NewFloat()
	b++
	if b.ToFloat32() != 1 {
		t.Error("FloatToFloat32 Failed!")
	} else {
		t.Log("FloatToFloat32 OK")
	}
	b = 1.2
	if b.ToFloat32() != 1.2 {
		t.Error("FloatToFloat32 Failed!")
	} else {
		t.Log("FloatToFloat32 OK")
	}
}

func Test_FloatToInt(t *testing.T) {
	b := NewFloat()
	b++
	if b.ToInt() != 1 {
		t.Error("FloatToInt Failed!")
	} else {
		t.Log("FloatToInt OK")
	}
	b = 1.2
	if b.ToInt() != 1 {
		t.Error("FloatToInt Failed!")
	} else {
		t.Log("FloatToInt OK")
	}
}

func Test_FloatToUInt(t *testing.T) {
	b := NewFloat()
	b++
	if b.ToUInt() != 1 {
		t.Error("FloatToUInt Failed!")
	} else {
		t.Log("FloatToUInt OK")
	}
	b = 1.2
	if b.ToUInt() != 1 {
		t.Error("FloatToUInt Failed!")
	} else {
		t.Log("FloatToUInt OK")
	}
}

func Test_FloatToFloat(t *testing.T) {
	b := NewFloat()
	b++
	if b.ToFloat() != 1 {
		t.Error("FloatToFloat Failed!")
	} else {
		t.Log("FloatToFloat OK")
	}
	b = 1.2
	if b.ToFloat() != 1.2 {
		t.Error("FloatToFloat Failed!")
	} else {
		t.Log("FloatToFloat OK")
	}
}

func Test_FloatToBool(t *testing.T) {
	b := NewFloat()
	if b.ToBool() != false {
		t.Error("FloatToBool Failed!")
	} else {
		t.Log("FloatToBool OK")
	}
	b = 1.2
	if b.ToBool() != true {
		t.Error("FloatToBool Failed!")
	} else {
		t.Log("FloatToBool OK")
	}
}

func Test_FloatToJSON(t *testing.T) {
	b := NewFloat()
	if b.ToJSON().ToString() != "0" {
		t.Error("FloatToJSON Failed!")
	} else {
		t.Log("FloatToJSON OK")
	}
	b = 1.2
	if b.ToJSON().ToString() != "1.2" {
		t.Errorf("FloatToJSON Failed! %v", b.ToJSON().ToString())
	} else {
		t.Log("FloatToJSON OK")
	}
}

func Test_FloatToString(t *testing.T) {
	b := NewFloat()
	if b.ToString() != "0" {
		t.Error("FloatToJSON Failed!")
	} else {
		t.Log("FloatToJSON OK")
	}
	b = 1.2
	if b.ToString() != "1.2" {
		t.Errorf("FloatToJSON Failed! %v", b.ToString())
	} else {
		t.Log("FloatToJSON OK")
	}
}

func Test_FloatToJSONString(t *testing.T) {
	b := NewFloat()
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
}
