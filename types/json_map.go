package types

import (
	"encoding/json"
)

func NewJSONMap() JSONMap {
	o := make(JSONMap, 0)
	return o
}

func (this JSONMap) ToJSON() *JSON {
	return &JSON{this}
}

func (this JSONMap) ToJSONString() (String, error) {
	if bs, err := json.Marshal(this); err != nil {
		return String(bs), nil
	} else {
		return "", err
	}
}
