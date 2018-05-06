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

func TestInitDatabaseStructure(test *testing.T) {
	fh.InitDatabaseStructure("test_db")
	if globals.NodesId.Name() != "databases/test_db/storage/nodes/id/nodes.id" {
		test.Errorf("Expected file was not created")
	}

	fileData, err := ioutil.ReadFile(globals.NodesId.Name())
	if err != nil {
		test.Errorf("Error reading a file")
	}

	if string(fileData) != "0" {
		test.Errorf("File content mismatch")
	}
	fh.DropDatabase("test_db")
}

func TestSwitchDatabase(test *testing.T) {
	fh.InitDatabaseStructure("test_db")
	fh.InitDatabaseStructure("test_db2")
	if globals.NodesId.Name() != "databases/test_db2/storage/nodes/id/nodes.id" {
		test.Errorf("File pointers mismatch")
	}
	fh.SwitchDatabaseStructure("test_db")
	if globals.NodesId.Name() != "databases/test_db/storage/nodes/id/nodes.id" {
		test.Errorf("File pointers mismatch")
	}
	fh.DropDatabase("test_db")
	fh.DropDatabase("test_db2")
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

func TestNodeGet(test *testing.T) {
	var (
		n structs.Node
		relId, propId, labelId []byte
		bs []byte
	)
	fh.InitDatabaseStructure("test_db")
	relId = utils.Int32ToByteArray(10)
	propId = utils.Int32ToByteArray(20)
	labelId = utils.Int32ToByteArray(30)
	bs = append(utils.BoolToByteArray(true), relId...)
	bs = append(bs, propId...)
	bs = append(bs, labelId...)
	err = fh.Write(globals.NodesStore, 0, bs)
	if err != nil {
		test.Errorf("Error writing to file")
	}
	n = n.Get(0)
	if n.GetRelationship().GetId() != 10 {
		test.Errorf("Id value mismatch")
	}

	if n.GetProperty().GetId() != 20 {
		test.Errorf("Id value mismatch")
	}

	if n.GetLabel().GetId() != 30 {
		test.Errorf("Id value mismatch")
	}

	fh.DropDatabase("test_db")
}

func TestCreateRelationship(test *testing.T) {
	var (
		relationship structs.Relationship
		bsExpected []byte
	)
	fh.InitDatabaseStructure("test_db")
	relationship.Create()
	bs := make([]byte, globals.RelationshipsSize)
	bsId := utils.Int32ToByteArray(-1)
	bsExpected = append(utils.BoolToByteArray(true), bsId...)
	bsExpected = append(bsExpected, bsId...)
	bsExpected = append(bsExpected, bsId...)
	bsExpected = append(bsExpected, bsId...)
	bsExpected = append(bsExpected, bsId...)
	bsExpected = append(bsExpected, bsId...)
	bsExpected = append(bsExpected, bsId...)
	bsExpected = append(bsExpected, bsId...)
	bsExpected = append(bsExpected, utils.BoolToByteArray(false)...)
	err := fh.Read(globals.RelationshipsStore, 0, bs)
	if err != nil {
		test.Errorf("Error reading from RelationshipsStore")
	}
	for i := 0; i < len(bs); i++ {
		if bs[i] != bsExpected[i] {
			test.Logf("%b %b", bs[i], bsExpected[i])
			test.Errorf("Read values mismatch")
		}
	}
	fh.DropDatabase("test_db")
}