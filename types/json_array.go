package types

import (
	"encoding/json"
)

func NewJSONArray() JSONArray {
	o := make(JSONArray, 0)
	return o
}

func (this *JSONArray) Push(item ...Any) {
	*this = append(*this, item...)
}

func (this *JSONArray) Pop() *JSON {
	index := this.GetLength() - 1
	json := &JSON{(*this)[index]}
	(*this) = (*this)[:index]
	return json
}

func (this JSONArray) IndexOf(item Any) Int {
	for i, v := range this {
		if v == item {
			return Int(i)
		}
	}
	return Int(-1)
}

func (this JSONArray) GetLength() Int {
	return Int(len(this))
}

func (this *JSONArray) Remove(index Int) {
	len := this.GetLength()
	if index <= 0 {
		*this = (*this)[1:]
	} else if index >= len-1 {
		*this = (*this)[:len]
	} else {
		*this = append((*this)[:index], (*this)[index+1:]...)
	}
}

func (this *JSONArray) RemoveItem(item Any) {
	index := this.IndexOf(item)
	this.Remove(index)
}

func (this JSONArray) ToJSON() JSON {
	return JSON{this}
}

// ToJSONString 转为json格式字符串
func (this JSONArray) ToJSONString() (String, error) {
	if bs, err := json.Marshal(this); err != nil {
		return "", err
	} else {
		return String(bs), nil
	}
}

func (this JSONArray) GetAsJSON(idx Int) *JSON {
	if idx < 0 || idx >= this.GetLength() {
		return nil
	}
	return &JSON{this[idx]}
}
