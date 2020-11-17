package types

import (
	"encoding/json"
	"strconv"
)

func NewFloat() Float {
	return Float(0)
}

func (this Float) ToNative() float64 {
	return float64(this)
}

func (this Float) ToInt64() Int64 {
	return Int64(this)
}
func (this Float) ToInt32() Int32 {
	return Int32(this)
}
func (this Float) ToInt16() Int16 {
	return Int16(this)
}
func (this Float) ToInt8() Int8 {
	return Int8(this)
}
func (this Float) ToUInt64() UInt64 {
	if this < 0 {
		this = 0
	}
	return UInt64(this)
}
func (this Float) ToUInt32() UInt32 {
	if this < 0 {
		this = 0
	}
	return UInt32(this)
}
func (this Float) ToUInt16() UInt16 {
	if this < 0 {
		this = 0
	}
	return UInt16(this)
}
func (this Float) ToUInt8() UInt8 {
	if this < 0 {
		this = 0
	}
	return UInt8(this)
}
func (this Float) ToFloat64() Float64 {
	return Float64(this)
}
func (this Float) ToFloat32() Float32 {
	return Float32(this)
}
func (this Float) ToInt() Int {
	return Int(this)
}
func (this Float) ToUInt() UInt {
	if this < 0 {
		this = 0
	}
	return UInt(this)
}
func (this Float) ToFloat() Float {
	return Float(this)
}
func (this Float) ToBool() Bool {
	var b bool = false
	if this != 0 {
		b = true
	}
	return Bool(b)
}
func (this Float) ToJSON() *JSON {
	return &JSON{this}
}

func (this Float) ToString() String {
	return String(strconv.FormatFloat(this.ToNative(), 'f', -1, 64))
}

func (this Float) ToJSONString() (String, error) {
	if bs, err := json.Marshal(this); err == nil {
		return String(bs), nil
	} else {
		return "", err
	}
}
