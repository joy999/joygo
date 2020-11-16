package types

import "fmt"

func AnyToString(any Any) String {
	switch value := any.(type) {
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
	case UInt:
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
	case Bytes:
		return value.ToString()
	case *JSON:
		if v, err := value.ToJSONString(); err == nil {
			return v
		} else {
			return ""
		}
	case JSON:
		if v, err := value.ToJSONString(); err == nil {
			return v
		} else {
			return ""
		}
	case JSONArray:
		if v, err := value.ToJSONString(); err == nil {
			return v
		} else {
			return ""
		}
	case JSONMap:
		if v, err := value.ToJSONString(); err == nil {
			return v
		} else {
			return ""
		}
	case int:
		return Int(value).ToString()
	case int8:
		return Int8(value).ToString()
	case int16:
		return Int16(value).ToString()
	case int32:
		return Int32(value).ToString()
	case int64:
		return Int64(value).ToString()
	case uint:
		return UInt(value).ToString()
	case uint8:
		return UInt8(value).ToString()
	case uint16:
		return UInt16(value).ToString()
	case uint32:
		return UInt32(value).ToString()
	case uint64:
		return UInt64(value).ToString()
	case float32:
		return Float32(value).ToString()
	case float64:
		return Float64(value).ToString()
	case string:
		return String(value)
	case []byte:
		return String(value)
	case nil:
		return String("")
	default:
		return String(fmt.Sprintf("%T", value))
	}
}

func AnyToFloat64(any Any) Float64 {
	switch value := any.(type) {
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
	case UInt:
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
		return value.ToFloat64()
	case Float64:
		return Float64(value)
	case String:
		return value.ToFloat64()
	case Bytes:
		return value.ToString().ToFloat64()
	case int:
		return Float64(value)
	case int8:
		return Float64(value)
	case int16:
		return Float64(value)
	case int32:
		return Float64(value)
	case int64:
		return Float64(value)
	case uint:
		return Float64(value)
	case uint8:
		return Float64(value)
	case uint16:
		return Float64(value)
	case uint32:
		return Float64(value)
	case uint64:
		return Float64(value)

	case float32:
		return Float32(value).ToString().ToFloat64()
	case float64:
		return Float64(value)
	case string:
		return String(value).ToFloat64()
	default:
		return Float64(0)
	}
}

func AnyToUInt(any Any) UInt {
	return UInt(AnyToUInt64(any))
}

func AnyToUInt64(any Any) UInt64 {
	switch v := any.(type) {
	case Int:
		if v < 0 {
			v = 0
		}
		return UInt64(v)
	case Int8:
		if v < 0 {
			v = 0
		}
		return UInt64(v)
	case Int16:
		if v < 0 {
			v = 0
		}
		return UInt64(v)
	case Int32:
		if v < 0 {
			v = 0
		}
		return UInt64(v)
	case Int64:
		if v < 0 {
			v = 0
		}
		return UInt64(v)
	case UInt:
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
		if v < 0 {
			v = 0
		}
		return UInt64(v)
	case Float32:
		if v < 0 {
			v = 0
		}
		return UInt64(v)
	case Float64:
		if v < 0 {
			v = 0
		}
		return UInt64(v)
	case String:
		return v.ToUInt64()
	case int:
		if v < 0 {
			v = 0
		}
		return UInt64(v)
	case int8:
		if v < 0 {
			v = 0
		}
		return UInt64(v)
	case int16:
		if v < 0 {
			v = 0
		}
		return UInt64(v)
	case int32:
		if v < 0 {
			v = 0
		}
		return UInt64(v)
	case int64:
		if v < 0 {
			v = 0
		}
		return UInt64(v)
	case uint:
		return UInt64(v)
	case uint8:
		return UInt64(v)
	case uint16:
		return UInt64(v)
	case uint32:
		return UInt64(v)
	case uint64:
		return UInt64(v)

	case float32:
		if v < 0 {
			v = 0
		}
		return UInt64(v)
	case float64:
		if v < 0 {
			v = 0
		}
		return UInt64(v)
	case string:
		return String(v).ToUInt64()
	default:
		return UInt64(0)
	}
}

