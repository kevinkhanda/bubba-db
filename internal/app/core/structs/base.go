package structs

var err error

type Store interface {
	read()
	write()
}

type ByteRepresentable interface {
	toBytes() []byte
	fromBytes([]byte)
}