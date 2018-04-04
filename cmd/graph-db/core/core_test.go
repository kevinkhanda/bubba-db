package core

import (
	"testing"
)

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
	bs := []byte {53, 64, 71, 59}
	number, err := byteArrayToInt32(bs)
	if err != nil {
		test.Errorf("Got error during conversion")
	}

	initialBS := int32ToByteArray(number)
	if len(initialBS) != len(bs) {
		test.Errorf("Arrays length mismatch")
	}
	for i := 0; i < len(initialBS); i++ {
		if initialBS[i] != bs[i] {
			test.Errorf("Arrays values mismatch")
		}
	}
}
