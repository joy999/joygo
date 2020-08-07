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
	switch v := this.Any.(type) {
	case Int:
		return Int64(v)
	case Int8:
		return Int64(v)
	case Int16:
		return Int64(v)
	case Int32:
		return Int64(v)
	case Int64:
		return Int64(v)
	case UInt8:
		return Int64(v)
	case UInt16:
		return Int64(v)
	case UInt32:
		return Int64(v)
	case UInt64:
		return Int64(v)
	case Float:
		return Int64(v)
	case Float32:
		return Int64(v)
	case Float64:
		return Int64(v)
	case String:
		return v.ToInt64()
	default:
		return Int64(0)
	}
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
	switch v := this.Any.(type) {
	case Int:
		return UInt64(v)
	case Int8:
		return UInt64(v)
	case Int16:
		return UInt64(v)
	case Int32:
		return UInt64(v)
	case Int64:
		return UInt64(v)
	case UInt8:
		return UInt64(v)
	case UInt16:
		return UInt64(v)
	case UInt32:
		return UInt64(v)
	case UInt64:
		return UInt64(v)
	case Float:
		return UInt64(v)
	case Float32:
		return UInt64(v)
	case Float64:
		return UInt64(v)
	case String:
		return v.ToUInt64()
	default:
		return UInt64(0)
	}
}

func (this JSON) ToFloat() Float {
	return Float(this.ToFloat64())
}

func (this JSON) ToFloat32() Float32 {
	return Float32(this.ToFloat64())
}

func (this JSON) ToFloat64() Float64 {
	switch value := this.Any.(type) {
	case Int:
		return Float64(value)
	case Int8:
		return Float64(value)
	case Int16:
		return Float64(value)
	case Int32:
		return Float64(value)
	case Int64:
		return Float64(value)
	case UInt8:
		return Float64(value)
	case UInt16:
		return Float64(value)
	case UInt32:
		return Float64(value)
	case UInt64:
		return Float64(value)
	case Float:
		return Float64(value)
	case Float32:
		return Float64(value)
	case Float64:
		return Float64(value)

	case String:
		return value.ToFloat64()
	default:
		return Float64(0)
	}
}

func (this JSON) ToString() String {
	switch value := this.Any.(type) {
	case Int:
		return value.ToString()
	case Int8:
		return value.ToString()
	case Int16:
		return value.ToString()
	case Int32:
		return value.ToString()
	case Int64:
		return value.ToString()
	case UInt8:
		return value.ToString()
	case UInt16:
		return value.ToString()
	case UInt32:
		return value.ToString()
	case UInt64:
		return value.ToString()
	case Float:
		return value.ToString()
	case Float32:
		return value.ToString()
	case Float64:
		return value.ToString()
	case String:
		return value
	default:
		return String("")
	}
}

func (this JSON) ToJSONArray() JSONArray {
	switch v := this.Any.(type) {
	case *JSONArray:
		return JSONArray(*v)
	case *[]JSON:
		ret := make([]Any, 0)
		for _, x := range *v {
			ret = append(ret, x.Any)
		}
		return JSONArray(ret)
	case *[]*JSON:
		ret := make([]Any, 0)
		for _, x := range *v {
			ret = append(ret, x.Any)
		}
		return JSONArray(ret)
	case *[]interface{}:
		ret := make([]Any, 0)
		for _, x := range *v {
			ret = append(ret, x)
		}
		return JSONArray(ret)
	case JSONArray:
		return v
	case []JSON:
		ret := make([]Any, 0)
		for _, x := range v {
			ret = append(ret, x.Any)
		}
		return JSONArray(ret)
	case []*JSON:
		ret := make([]Any, 0)
		for _, x := range v {
			ret = append(ret, x.Any)
		}
		return JSONArray(ret)
	case []interface{}:
		ret := make([]Any, 0)
		for _, x := range v {
			ret = append(ret, x)
		}
		return JSONArray(ret)
	default:
		return JSONArray{}
	}
}

func (this JSON) ToJSONMap() JSONMap {
	switch value := this.Any.(type) {
	case *JSONMap:
		return JSONMap(*value)
	case *map[String]JSON:
		ret := make(map[String]Any, 0)
		for k, v := range *value {
			ret[String(k)] = Any(v.Any)
		}
		return ret
	case *map[String]*JSON:
		ret := make(map[String]Any, 0)
		for k, v := range *value {
			ret[String(k)] = Any(v.Any)
		}
		return ret
	case *map[string]interface{}:
		ret := make(map[String]Any, 0)
		for k, v := range *value {
			ret[String(k)] = v
		}
		return ret
	case JSONMap:
		return JSONMap(value)
	case map[String]JSON:
		ret := make(map[String]Any, 0)
		for k, v := range value {
			ret[String(k)] = v.Any
		}
		return ret
	case map[String]*JSON:
		ret := make(map[String]Any, 0)
		for k, v := range value {
			ret[String(k)] = v.Any
		}
		return ret
	case map[string]interface{}:
		ret := make(map[String]Any, 0)
		for k, v := range value {
			ret[String(k)] = v
		}
		return ret
	default:
		return JSONMap{}
	}
}

func (this JSON) ToJSONString() (String, error) {
	if bs, err := json.Marshal(this.Any); err == nil {
		return String(bs), nil
	} else {
		return "", err
	}
}
