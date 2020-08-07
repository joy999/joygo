package types

import (
	"testing"
)

func Test_Push(t *testing.T) {
	arr := NewJSONArray()
	arr.Push(123)
	arr.Push("123")

}

func Test_ToJSONString(t *testing.T) {
	arr := NewJSONArray()
	arr.Push(123)
	arr.Push("123")

	subArr := NewJSONArray()
	subArr.Push("sub value")
	arr.Push(subArr)

	if js, err := arr.ToJSONString(); err != nil {
		t.Error(err)
	} else {
		t.Log(js)
	}
}

func Test_IndexOf(t *testing.T) {
	arr := NewJSONArray()
	arr.Push(123)
	arr.Push("123")

	n1 := arr.IndexOf(123)
	n2 := arr.IndexOf("123")
	n3 := arr.IndexOf("abc")

	if n1 != 0 || n2 != 1 || n3 != -1 {
		t.Error("Test_IndexOf FAILED!", n1, n2, n3)
	} else {
		t.Log("Test_IndexOf: ", n1, n2, n3)
	}
}
