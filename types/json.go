package types

import (
	"encoding/json"
)

func NewJSON(any Any) *JSON {
	return &JSON{any}
}

func ToJson(any Any) *JSON {
	return &JSON{any}
}

func (this JSON) ToInt() Int {
	return Int(this.ToInt64())
}

func (this JSON) ToInt8() Int8 {
	return Int8(this.ToInt64())
}

func (this JSON) ToInt16() Int16 {
	return Int16(this.ToInt64())
}

func (this JSON) ToInt32() Int32 {
	return Int32(this.ToInt64())
}

func (this JSON) ToInt64() Int64 {
	return AnyToInt64(this.Any)
}

func (this JSON) ToUInt() UInt {
	return UInt(this.ToUInt64())
}

func (this JSON) ToUInt8() UInt8 {
	return UInt8(this.ToUInt64())
}

func (this JSON) ToUInt16() UInt16 {
	return UInt16(this.ToUInt64())
}

func (this JSON) ToUInt32() UInt32 {
	return UInt32(this.ToUInt64())
}

func (this JSON) ToUInt64() UInt64 {
	return AnyToUInt64(this.Any)
}

func (this JSON) ToFloat() Float {
	return Float(this.ToFloat64())
}

func (this JSON) ToFloat32() Float32 {
	return Float32(this.ToFloat64())
}

func (this JSON) ToFloat64() Float64 {
	return AnyToFloat64(this.Any)
}

func (this JSON) ToString() String {
	return AnyToString(this.Any)
}

func (this JSON) ToJSONArray() JSONArray {
	return AnyToJSONArray(this.Any)
}

func (this JSON) ToJSONMap() JSONMap {
	return AnyToJSONMap(this.Any)
}

func (this JSON) ToJSONMapPtr() *JSONMap {
	return AnyToJSONMapPtr(this.Any)
}

func (this JSON) ToJSONString() (String, error) {
	if bs, err := json.Marshal(this.Any); err == nil {
		return String(bs), nil
	} else {
		return "", err
	}
}

func (this *JSON) Set(any Any) {
	this.Any = any
}

func (this *JSON) Assign(any Any) {
	this.Set(any)
}
