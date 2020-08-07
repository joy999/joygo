package types

import (
	"encoding/json"
	"strconv"
)

func NewInt8() Int8 {
	return Int8(0)
}

func (this Int8) ToNative() int8 {
	return int8(this)
}

func (this Int8) ToInt64() Int64 {
	return Int64(this)
}
func (this Int8) ToInt32() Int32 {
	return Int32(this)
}
func (this Int8) ToInt16() Int16 {
	return Int16(this)
}
func (this Int8) ToInt8() Int8 {
	return Int8(this)
}
func (this Int8) ToUInt64() UInt64 {
	return UInt64(this)
}
func (this Int8) ToUInt32() UInt32 {
	return UInt32(this)
}
func (this Int8) ToUInt16() UInt16 {
	return UInt16(this)
}
func (this Int8) ToUInt8() UInt8 {
	return UInt8(this)
}
func (this Int8) ToFloat64() Float64 {
	return Float64(this)
}
func (this Int8) ToFloat32() Float32 {
	return Float32(this)
}
func (this Int8) ToInt() Int {
	return Int(this)
}
func (this Int8) ToUInt() UInt {
	return UInt(this)
}
func (this Int8) ToFloat() Float {
	return Float(this)
}
func (this Int8) ToBool() Bool {
	var b bool = false
	if this != 0 {
		b = true
	}
	return Bool(b)
}
func (this Int8) ToJSON() *JSON {
	return &JSON{this.ToInt64()}
}

func (this Int8) ToString() String {
	return String(strconv.FormatInt(this.ToInt64().ToNative(), 8))
}

func (this Int8) ToJSONString() (String, error) {
	if bs, err := json.Marshal(this); err == nil {
		return String(bs), nil
	} else {
		return "", err
	}
}
