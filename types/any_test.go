package types

import (
	"fmt"
	"testing"
)

func Test_AnyToString(t *testing.T) {
	jArr := NewJSONArray()
	jArr.Push(1, 2)
	jArr.Push("a", "b")
	jMap := NewJSONMap()
	jMap.Set("a", 1)
	jMap.Set("b", "c")
	arr := [...]Any{
		Int(123), Int(-123), Int8(111), Int8(-111), Int16(11111), Int16(-12222), Int32(1100000000), Int32(-2122222222), Int64(1000000000000), Int64(-99999999999999),
		int(123), int(-123), int8(111), int8(-111), int16(11111), int16(-12222), int32(1100000000), int32(-2122222222), int64(1000000000000), int64(-99999999999999),
		UInt(223), UInt8(211), UInt16(61111), UInt32(4100000000), UInt64(19999000000000000),
		uint(223), uint8(211), uint16(61111), uint32(4100000000), uint64(19999000000000000),
		Float(1.22), Float32(2.333), Float64(12000.33333), String("hello,world!"),
		Float(-1.22), Float32(-2.333), Float64(-12000.33333),
		float32(2.333), float64(12000.33333), string("hello,world!"),
		Bytes("haha"), []byte("hello"), NewJSON("hehe"), *NewJSON("heheP"),
		jArr, jMap, func() {}, func(int) {}, nil,
		&JSON{nil}, JSON{nil}, JSONArray{}, JSONMap{},
	}
	mArr := [...]String{
		"123", "-123", "111", "-111", "11111", "-12222", "1100000000", "-2122222222", "1000000000000", "-99999999999999",
		"123", "-123", "111", "-111", "11111", "-12222", "1100000000", "-2122222222", "1000000000000", "-99999999999999",
		"223", "211", "61111", "4100000000", "19999000000000000",
		"223", "211", "61111", "4100000000", "19999000000000000",
		"1.22", "2.333", "12000.33333", "hello,world!",
		"-1.22", "-2.333", "-12000.33333",
		"2.333", "12000.33333", "hello,world!",
		"haha", "hello", `"hehe"`, `"heheP"`,
		`[1,2,"a","b"]`, `{"a":1,"b":"c"}`, "func()", "func(int)", "",
		"null", "null", "[]", "{}",
	}
	var s String

	for i := 0; i < len(arr); i++ {
		s = AnyToString(arr[i])
		if s != mArr[i] {
			ss := fmt.Sprintf("%T To String Failed! Any: %s Result: %s Match: %s", arr[i], arr[i], s, mArr[i])
			t.Error(ss)
		} else {
			ss := fmt.Sprintf("%T To String OK", arr[i])
			t.Log(ss)
		}
	}

}

func Test_AnyToFloat64(t *testing.T) {
	jArr := NewJSONArray()
	jArr.Push(1, 2)
	jArr.Push("a", "b")
	jMap := NewJSONMap()
	jMap.Set("a", 1)
	jMap.Set("b", "c")
	arr := [...]Any{
		Int(123), Int(-123), Int8(111), Int8(-111), Int16(11111), Int16(-12222), Int32(1100000000), Int32(-2122222222), Int64(1000000000000), Int64(-99999999999999),
		int(123), int(-123), int8(111), int8(-111), int16(11111), int16(-12222), int32(1100000000), int32(-2122222222), int64(1000000000000), int64(-99999999999999),
		UInt(223), UInt8(211), UInt16(61111), UInt32(4100000000), UInt64(19999000000000000),
		uint(223), uint8(211), uint16(61111), uint32(4100000000), uint64(19999000000000000),
		Float(1.22), Float32(2.33301), Float64(12000.33333), String("hello,world!"),
		Float(-1.22), Float32(-2.33301), Float64(-12000.33333),
		float32(2.33305), float64(12000.33333), string("hello,world!"),
		Bytes("haha"), []byte("hello"), NewJSON("hehe"), *NewJSON("heheP"),
		jArr, jMap, func() {}, func(int) {}, nil,
		Bool(true), Bool(false), true, false,
	}
	mArr := [...]Float64{
		123, -123, 111, -111, 11111, -12222, 1100000000, -2122222222, 1000000000000, -99999999999999,
		123, -123, 111, -111, 11111, -12222, 1100000000, -2122222222, 1000000000000, -99999999999999,
		223, 211, 61111, 4100000000, 19999000000000000,
		223, 211, 61111, 4100000000, 19999000000000000,
		1.22, 2.33301, 12000.33333, 0,
		-1.22, -2.33301, -12000.33333,
		2.33305, 12000.33333, 0,
		0, 0, 0, 0,
		0, 0, 0, 0, 0,
		1, 0, 1, 0,
	}
	var s Float64

	for i := 0; i < len(arr); i++ {
		s = AnyToFloat64(arr[i])
		if s != mArr[i] {
			ss := fmt.Sprintf("%T To Float64 Failed! Any: %v Result: %v Match: %v", arr[i], arr[i], s, mArr[i])
			t.Error(ss)
		} else {
			ss := fmt.Sprintf("%T To Float64 OK", arr[i])
			t.Log(ss)
		}
	}

}

