package structs

import (
	"fmt"
	"graph-db/internal/pkg/utils"
	"graph-db/internal/app/core/globals"
	"errors"
)

type Node struct {
	id       int
	isUsed   bool
	isWritten bool
	relationship *Relationship
	property *Property
	label    *Label
}

func (n Node) toBytes() (bs []byte) {
	var (
		rel *Relationship
		prop *Property
		label *Label
		relBs, propBs, labelBs []byte
	)
	isUsed := utils.BoolToByteArray(n.isUsed)
	rel = n.GetRelationship()
	id := IfNilAssignMinusOne(rel)
	relBs = utils.Int32ToByteArray(int32(id))

	prop = n.GetProperty()
	id = IfNilAssignMinusOne(prop)
	propBs = utils.Int32ToByteArray(int32(id))

	label = n.GetLabel()
	id = IfNilAssignMinusOne(label)
	labelBs = utils.Int32ToByteArray(int32(id))

	bs = append(isUsed, relBs...)
	bs = append(bs, propBs...)
	bs = append(bs, labelBs...)
	return bs
}

func (n Node) fromBytes(bs []byte) {
	var (
		id int32
		rel Relationship
		prop Property
		label Label

		err error
	)
	if len(bs) != globals.NodesSize {
		errorMessage := fmt.Sprintf("converter: wrong byte array length. expected array length is 13, actual length is %d", len(bs))
		panic(errorMessage)
	}
	n.isUsed, err = utils.ByteArrayToBool(bs[0:1])
	utils.CheckError(err)
	id, err = utils.ByteArrayToInt32(bs[1:5])
	utils.CheckError(err)
	rel.id = int(id)
	n.relationship = &rel
	id, err = utils.ByteArrayToInt32(bs[5:9])
	utils.CheckError(err)
	prop.id = int(id)
	n.property = &prop
	id, err = utils.ByteArrayToInt32(bs[9:13])
	utils.CheckError(err)
	label.id = int(id)
	n.label = &label
}

func (n Node) GetRelationship() *Relationship {
	if n.relationship != nil {
		return n.relationship
	} else if !n.isWritten {
		return nil
	} else {
		offset := n.id * globals.NodesSize
		bs := make([]byte, globals.NodesSize)
		globals.FileHandler.Read(globals.NodesStore, offset, bs)
		relId, err := utils.ByteArrayToInt32(bs[1:5])
		utils.CheckError(err)
		if relId == -1 {
			return nil
		} else {
			var rel Relationship
			rel.id = int(relId)
			n.relationship = &rel
			return n.relationship
		}
	}
}

func (n Node) GetId() int {
	return n.id
}

func (n Node) SetRelationship(rel *Relationship) {
	n.relationship = rel
	n.write()
}

func (n Node) GetProperty() *Property {
	if n.property != nil {
		return n.property
	} else if !n.isWritten {
		return nil
	} else {
		offset := n.id * globals.NodesSize
		bs := make([]byte, globals.NodesSize)
		globals.FileHandler.Read(globals.NodesStore, offset, bs)
		propId, err := utils.ByteArrayToInt32(bs[5:9])
		utils.CheckError(err)
		if propId == -1 {
			return nil
		} else {
			var prop Property
			prop.id = int(propId)
			n.property = &prop
			return n.property
		}
	}
}

func (n Node) SetProperty(prop *Property) {
	n.property = prop
	n.write()
}

func (n Node) GetLabel() *Label {
	if n.label != nil {
		return n.label
	} else if !n.isWritten {
		return nil
	} else {
		offset := n.id * globals.NodesSize
		bs := make([]byte, globals.NodesSize)
		globals.FileHandler.Read(globals.NodesStore, offset, bs)
		labelId, err := utils.ByteArrayToInt32(bs[9:13])
		utils.CheckError(err)
		if labelId == -1 {
			return nil
		} else {
			var label Label
			label.id = int(labelId)
			n.label = &label
			return n.label
		}
	}
}

func (n Node) SetLabel(label *Label) {
	n.label = label
	n.write()
}

func (n Node) write()  {
	offset := globals.NodesSize * n.id
	bs := n.toBytes()
	err := globals.FileHandler.Write(globals.NodesStore, offset, bs)
	utils.CheckError(err)
	n.isWritten = true
}

func (n Node) read() {
	bs := make([]byte, globals.NodesSize)
	offset := globals.NodesSize * n.id
	err := globals.FileHandler.Read(globals.NodesStore, offset, bs)
	utils.CheckError(err)
	n.fromBytes(bs)
}

func (n Node) Create() {
	id, err := globals.FileHandler.ReadId(globals.NodesId)
	utils.CheckError(err)
	n.id = id
	n.isUsed = true
	n.isWritten = false
	n.write()
}

func (n Node) Get(id int) Node {
	n.id = id
	n.read()
	n.isWritten = true	//Doesn't work if placed into read (-_-)
	return n
}

func (n Node) Delete(id int) (err error) {
	bs := make([]byte, globals.NodesSize)
	bs[0] = utils.BoolToByteArray(false)[0]
	err = globals.FileHandler.FreeId(globals.NodesId, id)
	if err != nil {
		return err
	}

	offset := globals.NodesSize * id
	err = globals.FileHandler.Write(globals.NodesStore, offset, bs)
	return err
}
////////////////////////////////////////////////////////////////////////////////////////////////////
type Label struct {
	id int
	isWritten bool
	isUsed bool
	numberOfLabels int
	labelNames [5]*LabelTitle
}

//todo Label from bytes
func (l Label) fromBytes(bs []byte) {

}

//todo Label to bytes
func (l Label) toBytes() (bs []byte) {
	return bs
}

func (l Label) write() {
	offset := globals.LabelsSize * l.id
	bs := l.toBytes()
	err := globals.FileHandler.Write(globals.LabelsStore, offset, bs)
	utils.CheckError(err)
	l.isWritten = true
}

func (l Label) read() {
	bs := make([]byte, globals.LabelsSize)
	offset := globals.LabelsSize * l.id
	err := globals.FileHandler.Read(globals.LabelsStore, offset, bs)
	utils.CheckError(err)
	l.fromBytes(bs)
}

//todo
func (l Label) GetLabelNames() []*LabelTitle  {
	return nil
}

func (l Label) AddLabelName(title *LabelTitle) (err error) {
	if l.numberOfLabels == 5 {
		err = errors.New("Already max amount of labels")

	} else {
		l.labelNames[l.numberOfLabels] = title
		l.numberOfLabels++
	}
	return err
}

//todo
func (l Label) RemoveLabelName() (err error)  {
	return nil
}

func (l Label) Get(id int) Label {
	l.id = id
	l.read()
	l.isWritten = true
	return l
}

func (l Label) Delete(id int) (err error) {
	bs := make([]byte, globals.LabelsSize)
	bs[0] = utils.BoolToByteArray(false)[0]
	err = globals.FileHandler.FreeId(globals.LabelsId, id)
	if err != nil {
		return err
	}

	offset := globals.LabelsSize * id
	err = globals.FileHandler.Write(globals.LabelsStore, offset, bs)
	return err
}

func (l Label) Create() {
	id, err := globals.FileHandler.ReadId(globals.LabelsId)
	utils.CheckError(err)
	l.id = id
	l.numberOfLabels = 0
	l.isUsed = true
	l.isWritten = false
	l.write()
}

func (l Label) GetId() int {
	return l.id
}
////////////////////////////////////////////////////////////////////////////////////////////
type LabelTitle struct {
	id int
	title string
	counter int
}
