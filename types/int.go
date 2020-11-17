package types

import (
	"encoding/json"
	"strconv"
)

func NewInt() Int {
	return Int(0)
}

func (this Int) ToNative() int {
	return int(this)
}

func (this Int) ToInt64() Int64 {
	return Int64(this)
}
func (this Int) ToInt32() Int32 {
	return Int32(this)
}
func (this Int) ToInt16() Int16 {
	return Int16(this)
}
func (this Int) ToInt8() Int8 {
	return Int8(this)
}
func (this Int) ToUInt64() UInt64 {
	if this < 0 {
		this = 0
	}
	return UInt64(this)
}
func (this Int) ToUInt32() UInt32 {
	if this < 0 {
		this = 0
	}
	return UInt32(this)
}
func (this Int) ToUInt16() UInt16 {
	if this < 0 {
		this = 0
	}
	return UInt16(this)
}
func (this Int) ToUInt8() UInt8 {
	if this < 0 {
		this = 0
	}
	return UInt8(this)
}
func (this Int) ToFloat64() Float64 {
	return Float64(this)
}
func (this Int) ToFloat32() Float32 {
	return Float32(this)
}
func (this Int) ToInt() Int {
	return Int(this)
}
func (this Int) ToUInt() UInt {
	if this < 0 {
		this = 0
	}
	return UInt(this)
}
func (this Int) ToFloat() Float {
	return Float(this)
}
func (this Int) ToBool() Bool {
	var b bool = false
	if this != 0 {
		b = true
	}
	return Bool(b)
}
func (this Int) ToJSON() *JSON {
	return &JSON{this.ToInt64()}
}

func (this Int) ToString() String {
	return String(strconv.Itoa(this.ToNative()))
}

func (this Int) ToJSONString() (String, error) {
	if bs, err := json.Marshal(this); err == nil {
		return String(bs), nil
	} else {
		return "", err
	}
}
