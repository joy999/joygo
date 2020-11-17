package types

import (
	"testing"
)

func Test_UIntToNative(t *testing.T) {
	b := NewUInt()
	b = 1
	if b.ToNative() != 1 {
		t.Error("UIntToNative Failed!")
	} else {
		t.Log("UIntToNative OK")
	}
}

func Test_UIntToInt64(t *testing.T) {
	b := NewUInt()
	b = 1
	if b.ToInt64() != 1 {
		t.Error("UIntToInt64 Failed!")
	} else {
		t.Log("UIntToInt64 OK")
	}
}

func Test_UIntToInt32(t *testing.T) {
	b := NewUInt()
	b = 1
	if b.ToInt32() != 1 {
		t.Error("UIntToInt32 Failed!")
	} else {
		t.Log("UIntToInt32 OK")
	}
}

func Test_UIntToInt16(t *testing.T) {
	b := NewUInt()
	b = 1
	if b.ToInt16() != 1 {
		t.Error("UIntToInt16 Failed!")
	} else {
		t.Log("UIntToInt16 OK")
	}
}

func Test_UIntToInt8(t *testing.T) {
	b := NewUInt()
	b = 1
	if b.ToInt8() != 1 {
		t.Error("UIntToInt8 Failed!")
	} else {
		t.Log("UIntToInt8 OK")
	}
}

func Test_UIntToUInt64(t *testing.T) {
	b := NewUInt()
	b = 1
	if b.ToUInt64() != 1 {
		t.Error("UIntToUInt64 Failed!")
	} else {
		t.Log("UIntToUInt64 OK")
	}
}

func Test_UIntToUInt32(t *testing.T) {
	b := NewUInt()
	b++
	if b.ToUInt32() != 1 {
		t.Error("UIntToUInt32 Failed!")
	} else {
		t.Log("UIntToUInt32 OK")
	}
}

func Test_UIntToUInt16(t *testing.T) {
	b := NewUInt()
	b++
	if b.ToUInt16() != 1 {
		t.Error("UIntToUInt16 Failed!")
	} else {
		t.Log("UIntToUInt16 OK")
	}
}

func Test_UIntToUInt8(t *testing.T) {
	b := NewUInt()
	b++
	if b.ToUInt8() != 1 {
		t.Error("UIntToUInt8 Failed!")
	} else {
		t.Log("UIntToUInt8 OK")
	}
}

func Test_UIntToFloat64(t *testing.T) {
	b := NewUInt()
	b++
	if b.ToFloat64() != 1 {
		t.Error("UIntToFloat64 Failed!")
	} else {
		t.Log("UIntToFloat64 OK")
	}
}

func Test_UIntToFloat32(t *testing.T) {
	b := NewUInt()
	b++
	if b.ToFloat32() != 1 {
		t.Error("UIntToFloat32 Failed!")
	} else {
		t.Log("UIntToFloat32 OK")
	}
}

func Test_UIntToInt(t *testing.T) {
	b := NewUInt()
	b++
	if b.ToInt() != 1 {
		t.Error("UIntToInt Failed!")
	} else {
		t.Log("UIntToInt OK")
	}
}

func Test_UIntToUInt(t *testing.T) {
	b := NewUInt()
	b++
	if b.ToUInt() != 1 {
		t.Error("UIntToUInt Failed!")
	} else {
		t.Log("UIntToUInt OK")
	}
}

func Test_UIntToFloat(t *testing.T) {
	b := NewUInt()
	b++
	if b.ToFloat() != 1 {
		t.Error("UIntToFloat Failed!")
	} else {
		t.Log("UIntToFloat OK")
	}
}

func Test_UIntToBool(t *testing.T) {
	b := NewUInt()
	if b.ToBool() != false {
		t.Error("UIntToBool Failed!")
	} else {
		t.Log("UIntToBool OK")
	}
	b = 333
	if b.ToBool() != true {
		t.Error("UIntToBool Failed!")
	} else {
		t.Log("UIntToBool OK")
	}
}

func Test_UIntToJSON(t *testing.T) {
	b := NewUInt()
	if b.ToJSON().ToString() != "0" {
		t.Error("UIntToJSON Failed!")
	} else {
		t.Log("UIntToJSON OK")
	}
}

func Test_UIntToString(t *testing.T) {
	b := NewUInt()
	if b.ToString() != "0" {
		t.Error("UIntToJSON Failed!")
	} else {
		t.Log("UIntToJSON OK")
	}
}

func Test_UIntToJSONString(t *testing.T) {
	b := NewUInt()
	if s, _ := b.ToJSONString(); s != "0" {
		t.Error("ToJSONString Failed!")
	} else {
		t.Log("ToJSONString OK")
	}
}
