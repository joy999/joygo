package types

func NewJSONMap() JSONMap {
	o := make(JSONMap, 0)
	return o
}

func NewJSONArray() JSONArray {
	o := make(JSONArray, 0)
	return o
}

func (this JSONArray) ToJSON() JSON {
	return JSON(this)
}

func (this JSONMap) ToJSON() JSON {
	return JSON(this)
}
