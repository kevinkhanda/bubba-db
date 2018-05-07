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
	InitDb("test_db", "local")
	InitDb("test_db2", "local")
	bs1 := make([]byte, globals.LabelsTitlesSize)
	bs2 := make([]byte, globals.LabelsTitlesSize)
	titleBs1 := utils.StringToByteArray("deleted#")
	titleBs2 := utils.StringToByteArray("present#")
	for i := 0; i < len(titleBs1); i++ {
		bs1[i] = titleBs1[i]
	}

	for i := 0; i < len(titleBs2); i++ {
		bs2[i] = titleBs2[i]
	}

	counterBs1 := utils.Int32ToByteArray(int32(0))
	counterBs2 := utils.Int32ToByteArray(int32(7))
	for j := 0; j < 4; j++ {
		bs1[globals.LabelsTitlesSize - 4 + j] = counterBs1[j]
	}

	for j := 0; j < 4; j++ {
		bs2[globals.LabelsTitlesSize - 4 + j] = counterBs2[j]
	}

	err := fh.Write(globals.LabelsTitlesStore, 0 * globals.LabelsTitlesSize, bs1, 0)
	if err != nil {
		test.Errorf("Error writing to file")
	}

	err = fh.Write(globals.LabelsTitlesStore, 1 * globals.LabelsTitlesSize, bs2, 1)
	if err != nil {
		test.Errorf("Error writing to file")
	}

	if globals.NodesId.Name() != "databases/test_db2/storage/nodes/id/nodes.id" {
		test.Errorf("File pointers mismatch")
	}

	SwitchDb("test_db")
	_, present := globals.LabelTitleMap["present"]
	if present {
		test.Errorf("Unexpected presence of map entry")
	}

	if globals.NodesId.Name() != "databases/test_db/storage/nodes/id/nodes.id" {
		test.Errorf("File pointers mismatch")
	}

	SwitchDb("test_db2")
	_, present = globals.LabelTitleMap["deleted"]
	if present {
		test.Errorf("Unexpected presence of map entry")
	}

	value, present := globals.LabelTitleMap["present"]
	if !present {
		test.Errorf("Unexpected absence of map entry")
	} else {
		if value.Id != 1 {
			test.Errorf("Id value mismatch")
		}

		if value.Counter != 7 {
			test.Errorf("Counter value mismatch")
		}
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
	err = fh.Write(testFile, 0, bs, 0)
	if err != nil {
		test.Errorf("Error writing to file")
	}

	readBs := make([]byte, 4)
	fh.Read(testFile, 0, readBs, 0)
	for i := 0; i < len(bs); i++ {
		if bs[i] != readBs[i] {
			test.Errorf("Read values mismatch")
		}
	}

	bs = []byte{79, 11, 254, 98}
	err = fh.Write(testFile, 1, bs, 1)
	if err != nil {
		test.Errorf("Error writing to file")
	}

	readBs = make([]byte, 4)
	fh.Read(testFile, 1, readBs, 1)
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
	}
}

