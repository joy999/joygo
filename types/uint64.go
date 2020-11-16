package types

import (
	"encoding/json"
	"strconv"
)

func NewUInt64() UInt64 {
	return UInt64(0)
}

func (this UInt64) ToNative() uint64 {
	return uint64(this)
}
func (this UInt64) ToInt64() Int64 {
	return Int64(this)
}
func (this UInt64) ToInt32() Int32 {
	return Int32(this)
}
func (this UInt64) ToInt16() Int16 {
	return Int16(this)
}
func (this UInt64) ToInt8() Int8 {
	return Int8(this)
}
func (this UInt64) ToUInt64() UInt64 {
	return UInt64(this)
}
func (this UInt64) ToUInt32() UInt32 {
	return UInt32(this)
}
func (this UInt64) ToUInt16() UInt16 {
	return UInt16(this)
}
func (this UInt64) ToUInt8() UInt8 {
	return UInt8(this)
}
func (this UInt64) ToFloat64() Float64 {
	return Float64(this)
}
func (this UInt64) ToFloat32() Float32 {
	return Float32(this)
}
func (this UInt64) ToInt() Int {
	return Int(this)
}
func (this UInt64) ToUInt() UInt {
	return UInt(this)
}
func (this UInt64) ToFloat() Float {
	return Float(this)
}
func (this UInt64) ToBool() Bool {
	var b bool = false
	if this != 0 {
		b = true
	}
	return Bool(b)
}
func (this UInt64) ToJSON() *JSON {
	return &JSON{this.ToInt64()}
}
func (this UInt64) ToString() String {
	return String(strconv.FormatUint(this.ToUInt64().ToNative(), 10))
}

func (this UInt64) ToJSONString() (String, error) {
	if bs, err := json.Marshal(this); err == nil {
		return String(bs), nil
	} else {
		return "", err
	}
}
