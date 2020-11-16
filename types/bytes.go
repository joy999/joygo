package types

func NewBytes() Bytes {
	return make(Bytes, 0)
}

func (this Bytes) ToString() String {
	return String(this)
}

func (this *Bytes) AppendByte(b Byte) {
	*this = append(*this, byte(b))
}

func (this *Bytes) Append(bs Bytes) {
	*this = append(*this, bs...)
}
