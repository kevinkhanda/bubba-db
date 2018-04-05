package utils

import "testing"

func TestInt32ToByte(test *testing.T) {
	number := 123
	array := Int32ToByteArray(int32(number))
	if len(array) != 4  {
		test.Errorf("Array length mismatch")
	}

	initialNumber, err := ByteArrayToInt32(array)
	if int32(number) != initialNumber || err != nil {
		test.Errorf("Initial number mismatch")
	}
}

func TestByteToInt32(test *testing.T) {
	bs := []byte {0x53, 0x64, 0x71, 0x59}
	number, err := ByteArrayToInt32(bs)
	if err != nil {
		test.Errorf("Got error during conversion")
	}

	initialBs := Int32ToByteArray(number)
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
	array := Float64ToByteArray(float)
	if len(array) != 8 {
		test.Errorf("Array length mismatch")
	}

	initialFloat, err := ByteArrayToFloat64(array)
	if err != nil {
		test.Errorf("Got error during conversion")
	}

	if float64(initialFloat) != float {
		test.Errorf("Initial float mismatch")
	}
}

func TestByteToFloat64(test *testing.T) {
	bs := []byte {0x53, 0x64, 0x71, 0x59, 0x55, 0x67, 0x43, 0x29}
	float, err := ByteArrayToFloat64(bs)

	initialBs := Float64ToByteArray(float)
	if len(initialBs) != len(bs) || err != nil {
		test.Errorf("Arrays length mismatch")
	}

	for i := 0; i < len(initialBs); i++ {
		if initialBs[i] != bs[i] {
			test.Errorf("Arrays values mismatch")
		}
	}
}

func TestBoolToByteArray(test *testing.T) {
	expectedBs := []byte{0x01}
	bs := BoolToByteArray(true)
	if len(bs) != 1 {
		test.Errorf("Array length mismatch")
	}

	if bs[0] != expectedBs[0] {
		test.Errorf("Array value mismatch")
	}
}

func TestByteArrayToBool(test *testing.T) {
	bs := []byte{0x01}
	boolValue, err := ByteArrayToBool(bs)
	if err != nil {
		test.Errorf("Got error during conversion")
	}

	if boolValue != true {
		test.Errorf("Bool value mismatch")
	}
}

func TestStringToByteArray(test *testing.T) {
	str := "test"
	expectedBs := []byte{0x74, 0x65, 0x73, 0x74}
	bs := StringToByteArray(str)
	if len(bs) != len(str) {
		test.Errorf("Array length mismatch")
	}

	for i := 0; i < len(bs); i++ {
		if bs[i]  != expectedBs[i] {
			test.Errorf("Array values mismatch")
		}
	}
}

func TestByteArrayToString(test *testing.T) {
	bs := []byte{0x74, 0x65, 0x73, 0x74}
	expectedStr := "test"
	str := ByteArrayToString(bs)
	if len(str) != len(bs) {
		test.Errorf("String length mismatch")
	}

	if str != expectedStr {
		test.Errorf("String value mismatch")
	}
}