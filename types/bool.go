package types

func NewBool() Bool {
	return Bool(false)
}

func (this Bool) ToNative() bool {
	return bool(this)
}

func (this Bool) ToInt64() Int64 {
	if this {
		return 1
	} else {
		return 0
	}

}
func (this Bool) ToInt32() Int32 {
	if this {
		return 1
	} else {
		return 0
	}
}
func (this Bool) ToInt16() Int16 {
	if this {
		return 1
	} else {
		return 0
	}
}
func (this Bool) ToInt8() Int8 {
	if this {
		return 1
	} else {
		return 0
	}
}
func (this Bool) ToUInt64() UInt64 {
	if this {
		return 1
	} else {
		return 0
	}
}
func (this Bool) ToUInt32() UInt32 {
	if this {
		return 1
	} else {
		return 0
	}
}
func (this Bool) ToUInt16() UInt16 {
	if this {
		return 1
	} else {
		return 0
	}
}
func (this Bool) ToUInt8() UInt8 {
	if this {
		return 1
	} else {
		return 0
	}
}
func (this Bool) ToFloat64() Float64 {
	if this {
		return 1
	} else {
		return 0
	}
}
func (this Bool) ToFloat32() Float32 {
	if this {
		return 1
	} else {
		return 0
	}
}
func (this Bool) ToInt() Int {
	if this {
		return 1
	} else {
		return 0
	}
}
func (this Bool) ToUInt() UInt {
	if this {
		return 1
	} else {
		return 0
	}
}
func (this Bool) ToFloat() Float {
	if this {
		return 1
	} else {
		return 0
	}
}
func (this Bool) ToBool() Bool {
	return this
}

func (this Bool) ToJSON() JSON {
	return JSON(this)
}

func (this Bool) ToString() String {
	if this {
		return "true"
	} else {
		return "false"
	}
}
