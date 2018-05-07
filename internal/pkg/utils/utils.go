package utils

import (
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"math"
	"strings"
)

func CheckError(err error)  {
	if err != nil {
		log.Panic(err)
	}
}

// Int32ToByteArray transforms int32 value to byte array using fixed length encoding
// Returns byte array of size 4
// Reversed conversion is possible only with ByteArrayToInt32 function
func Int32ToByteArray(number int32) []byte {
	bs := make([]byte, 4)
	unsigned := uint32(number)
	binary.LittleEndian.PutUint32(bs, unsigned)
	return bs
}

// ByteArrayToInt32 transforms byte array of size 4 to int32
// Returns int32 value and error if length of byte array != 4
// Reversed conversion is possible only with Int32ToByteArray function
func ByteArrayToInt32(bs []byte) (int32, error) {
	if len(bs) != 4 {
		errorMessage := fmt.Sprintf("converter: wrong bs array length. Expected array length of 4, " +
			"actual length is %d", len(bs))
		return -1, errors.New(errorMessage)
	}
	unsigned := binary.LittleEndian.Uint32(bs)
	number := int32(unsigned)
	return number, nil
}

func Int8ToByteArray(number int8) []byte {
	bs := make([]byte, 1)
	bs[0] = byte(number)
	return bs
}

func ByteArrayToInt8(bs []byte) int8 {
	var number int8
	number = int8(bs[0])
	return number
}

// BoolToByteArray transforms bool value to byte array
// Returns byte array of size 1
// Reversed conversion is possible only with ByteArrayToBool function
func BoolToByteArray(flag bool) []byte {
	bs := make([]byte, 1)
	if flag {
		bs[0] = 0x01
	} else {
		bs[0] = 0x00
	}
	return bs
}

// ByteArrayToBool tansforms byte array of size 1 to bool
// Returns bool value and error if byte array size is not 1 or it contains bad data
// Reversed conversion is possible only with BoolToByteArray function
func ByteArrayToBool(bs []byte) (bool, error)  {
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
func StringToByteArray(s string) []byte {
	return []byte(s)
}

// ByteArrayToString transforms byte array to string
// Returns string of length equal to array size
func ByteArrayToString(bs []byte) string {
	return string(bs)
}

// Float64ToByteArray transforms float64 to byte array
// Returns byte array of size 8
func Float64ToByteArray(number float64) []byte {
	bits := math.Float64bits(number)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)
	return bytes
}

// ByteArrayToFloat64 transforms byte array of size 8 to float64
// Returns float64 value and error if byte array size is not 8
func ByteArrayToFloat64(bs []byte) (float64, error) {
	if len(bs) != 8 {
		errorMessage := fmt.Sprintf("converter: wrong bs array length. Expected array length of 8, " +
			"actual length is %d", len(bs))
		return -1, errors.New(errorMessage)
	}
	bits := binary.LittleEndian.Uint64(bs)
	float := math.Float64frombits(bits)
	return float, nil
}

// If string parameter length is less than requiredLength then '#' character is added to the end of string
func AddStopCharacter(str string, requiredLength int) string {
	for len(str) < requiredLength {
		str += "#"
	}

	return str
}

// Returns string slice until the '#' character
func RemoveStopCharacter(str string) string {
	i := strings.Index(str, "#")
	if i > -1 {
		return str[:i]
	} else {
		return str
	}
}