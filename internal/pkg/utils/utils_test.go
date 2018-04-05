package utils

import "testing"

func TestInt32ToByteConversion(test *testing.T) {
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

func TestByteToInt32Conversion(test *testing.T) {
	bs := []byte {53, 64, 71, 59}  // 994525237
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

func TestFloat64ToByteArrayConversion(test *testing.T) {
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

func TestByteToFloat64Conversion(test *testing.T) {
	bs := []byte {53, 64, 71, 59, 55, 67, 43, 29}
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
