package core

import (
	"testing"
	"os"
)

var testFile *os.File

func init() {
	testFile, _ = os.Create("test.txt")
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