func AnyToInt(any Any) Int {
	return Int(AnyToInt64(any))
}

func AnyToInt64(any Any) Int64 {
	switch v := any.(type) {
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
	case UInt:
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
	case int:
		return Int64(v)
	case int8:
		return Int64(v)
	case int16:
		return Int64(v)
	case int32:
		return Int64(v)
	case int64:
		return Int64(v)
	case uint:
		return Int64(v)
	case uint8:
		return Int64(v)
	case uint16:
		return Int64(v)
	case uint32:
		return Int64(v)
	case uint64:
		return Int64(v)
	case float32:
		return Int64(v)
	case float64:
		return Int64(v)
	case string:
		return String(v).ToInt64()
	default:
		return Int64(0)
	}
}

func AnyToJSONArray(any Any) JSONArray {
	switch value := any.(type) {
	case *JSONArray:
		return JSONArray(*value)
	case *[]JSON:
		ret := make(JSONArray, 0)
		for _, v := range *value {
			ret.Push(v.Any)
		}
		return ret
	case *[]*JSON:
		ret := make(JSONArray, 0)
		for _, v := range *value {
			ret.Push(v.Any)
		}
		return ret
	case *[]interface{}:
		ret := make(JSONArray, 0)
		for _, v := range *value {
			ret = append(ret, Any(v))
		}
		return ret
	case JSONArray:
		return value
	case []JSON:
		ret := make(JSONArray, 0)
		for _, v := range value {
			ret.Push(v.Any)
		}
		return ret
	case []*JSON:
		ret := make(JSONArray, 0)
		for _, v := range value {
			ret.Push(v.Any)
		}
		return ret
	case []interface{}:
		ret := make(JSONArray, 0)
		for _, v := range value {
			ret.Push(Any(v))
		}
		return ret
	default:
		return JSONArray{}
	}
}

func AnyToJSONMap(any Any) JSONMap {
	return *AnyToJSONMapPtr(any)
}

func AnyToJSONMapPtr(any Any) *JSONMap {
	switch value := any.(type) {
	case *JSONMap:
		return value
	case *map[String]JSON:
		ret := make(JSONMap, 0)
		for k, v := range *value {
			ret[k] = v.Any
		}
		return &ret
	case *map[String]*JSON:
		ret := make(JSONMap, 0)
		for k, v := range *value {
			ret[k] = v.Any
		}
		return &ret
	case JSONMap:
		return &value
	case map[String]JSON:
		ret := make(JSONMap, 0)
		for k, v := range value {
			ret[k] = v.Any
		}
		return &ret
	case map[String]*JSON:
		ret := make(JSONMap, 0)
		for k, v := range value {
			ret[k] = v.Any
		}
		return &ret
	case *map[string]interface{}:
		ret := make(JSONMap, 0)
		for k, v := range *value {
			ret[String(k)] = Any(v)
		}
		return &ret
	case map[string]JSON:
		ret := make(JSONMap, 0)
		for k, v := range value {
			ret[String(k)] = v.Any
		}
		return &ret
	case map[string]*JSON:
		ret := make(JSONMap, 0)
		for k, v := range value {
			ret[String(k)] = v.Any
		}
		return &ret
	case *map[string]JSON:
		ret := make(JSONMap, 0)
		for k, v := range *value {
			ret[String(k)] = v.Any
		}
		return &ret
	case *map[string]*JSON:
		ret := make(JSONMap, 0)
		for k, v := range *value {
			ret[String(k)] = v.Any
		}
		return &ret
	case map[string]interface{}:
		ret := make(JSONMap, 0)
		for k, v := range value {
			ret[String(k)] = Any(v)
		}
		return &ret
	default:
		return &JSONMap{}
	}
}
