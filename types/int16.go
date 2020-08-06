package types

import (
	"strconv"
)

func NewInt16() Int16 {
	return Int16(0)
}

func (this Int16) ToNative() int16 {
	return int16(this)
}

func (this Int16) ToInt64() Int64 {
	return Int64(this)
}
func (this Int16) ToInt32() Int32 {
	return Int32(this)
}
func (this Int16) ToInt16() Int16 {
	return Int16(this)
}
func (this Int16) ToInt8() Int8 {
	return Int8(this)
}
func (this Int16) ToUInt64() UInt64 {
	return UInt64(this)
}
func (this Int16) ToUInt32() UInt32 {
	return UInt32(this)
}
func (this Int16) ToUInt16() UInt16 {
	return UInt16(this)
}
func (this Int16) ToUInt8() UInt8 {
	return UInt8(this)
}
func (this Int16) ToFloat64() Float64 {
	return Float64(this)
}
func (this Int16) ToFloat32() Float32 {
	return Float32(this)
}
func (this Int16) ToInt() Int {
	return Int(this)
}
func (this Int16) ToUInt() UInt {
	return UInt(this)
}
func (this Int16) ToFloat() Float {
	return Float(this)
}
func (this Int16) ToBool() Bool {
	var b bool = false
	if this != 0 {
		b = true
	}
	return Bool(b)
}
func (this Int16) ToJSON() JSON {
	return JSON(this.ToInt64())
}

func (this Int16) ToString() String {
	return String(strconv.FormatInt(this.ToInt64().ToNative(), 16))
}
