package core

import (
	"testing"
	"os"
	"github.com/KKhanda/graph-db/internal/pkg/utils"
)

var testFile *os.File

func init() {
	testFile, _ = os.Create("test.txt")
}

func TestIntToByteConversion(test *testing.T) {
	number := 123
	array := utils.Int32ToByteArray(int32(number))
	if len(array) != 4  {
		test.Errorf("Array length mismatch")
	}

	initialNumber, err := utils.ByteArrayToInt32(array)
	if int32(number) != initialNumber || err != nil {
		test.Errorf("Initial number mismatch")
	}
}

func TestByteToIntConversion(test *testing.T) {
	bs := []byte {53, 64, 71, 59}  // 994525237
	number, err := utils.ByteArrayToInt32(bs)
	if err != nil {
		test.Errorf("Got error during conversion")
	}

	initialBs := utils.Int32ToByteArray(number)
	if len(initialBs) != len(bs) {
		test.Errorf("Arrays length mismatch")
	}

	for i := 0; i < len(initialBs); i++ {
		if initialBs[i] != bs[i] {
			test.Errorf("Arrays values mismatch")
		}
	}
}

func TestFloat64ToByteArray(test *testing.T) {
	float  := float64(10)
	array := utils.Float64ToByteArray(float)
	if len(array) != 8 {
		test.Errorf("Array length mismatch")
	}

	initialFloat, err := utils.ByteArrayToFloat64(array)
	if err != nil {
		test.Errorf("Got error during conversion")
	}

	if float64(initialFloat) != float {
		test.Errorf("Initial float mismatch")
	}
}

func TestByteToFloatConversation(test *testing.T) {
	bs := []byte {53, 64, 71, 59, 55, 67, 43, 29}
	float, err := utils.ByteArrayToFloat64(bs)

	initialBs := utils.Float64ToByteArray(float)
	if len(initialBs) != len(bs) || err != nil {
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
	err := Write(testFile, 0, bs)
	if err != nil {
		test.Errorf("Error writing to file")
	}

	readBs := make([]byte, 4)
	Read(testFile, 0, readBs)
	for i := 0; i < len(bs); i++ {
		if bs[i] != readBs[i] {
			test.Errorf("Read values mismatch")
		}
	}

	bs = []byte{79, 11, 254, 98}
	err = Write(testFile, 1, bs)
	if err != nil {
		test.Errorf("Error writing to file")
	}

	readBs = make([]byte, 4)
	Read(testFile, 1, readBs)
	for i := 0; i < len(bs); i++ {
		if bs[i] != readBs[i] {
			test.Errorf("Read values mismatch")
		}
	}
}
