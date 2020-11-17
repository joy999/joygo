package types

import (
	"encoding/json"
)

func NewJSONMap() JSONMap {
	o := make(JSONMap, 0)
	return o
}

func (this *JSONMap) Set(key String, any Any) {
	(*this)[key] = any
}

func (this JSONMap) Get(key String) Any {

	if v, ok := this[key]; ok {
		return Any(v)
	} else {
		return nil
	}
}

func (this JSONMap) GetAsJSON(key String) *JSON {
	if v, ok := this[key]; ok {
		return &JSON{v}
	} else {
		return nil
	}
}

func (this JSONMap) ToJSON() JSON {
	return JSON{this}
}

func (this JSONMap) Exists(key String) Bool {
	if _, ok := this[key]; ok {
		return true
	} else {
		return false
	}
}

func (this *JSONMap) Remove(key String) {
	delete(*this, key)
}

func (this JSONMap) ToJSONString() (String, error) {
	if bs, err := json.Marshal(this); err == nil {
		return String(bs), nil
	} else {
		return "", err
	}
}

func (this *JSONMap) Clear() {
	*this = make(JSONMap, 0)
}
