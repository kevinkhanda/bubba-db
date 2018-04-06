package core

import (
	"testing"
	"os"
	"fmt"
	"io/ioutil"
	"graph-db/internal/app/core/globals"
)
var fh FileHandler

func init()  {
	globals.FileHandler = fh
}

func TestFileReadWrite(test *testing.T) {
	testFile, err := os.Create("test.txt")
	if err != nil {
		test.Errorf("Error creating file")
	}

	defer testFile.Close()
	defer os.Remove(testFile.Name())
	bs := []byte{53, 57, 50, 54}
	err = fh.Write(testFile, 0, bs)
	if err != nil {
		test.Errorf("Error writing to file")
	}

	readBs := make([]byte, 4)
	fh.Read(testFile, 0, readBs)
	for i := 0; i < len(bs); i++ {
		if bs[i] != readBs[i] {
			test.Errorf("Read values mismatch")
		}
	}

	bs = []byte{79, 11, 254, 98}
	err = fh.Write(testFile, 1, bs)
	if err != nil {
		test.Errorf("Error writing to file")
	}

	readBs = make([]byte, 4)
	fh.Read(testFile, 1, readBs)
	for i := 0; i < len(bs); i++ {
		if bs[i] != readBs[i] {
			test.Errorf("Read values mismatch")
		}
	}
}

func TestIdReading(test *testing.T) {
	testFile, err := os.Create("test.txt")
	if err != nil {
		test.Errorf("Error creating file")
	}

	defer testFile.Close()
	defer os.Remove(testFile.Name())
	testFile.WriteString(fmt.Sprintf("%d\n%d", 12, 17))
	id, err := fh.ReadId(testFile)
	if err != nil {
		test.Errorf("Error in reading id")
	}

	if id != 12 {
		test.Errorf("Id value mismatch")
	}

	id, err = fh.ReadId(testFile)
	if err != nil {
		test.Errorf("Error in reading id")
	}

	if id != 17 {
		test.Errorf("Id value mismatch")
	}

	newId, err := ioutil.ReadFile(testFile.Name())
	if string(newId) != "18" {
		test.Errorf("New id was not written")
	}
}
