package types

import (
	"encoding/json"
	"strconv"
)

func NewUInt() UInt {
	return UInt(0)
}

func (this UInt) ToNative() uint {
	return uint(this)
}
func (this UInt) ToInt64() Int64 {
	return Int64(this)
}
func (this UInt) ToInt32() Int32 {
	return Int32(this)
}
func (this UInt) ToInt16() Int16 {
	return Int16(this)
}
func (this UInt) ToInt8() Int8 {
	return Int8(this)
}
func (this UInt) ToUInt64() UInt64 {
	return UInt64(this)
}
func (this UInt) ToUInt32() UInt32 {
	return UInt32(this)
}
func (this UInt) ToUInt16() UInt16 {
	return UInt16(this)
}
func (this UInt) ToUInt8() UInt8 {
	return UInt8(this)
}
func (this UInt) ToFloat64() Float64 {
	return Float64(this)
}
func (this UInt) ToFloat32() Float32 {
	return Float32(this)
}
func (this UInt) ToInt() Int {
	return Int(this)
}
func (this UInt) ToUInt() UInt {
	return UInt(this)
}
func (this UInt) ToFloat() Float {
	return Float(this)
}
func (this UInt) ToBool() Bool {
	var b bool = false
	if this != 0 {
		b = true
	}
	return Bool(b)
}
func (this UInt) ToJSON() *JSON {
	return &JSON{this.ToInt64()}
}

func (this UInt) ToString() String {
	return String(strconv.FormatUint(this.ToUInt64().ToNative(), 10))
}

func (this UInt) ToJSONString() (String, error) {
	if bs, err := json.Marshal(this); err == nil {
		return String(bs), nil
	} else {
		return "", err
	}
}
