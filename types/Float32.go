package types

import (
	"encoding/json"
	"strconv"
)

func NewFloat32() Float32 {
	return Float32(0)
}

func (this Float32) ToNative() float32 {
	return float32(this)
}

func (this Float32) ToInt64() Int64 {
	return Int64(this)
}
func (this Float32) ToInt32() Int32 {
	return Int32(this)
}
func (this Float32) ToInt16() Int16 {
	return Int16(this)
}
func (this Float32) ToInt8() Int8 {
	return Int8(this)
}
func (this Float32) ToUInt64() UInt64 {
	return UInt64(this)
}
func (this Float32) ToUInt32() UInt32 {
	return UInt32(this)
}
func (this Float32) ToUInt16() UInt16 {
	return UInt16(this)
}
func (this Float32) ToUInt8() UInt8 {
	return UInt8(this)
}
func (this Float32) ToFloat64() Float64 {
	return this.ToString().ToFloat64()
}
func (this Float32) ToFloat32() Float32 {
	return Float32(this)
}
func (this Float32) ToInt() Int {
	return Int(this)
}
func (this Float32) ToUInt() UInt {
	return UInt(this)
}
func (this Float32) ToFloat() Float {
	return Float(this)
}
func (this Float32) ToBool() Bool {
	var b bool = false
	if this != 0 {
		b = true
	}
	return Bool(b)
}
func (this Float32) ToJSON() *JSON {
	return &JSON{this.ToInt64()}
}

func (this Float32) ToString() String {
	return String(strconv.FormatFloat(Float64(this).ToNative(), 'f', -1, 32))
}

func (this Float32) ToJSONString() (String, error) {
	if bs, err := json.Marshal(this); err == nil {
		return String(bs), nil
	} else {
		return "", err
	}
}
