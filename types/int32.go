package types

import (
	"encoding/json"
	"strconv"
)

func NewInt32() Int32 {
	return Int32(0)
}

func (this Int32) ToNative() int32 {
	return int32(this)
}
func (this Int32) ToInt64() Int64 {
	return Int64(this)
}
func (this Int32) ToInt32() Int32 {
	return Int32(this)
}
func (this Int32) ToInt16() Int16 {
	return Int16(this)
}
func (this Int32) ToInt8() Int8 {
	return Int8(this)
}
func (this Int32) ToUInt64() UInt64 {
	if this < 0 {
		this = 0
	}
	return UInt64(this)
}
func (this Int32) ToUInt32() UInt32 {
	if this < 0 {
		this = 0
	}
	return UInt32(this)
}
func (this Int32) ToUInt16() UInt16 {
	if this < 0 {
		this = 0
	}
	return UInt16(this)
}
func (this Int32) ToUInt8() UInt8 {
	if this < 0 {
		this = 0
	}
	return UInt8(this)
}
func (this Int32) ToFloat64() Float64 {
	return Float64(this)
}
func (this Int32) ToFloat32() Float32 {
	return Float32(this)
}
func (this Int32) ToInt() Int {
	return Int(this)
}
func (this Int32) ToUInt() UInt {
	if this < 0 {
		this = 0
	}
	return UInt(this)
}
func (this Int32) ToFloat() Float {
	return Float(this)
}
func (this Int32) ToBool() Bool {
	var b bool = false
	if this != 0 {
		b = true
	}
	return Bool(b)
}
func (this Int32) ToJSON() *JSON {
	return &JSON{this.ToInt64()}
}

func (this Int32) ToString() String {
	return String(strconv.FormatInt(this.ToInt64().ToNative(), 10))
}

func (this Int32) ToJSONString() (String, error) {
	if bs, err := json.Marshal(this); err == nil {
		return String(bs), nil
	} else {
		return "", err
	}
}
