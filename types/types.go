package types

type Any interface{}

type Byte byte

type Bytes []Byte

type Int64 int64

type Int32 int32

type Int16 int16

type Int8 int8

type UInt64 uint64

type UInt32 uint32

type UInt16 uint16

type UInt8 uint8

type Float64 float64

type Float32 float32

type Int int

type UInt uint

type Float Float64

type Bool bool

type String string

type JSON struct {
	Any
}

type Array []Any

type StringArray []String

type JSONMap map[String]Any

type JSONArray Array

type JSONString struct {
	String
}

type Number interface {
	ToInt64() Int64
	ToInt32() Int32
	ToInt16() Int16
	ToInt8() Int8
	ToUInt64() UInt64
	ToUInt32() UInt32
	ToUInt16() UInt16
	ToUInt8() UInt8
	ToFloat64() Float64
	ToFloat32() Float32
	ToInt() Int
	ToUInt() UInt
	ToFloat() Float
	ToBool() Bool
	ToJSON() JSON
	ToString() String
}

type NumberArray []Number
