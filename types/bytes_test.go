package types

import (
	"testing"
)

func Test_BytesToNative(t *testing.T) {
	b := NewBytes()
	b.AssignString("hello,world!")

	bs := []byte("hello,world!")

	if string(b.ToNative()) != string(bs) {
		t.Error("Bytes ToNative Failed!")
	} else {
		t.Log("Bytes ToNative OK")
	}
}

func Test_BytesToString(t *testing.T) {
	b := NewBytes()
	b.AssignString("hello,world!")

	if b.ToString().ToNative() != "hello,world!" {
		t.Error("Bytes ToString Failed!")
	} else {
		t.Log("Bytes ToString OK")
	}
}

func Test_BytesAppendByte(t *testing.T) {
	b := NewBytes()
	b.AppendByte('a')
	b.AppendByte('b')
	b.AppendByte('c')

	if b.ToString().ToNative() != "abc" {
		t.Error("Bytes AppendByte Failed!")
	} else {
		t.Log("Bytes AppendByte OK")
	}
}

func Test_BytesAppend(t *testing.T) {
	b := NewBytes()
	b.Append(Bytes{'a', 'b', 'c'})

	if b.ToString().ToNative() != "abc" {
		t.Error("Bytes Append Failed!")
	} else {
		t.Log("Bytes Append OK")
	}
}