func Test_AnyToUInt(t *testing.T) {
	jArr := NewJSONArray()
	jArr.Push(1, 2)
	jArr.Push("a", "b")
	jMap := NewJSONMap()
	jMap.Set("a", 1)
	jMap.Set("b", "c")
	arr := [...]Any{
		Int(123), Int(-123), Int8(111), Int8(-111), Int16(11111), Int16(-12222), Int32(1100000000), Int32(-2122222222), Int64(1000000000000), Int64(-99999999999999),
		int(123), int(-123), int8(111), int8(-111), int16(11111), int16(-12222), int32(1100000000), int32(-2122222222), int64(1000000000000), int64(-99999999999999),
		UInt(223), UInt8(211), UInt16(61111), UInt32(4100000000), UInt64(19999000000000000),
		uint(223), uint8(211), uint16(61111), uint32(4100000000), uint64(19999000000000000),
		Float(1.22), Float32(2.33301), Float64(12000.33333), String("hello,world!"),
		Float(-1.22), Float32(-2.33301), Float64(-12000.33333),
		float32(2.33305), float64(12000.33333), string("hello,world!"),
		float32(-2.33305), float64(-12000.33333),
		Bytes("haha"), []byte("hello"), NewJSON("hehe"), *NewJSON("heheP"),
		jArr, jMap, func() {}, func(int) {}, nil,
		Bool(true), Bool(false), true, false,
	}
	mArr := [...]UInt{
		123, 0, 111, 0, 11111, 0, 1100000000, 0, 1000000000000, 0,
		123, 0, 111, 0, 11111, 0, 1100000000, 0, 1000000000000, 0,
		223, 211, 61111, 4100000000, 19999000000000000,
		223, 211, 61111, 4100000000, 19999000000000000,
		1, 2, 12000, 0,
		0, 0, 0,
		2, 12000, 0,
		0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0, 0,
		1, 0, 1, 0,
	}
	var s UInt

	for i := 0; i < len(arr); i++ {
		s = AnyToUInt(arr[i])
		if s != mArr[i] {
			ss := fmt.Sprintf("%T To UInt Failed! Any: %v Result: %v Match: %v", arr[i], arr[i], s, mArr[i])
			t.Error(ss)
		} else {
			ss := fmt.Sprintf("%T To UInt OK", arr[i])
			t.Log(ss)
		}
	}

}

func Test_AnyToInt(t *testing.T) {
	jArr := NewJSONArray()
	jArr.Push(1, 2)
	jArr.Push("a", "b")
	jMap := NewJSONMap()
	jMap.Set("a", 1)
	jMap.Set("b", "c")
	arr := [...]Any{
		Int(123), Int(-123), Int8(111), Int8(-111), Int16(11111), Int16(-12222), Int32(1100000000), Int32(-2122222222), Int64(1000000000000), Int64(-99999999999999),
		int(123), int(-123), int8(111), int8(-111), int16(11111), int16(-12222), int32(1100000000), int32(-2122222222), int64(1000000000000), int64(-99999999999999),
		UInt(223), UInt8(211), UInt16(61111), UInt32(4100000000), UInt64(19999000000000000),
		uint(223), uint8(211), uint16(61111), uint32(4100000000), uint64(19999000000000000),
		Float(1.22), Float32(2.33301), Float64(12000.33333), String("hello,world!"),
		Float(-1.22), Float32(-2.33301), Float64(-12000.33333),
		float32(2.33305), float64(12000.33333), string("hello,world!"),
		float32(-2.33305), float64(-12000.33333),
		Bytes("haha"), []byte("hello"), NewJSON("hehe"), *NewJSON("heheP"),
		jArr, jMap, func() {}, func(int) {}, nil,
		Bool(true), Bool(false), true, false,
	}
	mArr := [...]Int{
		123, -123, 111, -111, 11111, -12222, 1100000000, -2122222222, 1000000000000, -99999999999999,
		123, -123, 111, -111, 11111, -12222, 1100000000, -2122222222, 1000000000000, -99999999999999,
		223, 211, 61111, 4100000000, 19999000000000000,
		223, 211, 61111, 4100000000, 19999000000000000,
		1, 2, 12000, 0,
		-1, -2, -12000,
		2, 12000, 0,
		-2, -12000,
		0, 0, 0, 0,
		0, 0, 0, 0, 0,
		1, 0, 1, 0,
	}
	var s Int

	for i := 0; i < len(arr); i++ {
		s = AnyToInt(arr[i])
		if s != mArr[i] {
			ss := fmt.Sprintf("%T To UInt Failed! Any: %v Result: %v Match: %v", arr[i], arr[i], s, mArr[i])
			t.Error(ss)
		} else {
			ss := fmt.Sprintf("%T To UInt OK", arr[i])
			t.Log(ss)
		}
	}

}

