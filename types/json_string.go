package types

import (
	"encoding/json"
	"errors"
)

func NewJSONString(s String) *JSONString {
	return &JSONString{s}
}

func (this JSONString) DecodeToJSON() (*JSON, error) {
	s := this.String
	if s.Match(`^\d+$`) { //整数
		var d Int64
		if err := json.Unmarshal(s.ToBytes(), &d); err != nil {
			return nil, err
		} else {
			return &JSON{d}, nil
		}
	}

	if s.Match("^\\\"[.\\w\\W]*\\\"$") {
		var str String
		if err := json.Unmarshal(s.ToBytes(), &str); err != nil {
			return nil, err
		} else {
			return &JSON{str}, nil
		}
	}

	if s.Match("^\\[[.\\w\\W]*\\]$") {
		var arr JSONArray
		if err := json.Unmarshal(s.ToBytes(), &arr); err != nil {
			return nil, err
		} else {
			return &JSON{arr}, nil
		}
	}

	if s.Match("^\\{[.\\w\\W]*\\}$") {
		var m JSONMap
		if err := json.Unmarshal(s.ToBytes(), &m); err != nil {
			return nil, err
		} else {
			return &JSON{m}, nil
		}
	}

	return nil, errors.New("不能解析的格式")
}
