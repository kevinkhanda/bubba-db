package core

import (
	"testing"
	"os"
	"fmt"
	"io/ioutil"
	"graph-db/internal/app/core/globals"
	"graph-db/internal/app/core/structs"
	"graph-db/internal/pkg/utils"
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

func TestFreeId(test *testing.T)  {
	var str []byte
	testFile, err := os.Create("test.txt")
	if err != nil {
		test.Errorf("Error creating file")
	}

	defer testFile.Close()
	defer os.Remove(testFile.Name())
	testFile.WriteString(fmt.Sprintf("%d\n%d", 12, 17))
	err = fh.FreeId(testFile, 10)
	if err != nil {
		test.Errorf("Error during id release")
	}

	str, err = ioutil.ReadFile(testFile.Name())
	if string(str) != "10\n12\n17" {
		test.Errorf("Read values mismatch")
	}

	err = fh.FreeId(testFile, 18)
	if err != nil {
		if err.Error() != "Bad id (specified id is out of range)" {
			test.Errorf("Error during id release")
		}
	} else {
		test.Errorf("Absence of expected error")
	}

	err = fh.FreeId(testFile, 12)
	if err != nil {
		if err.Error() != "Bad id (specified id is already free)" {
			test.Errorf("Error during id release")
		}
	} else {
		test.Errorf("Absence of expected error")
	}

	err = fh.FreeId(testFile, 15)
	if err != nil {
		test.Errorf("Error during id release")
	}

	str, err = ioutil.ReadFile(testFile.Name())
	if string(str) != "10\n12\n15\n17" {
		test.Errorf("Read values mismatch")
		println(string(str))
	}
}

func TestNodeCreate(test *testing.T) {
	fh.InitDatabaseStructure("test_db")
	var (
		n structs.Node
		bsExpected []byte
	)
	n.Create()
	bs := make([]byte, globals.NodesSize)
	bsId := utils.Int32ToByteArray(-1)
	bsExpected = append(utils.BoolToByteArray(true), bsId...)
	bsExpected = append(bsExpected, bsId...)
	bsExpected = append(bsExpected, bsId...)
	err := fh.Read(globals.NodesStore, 0, bs)
	if err != nil {
		test.Errorf("Error reading from NodeStore")
	}

	for i := 0; i < len(bs); i++ {
		if bs[i] != bsExpected[i] {
			test.Errorf("Read values mismatch")
		}
	}

	fh.DropDatabase("test_db")
}

func TestNodeDelete(test *testing.T) {
	fh.InitDatabaseStructure("test_db")
	var n structs.Node
	n.Create()
	n.Delete(n.GetId())
	bs := make([]byte, globals.NodesSize)
	bsExpected := make([]byte, globals.NodesSize)
	bsExpected[0] = utils.BoolToByteArray(false)[0]
	err := fh.Read(globals.NodesStore, 0, bs)
	if err != nil {
		test.Errorf("Error reading from NodeStore")
	}

	for i := 0; i < len(bs); i++ {
		if bs[i] != bsExpected[i] {
			test.Errorf("Read values mismatch")
		}
	}

	fh.DropDatabase("test_db")
}