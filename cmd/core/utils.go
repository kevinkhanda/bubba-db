package core

import "log"
import "encoding/binary"
import (
	"errors"
	"fmt"
)

func checkError(err error)  {
	if err != nil {
		log.Panic(err)
	}
}

// int32ToByteArray transforms int32 value to byte array using fixed length encoding
// Returns byte array of size 4
// Reversed conversion is possible only with byteArrayToInt32 function
func int32ToByteArray(number int32) []byte {
	bs := make([]byte, 4)
	unsigned := uint32(number)
	binary.LittleEndian.PutUint32(bs, unsigned)
	return bs
}

// byteArrayToInt32 transforms byte array of size 4 to int32
// Returns int32 value and error if length of byte array != 4
// Reversed conversion is possible only with int32ToByteArray function
func byteArrayToInt32(byte []byte) (int32, error) {
	if len(byte) != 4 {
		errorMessage := fmt.Sprintf("wrong byte array length. expected array length is 4, actual length is %d", len(byte))
		return -1, errors.New(errorMessage)
	}
	unsigned := binary.LittleEndian.Uint32(byte)
	number := int32(unsigned)
	return number, nil
}