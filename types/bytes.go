package types

func NewBytes() Bytes {
	return make(Bytes, 0)
}

func (this Bytes) ToNative() []byte {
	return []byte(this)
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

func (this *Bytes) AssignString(s String) {
	*this = s.ToBytes()
}
