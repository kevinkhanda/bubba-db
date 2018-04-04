package core

import "log"
import "encoding/binary"
import (
	"errors"
	"fmt"
	"math"
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
func byteArrayToInt32(bs []byte) (int32, error) {
	if len(bs) != 4 {
		errorMessage := fmt.Sprintf("converter: wrong bs array length. Expected array length of 4, " +
			"actual length is %d", len(bs))
		return -1, errors.New(errorMessage)
	}
	unsigned := binary.LittleEndian.Uint32(bs)
	number := int32(unsigned)
	return number, nil
}

// boolToByteArray transforms bool value to byte array
// Returns byte array of size 1
// Reversed conversion is possible only with byteArrayToBool function
func boolToByteArray(flag bool) []byte {
	bs := make([]byte, 1)
	if flag {
		bs[0] = 0x01
	} else {
		bs[0] = 0x00
	}
	return bs
}

// byteArrayToBool tansforms byte array of size 1 to bool
// Returns bool value and error if byte array size is not 1 or it contains bad data
// Reversed conversion is possible only with boolToByteArray function
func byteArrayToBool(bs []byte) (bool, error)  {
	if len(bs) != 1 {
		errorMessage := fmt.Sprintf("converter: wrong byte array length. expected array length is 1, actual length is %d", len(bs))
		return false, errors.New(errorMessage)
	}
	if bs[0] == 0x00 {
		return false, nil
	} else if bs[0] == 0x01 {
		return true, nil
	} else {
		errorMessage := fmt.Sprintf("converter: byte array contains bad data")
		return false, errors.New(errorMessage)
	}
}

// StringToByteArray transforms string to byte array
// Returns byte array of size len(string)
func stringToByteArray(s string) []byte {
	return []byte(s)
}

// ByteArrayToString transforms byte array to string
// Returns string of length equal to array size
func byteArrayToString(bs []byte) string {
	return string(bs)
}

func Float64ToByteArray(number float64) []byte {
	bits := math.Float64bits(number)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)
	return bytes
}

func ByteArrayToFloat64(bs []byte) float64 {
	bits := binary.LittleEndian.Uint64(bs)
	float := math.Float64frombits(bits)
	return float
}