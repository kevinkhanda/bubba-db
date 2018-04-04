package core

import (
	"testing"
	"os"
)

var testFile *os.File

func init() {
	testFile, _ = os.Create("test.txt")
}

func TestIntToByteConversion(test *testing.T) {
	number := 123
	array := int32ToByteArray(int32(number))

	if len(array) != 4  {
		test.Errorf("Array length mismatch")
	}

	initialNumber, err := byteArrayToInt32(array)
	if int32(number) != initialNumber || err != nil {
		test.Errorf("Initial number mismatch")
	}
}

func TestByteToIntConversion(test *testing.T) {
	bs := []byte {53, 64, 71, 59}  // 994525237
	number, err := byteArrayToInt32(bs)
	if err != nil {
		test.Errorf("Got error during conversion")
	}

	initialBs := int32ToByteArray(number)
	if len(initialBs) != len(bs) {
		test.Errorf("Arrays length mismatch")
	}

	for i := 0; i < len(initialBs); i++ {
		if initialBs[i] != bs[i] {
			test.Errorf("Arrays values mismatch")
		}
	}
}

func TestFileReadWrite(test *testing.T) {
	bs := []byte{53, 57, 50, 54}
	err := write(testFile, 0, bs)
	if err != nil {
		test.Errorf("Error writing to file")
	}
	readBs := make([]byte, 4)
	read(testFile, 0, readBs)

	for i := 0; i < len(bs); i++ {
		if bs[i] != readBs[i] {
			test.Errorf("Read values mismatch")
		}
	}

	bs = []byte{79, 11, 254, 98}
	err = write(testFile, 1, bs)
	if err != nil {
		test.Errorf("Error writing to file")
	}
	readBs = make([]byte, 4)
	read(testFile, 1, readBs)

	for i := 0; i < len(bs); i++ {
		if bs[i] != readBs[i] {
			test.Errorf("Read values mismatch")
		}
	}
}