func Test_AnyToJSONArray(t *testing.T) {
	jArr := NewJSONArray()
	jArr.Push(1, 2)
	jArr.Push("a", "b")
	jMap := NewJSONMap()
	jMap.Set("a", 1)

	aj := make([]JSON, 0)
	aj = append(aj, JSON{1})
	aj = append(aj, JSON{2})
	aj = append(aj, JSON{jArr})

	ajp := make([]*JSON, 0)
	ajp = append(ajp, &JSON{1})
	ajp = append(ajp, &JSON{2})
	ajp = append(ajp, &JSON{jArr})

	ai := make([]interface{}, 0)
	ai = append(ai, 1)
	ai = append(ai, 2)
	ai = append(ai, jArr)

	arr := [...]Any{
		&jArr, jArr,
		&jMap, jMap,
		aj, &aj,
		ajp, &ajp,
		ai, &ai,
		1, func() {},
	}
	mArr := [...]String{
		`[1,2,"a","b"]`, `[1,2,"a","b"]`,
		`["a",1]`, `["a",1]`,
		`[1,2,[1,2,"a","b"]]`, `[1,2,[1,2,"a","b"]]`,
		`[1,2,[1,2,"a","b"]]`, `[1,2,[1,2,"a","b"]]`,
		`[1,2,[1,2,"a","b"]]`, `[1,2,[1,2,"a","b"]]`,
		`[]`, `[]`,
	}
	var res JSONArray
	var s String

	for i := 0; i < len(arr); i++ {
		res = AnyToJSONArray(arr[i])
		s, _ = res.ToJSONString()

		if s != mArr[i] {
			ss := fmt.Sprintf("%T To JSONArray Failed! Any: %v Result: %v Match: %v", arr[i], arr[i], s, mArr[i])
			t.Error(ss)
		} else {
			ss := fmt.Sprintf("%T To JSONArray OK", arr[i])
			t.Log(ss)
		}
	}

}

func Test_AnyToJSONMap(t *testing.T) {
	jArr := NewJSONArray()
	jArr.Push(1, 2)
	jArr.Push("a", "b")
	jMap := NewJSONMap()
	jMap.Set("a", 1)

	aj := make([]JSON, 0)
	aj = append(aj, JSON{1})
	aj = append(aj, JSON{2})
	aj = append(aj, JSON{jArr})

	ajp := make([]*JSON, 0)
	ajp = append(ajp, &JSON{1})
	ajp = append(ajp, &JSON{2})
	ajp = append(ajp, &JSON{jArr})

	ai := make([]interface{}, 0)
	ai = append(ai, 1)
	ai = append(ai, 2)
	ai = append(ai, jArr)

	mSj := make(map[String]JSON, 0)
	mSj[String("a")] = JSON{1}

	mSjP := make(map[String]*JSON, 0)
	mSjP[String("a")] = &JSON{"b"}

	msj := make(map[string]JSON, 0)
	msj["a"] = JSON{1}

	msjP := make(map[string]*JSON, 0)
	msjP["a"] = &JSON{"b"}

	msi := make(map[string]interface{}, 0)
	msi["a"] = 1

	arr := [...]Any{
		&jArr, jArr,
		&jMap, jMap,
		aj, &aj,
		ajp, &ajp,
		ai, &ai,
		1, func() {},

		mSj, &mSj,
		mSjP, &mSjP,
		msj, &msj,
		msjP, &msjP,
		msi, &msi,
	}
	mArr := [...]String{
		`{"0":1,"1":2,"2":"a","3":"b"}`, `{"0":1,"1":2,"2":"a","3":"b"}`,
		`{"a":1}`, `{"a":1}`,
		`{"0":1,"1":2,"2":[1,2,"a","b"]}`, `{"0":1,"1":2,"2":[1,2,"a","b"]}`,
		`{"0":1,"1":2,"2":[1,2,"a","b"]}`, `{"0":1,"1":2,"2":[1,2,"a","b"]}`,
		`{"0":1,"1":2,"2":[1,2,"a","b"]}`, `{"0":1,"1":2,"2":[1,2,"a","b"]}`,
		`{}`, `{}`,

		`{"a":1}`, `{"a":1}`,
		`{"a":"b"}`, `{"a":"b"}`,
		`{"a":1}`, `{"a":1}`,
		`{"a":"b"}`, `{"a":"b"}`,
		`{"a":1}`, `{"a":1}`,
	}
	var res JSONMap
	var s String

	for i := 0; i < len(arr); i++ {
		res = AnyToJSONMap(arr[i])
		s, _ = res.ToJSONString()

		if s != mArr[i] {
			ss := fmt.Sprintf("%T To JSONMap Failed! Any: %v Result: %v Match: %v", arr[i], arr[i], s, mArr[i])
			t.Error(ss)
		} else {
			ss := fmt.Sprintf("%T To JSONMap OK", arr[i])
			t.Log(ss)
		}
	}

}