func TestNodeCreate(test *testing.T) {
	fh.InitDatabaseStructure("test_db")
	var (
		bsExpected []byte
	)
	structs.CreateNode()
	bs := make([]byte, globals.NodesSize)
	bsId := utils.Int32ToByteArray(-1)
	bsExpected = append(utils.BoolToByteArray(true), bsId...)
	bsExpected = append(bsExpected, bsId...)
	bsExpected = append(bsExpected, bsId...)
	err := fh.Read(globals.NodesStore, 0, bs, 0)
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
	n = *structs.CreateNode()
	n.Delete(n.GetId())
	bs := make([]byte, globals.NodesSize)
	bsExpected := make([]byte, globals.NodesSize)
	bsExpected[0] = utils.BoolToByteArray(false)[0]
	err := fh.Read(globals.NodesStore, 0, bs, 0)
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
	err = fh.Write(globals.NodesStore, 0, bs, 0)
	if err != nil {
		test.Errorf("Error writing to file")
	}
	n = *n.Get(0)
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

func TestLabelCreate(test *testing.T) {
	var (
		bsExpected []byte
	)
	InitDb("test_db", "local")
	structs.CreateLabel()
	bs := make([]byte, globals.LabelsSize)
	bsNumber := utils.Int32ToByteArray(0)
	bsExpected = append(utils.BoolToByteArray(true), bsNumber...)
	err := fh.Read(globals.LabelsStore, 0, bs, 0)
	if err != nil {
		test.Errorf("Error reading a file")
		println(err.Error())
	}

	for i := 0; i < len(bsExpected); i++ {
		if bsExpected[i] != bs[i] {
			test.Errorf("Read value mismatch")
		}
	}

	DropDb("test_db")
}

//TODO: Test id availability after deletion for all delete methods
func TestLabelDelete(test *testing.T) {
	InitDb("test_db", "local")
	var l structs.Label
	l = *structs.CreateLabel()
	l.Delete(l.GetId())
	bs := make([]byte, globals.LabelsSize)
	bsExpected := make([]byte, globals.LabelsSize)
	bsExpected[0] = utils.BoolToByteArray(false)[0]
	err := fh.Read(globals.LabelsStore, 0, bs, 0)
	if err != nil {
		test.Errorf("Error reading from LabelStore")
	}

	for i := 0; i < len(bs); i++ {
		if bs[i] != bsExpected[i] {
			test.Errorf("Read values mismatch")
		}
	}

	DropDb("test_tb")
}

func TestLabelGet(test *testing.T)  {
	var(
		l structs.Label
		numberBs, firstBs, secondBs, emptyBs,
		bs []byte
	)
	InitDb("test_db", "local")
	numberBs = utils.Int32ToByteArray(2)
	firstBs = utils.Int32ToByteArray(10)
	secondBs = utils.Int32ToByteArray(20)
	emptyBs = make([]byte, 4)
	bs = append(utils.BoolToByteArray(true), numberBs...)
	bs = append(bs, firstBs...)
	bs = append(bs, secondBs...)
	bs = append(bs, emptyBs...)
	bs = append(bs, emptyBs...)
	bs = append(bs, emptyBs...)
	err = fh.Write(globals.LabelsStore, 0, bs, 0)
	if err != nil {
		test.Errorf("Error writing to file")
	}

	l = l.Get(0)
	if l.GetNumberOfLabels() != 2 {
		test.Errorf("Label number mismatch")
	}

	if l.GetLabelNames()[0].GetId() != 10 {
		test.Errorf("Id value mismatch")
	}

	if l.GetLabelNames()[1].GetId() != 20 {
		test.Errorf("Id value mismatch")
	}

	DropDb("test_tb")
}

func TestLabelNameMethods(test *testing.T)  {
	var l1, l2 structs.Label
	l1 = *structs.CreateLabel()
	l2 = *structs.CreateLabel()
	l1.AddLabelName("test")
	l2.AddLabelName("test")
	if l1.GetNumberOfLabels() != 1 {
		test.Errorf("Number of labels mismatch")
	}
	if l1.GetLabelNames()[0].GetTitle() != "test" {
		test.Errorf("Label title mismatch")
	}
	if l1.GetLabelNames()[0].GetId() != 0 {
		test.Errorf("Id value mismatch")
	}
	if l2.GetLabelNames()[0].GetId() != 0 {
		test.Errorf("Id value mismatch")
	}

	value, present := globals.LabelTitleMap["test"]
	if !present {
		test.Errorf("Absence of value in map")
	} else {
		if value.Id != 0 {
			test.Errorf("Id value mismatch")
		}
		if value.Counter != 2 {
			test.Errorf("Counter value mismatch")
		}
	}

	l2.RemoveLabelName("test")
	if l2.GetNumberOfLabels() != 0 {
		test.Errorf("Number of labels mismatch")
	}
	value, present = globals.LabelTitleMap["test"]
	if !present {
		test.Errorf("Absence of value in map")
	} else {
		if value.Id != 0 {
			test.Errorf("Id value mismatch")
		}
		if value.Counter != 1 {
			test.Errorf("Counter value mismatch")
		}
	}
	l1.RemoveLabelName("test")
	if l1.GetNumberOfLabels() != 0 {
		test.Errorf("Number of labels mismatch")
	}
	_, present = globals.LabelTitleMap["test"]
	if present {
		test.Errorf("Unexpected presence of value in map")
	}


}

func TestRemoveLabelName(test *testing.T) {

}

func TestCreateRelationship(test *testing.T) {
	var (
		bsExpected []byte
	)
	fh.InitDatabaseStructure("test_db")
	structs.CreateRelationship(false)
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
	err := fh.Read(globals.RelationshipsStore, 0, bs, 0)
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

func TestDeleteRelationship(test *testing.T) {
	fh.InitDatabaseStructure("test_db")
	var relationship structs.Relationship
	relationship = *structs.CreateRelationship(false)
	relationship.Delete(relationship.GetId())
	bs := make([]byte, globals.RelationshipsSize)
	bsExpected := make([]byte, globals.RelationshipsSize)
	bsExpected[0] = utils.BoolToByteArray(false)[0]
	err := fh.Read(globals.RelationshipsStore, 0, bs, 0)
	if err != nil {
		test.Errorf("Error reading from RelationshipStore")
	}
	for i := 0; i < len(bs); i++ {
		if bs[i] != bsExpected[i] {
			test.Errorf("Read values mismatch")
		}
	}
	fh.DropDatabase("test_db")
}

func TestGetRelationship(test *testing.T) {
	var (
		relationship structs.Relationship
		node1Id, node2Id, titleId, prevRel1Id, prevRel2Id,
		nextRel1Id, nextRel2Id, propertyId, bs []byte
	)
	fh.InitDatabaseStructure("test_db")
	node1Id = utils.Int32ToByteArray(10)
	node2Id = utils.Int32ToByteArray(20)
	titleId = utils.Int32ToByteArray(30)
	prevRel1Id = utils.Int32ToByteArray(40)
	prevRel2Id = utils.Int32ToByteArray(-1)
	nextRel1Id = utils.Int32ToByteArray(10)
	nextRel2Id = utils.Int32ToByteArray(15)
	propertyId = utils.Int32ToByteArray(1)

	bs = append(utils.BoolToByteArray(true), node1Id...)
	bs = append(bs, node2Id...)
	bs = append(bs, titleId...)
	bs = append(bs, prevRel1Id...)
	bs = append(bs, prevRel2Id...)
	bs = append(bs, nextRel1Id...)
	bs = append(bs, nextRel2Id...)
	bs = append(bs, propertyId...)
	bs = append(bs, utils.BoolToByteArray(false)...)

	err = fh.Write(globals.RelationshipsStore, 0, bs, 0)
	if err != nil {
		test.Errorf("Error writing to file")
	}
	relationship = *relationship.Get(0)
	if relationship.GetId() != 0 {
		test.Errorf("Id value mismatch")
	}

	if relationship.GetFirstNode().GetId() != 10 {
		test.Errorf("First node id value mismatch")
	}
	if relationship.GetSecondNode().GetId() != 20 {
		test.Errorf("Second node id value mismatch")
	}
	if relationship.GetTitle().GetId() != 30 {
		test.Errorf("Title id value mismatch")
	}
	if relationship.GetFirstPreviousRelationship().GetId() != 40 {
		test.Errorf("First previous relationship id value mismatch")
	}
	if relationship.GetSecondPreviousRelationship().GetId() != -1 {
		test.Errorf("Second previous relationship id value mismatch")
	}
	if relationship.GetFirstNextRelationship().GetId() != 10 {
		test.Errorf("First next relationship id value mismatch")
	}
	if relationship.GetSecondNextRelationship().GetId() != 15 {
		test.Errorf("Second next relationship id value mismatch")
	}
	if relationship.GetProperty().GetId() != 1 {
		test.Errorf("Property id value mismatch")
	}

	fh.DropDatabase("test_db")
}
