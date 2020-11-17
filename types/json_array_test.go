package types

import (
	"testing"
)

func Test_JSONArrayPush(t *testing.T) {
	arr := NewJSONArray()
	arr.Push(123)
	arr.Push("123")

}

func Test_JSONArrayToJSONString(t *testing.T) {
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

func Test_JSONArrayIndexOf(t *testing.T) {
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

func Test_JSONArrayPop(t *testing.T) {
	arr := NewJSONArray()
	arr.Push(123)
	arr.Push("abc")

	if arr.Pop().ToString() != "abc" {
		t.Error("Test_JSONArrayPop Failed")
	}
	if arr.Pop().ToInt() != 123 {
		t.Error("Test_JSONArrayPop Failed")
	}
}

func Test_JSONArrayRemove(t *testing.T) {
	arr := NewJSONArray()
	arr.Push(123)
	arr.Push("abc")
	arr.Push(111)

	if s, _ := arr.ToJSONString(); s != `[123,"abc",111]` {
		t.Error("Test_JSONArrayRemove Failed")
	}

	arr.Remove(1)

	if s, _ := arr.ToJSONString(); s != `[123,111]` {
		t.Error("Test_JSONArrayRemove Failed")
	}

	arr.Push(222)
	arr.Push(333)

	if arr.GetLength() != 4 {
		t.Error("Test_JSONArrayRemove Failed")
	}

	arr.Remove(3)
	if s, _ := arr.ToJSONString(); s != `[123,111,222]` {
		t.Errorf("Test_JSONArrayRemove Failed %s %v", arr, s)
	}
	arr.Remove(0)
	if s, _ := arr.ToJSONString(); s != `[111,222]` {
		t.Errorf("Test_JSONArrayRemove Failed %s %v", arr, s)
	}

	arr.Push("abc")
	arr.Push(1)

	arr.RemoveItem(111)
	if s, _ := arr.ToJSONString(); s != `[222,"abc",1]` {
		t.Errorf("Test_JSONArrayRemoveItem Failed %s %v", arr, s)
	}

	if s, _ := arr.ToJSON().ToJSONString(); s != `[222,"abc",1]` {
		t.Errorf("Test_JSONArrayToJSON Failed %s %v", arr, s)
	}

	if arr.GetAsJSON(-1) != nil {
		t.Errorf("Test_JSONArrayGetAsJSON -1 Failed %s ", arr)
	}
}
