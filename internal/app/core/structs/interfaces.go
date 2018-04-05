package structs

type Store interface {
	read()
	write()
}

type ByteRepresentable interface {
	toBytes() []byte
	fromBytes([]byte)
}
