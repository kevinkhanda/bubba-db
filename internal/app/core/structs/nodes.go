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
	labelNames []*LabelTitle
}

func (l Label) GetNumberOfLabels() int {
	return l.numberOfLabels
}

func (l Label) fromBytes(bs []byte) {
	var (
		id int32
		numberOfLabels int32
		title []LabelTitle
	)

	if len(bs) != globals.LabelsSize {
		errorMessage := fmt.Sprintf("converter: wrong byte array length. expected array length is 34, actual length is %d", len(bs))
		panic(errorMessage)
	}

	l.isUsed, err = utils.ByteArrayToBool(bs[0:1])
	utils.CheckError(err)
	numberOfLabels, err = utils.ByteArrayToInt32(bs[1:5])
	utils.CheckError(err)
	title = make([]LabelTitle, numberOfLabels)
	l.numberOfLabels = int(numberOfLabels)
	for i := 0; i < int(numberOfLabels); i++ {
		id, err = utils.ByteArrayToInt32(bs[5 + i * 4:5 + (i + 1) * 4])
		utils.CheckError(err)
		title[i].id = int(id)
		l.labelNames[i] = &title[i]
	}


}

func (l Label) toBytes() (bs []byte) {
	var (
		titleBs []byte
		numberBs []byte
		titles []*LabelTitle
	)

	titles = l.GetLabelNames()
	isUsed := utils.BoolToByteArray(l.isUsed)
	numberBs = utils.Int32ToByteArray(int32(l.numberOfLabels))
	bs = append(isUsed, numberBs...)
	for i := 0; i < l.numberOfLabels; i++ {
		titleBs = utils.Int32ToByteArray(int32(titles[i].GetId()))
		bs = append(bs, titleBs...)
	}

	//Need to append some stuff to make bs of needed size
	//numberBs was chosen because it was already created and has suitable size
	for len(bs) < globals.LabelsSize {
		bs = append(bs, numberBs...)
	}
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

func (l Label) GetLabelNames() []*LabelTitle  {
	if l.numberOfLabels == 0 || !l.isWritten {
		return nil
	} else if l.labelNames[0] != nil {
		return l.labelNames
	} else {
		var id int32
		offset := l.id * globals.LabelsSize
		bs := make([]byte, globals.LabelsSize)
		err := globals.FileHandler.Read(globals.LabelsStore, offset, bs)
		utils.CheckError(err)
		titles := make([]*LabelTitle, l.numberOfLabels)
		for i := 0; i < l.numberOfLabels; i++ {
			id, err = utils.ByteArrayToInt32(bs[5 + i * 4:5 + (i + 1) * 4])
			utils.CheckError(err)
			titles[i].id = int(id)
			l.labelNames[i] = titles[i]
		}
		return l.labelNames
	}

}

func (l Label) AddLabelName(title string) (err error) {
	if l.numberOfLabels == 5 {
		err = errors.New("Already max amount of labels")

	} else {
		l.labelNames[l.numberOfLabels] = AddLabelTitle(title)
		l.numberOfLabels++
		l.write()
	}
	return err
}

func (l Label) RemoveLabelName(id int) (err error)  {
	if l.numberOfLabels == 0 {
		err = errors.New("There is no such label")
		return err
	}
	_ = l.GetLabelNames()
	for i := 0; i < l.numberOfLabels; i++ {
		if (l.labelNames[i].GetId() == id) {
			title := l.labelNames[i].title
			if i == l.numberOfLabels - 1 {
				l.numberOfLabels--
				l.write()
			} else {
				l.labelNames[i] = l.labelNames[l.numberOfLabels - 1]
				l.numberOfLabels--;
				l.write()
			}
			DecreaseLabelTitleCounter(title)
			return nil
		}
	}
	err = errors.New("There is no such label")
	return err
}

func (l Label) Get(id int) Label {
	l.id = id
	l.labelNames = make([]*LabelTitle, 5)
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
	l.labelNames = make([]*LabelTitle, 5)
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

func (lt LabelTitle) GetId() int  {
	return lt.id
}

func  WriteLabelTitle(id int, title string, counter int)  {
	offset := id * globals.LabelsTitlesSize
	bs := make([]byte, globals.LabelsTitlesSize)
	titleBs := utils.StringToByteArray(utils.AddStopCharacter(title, globals.LabelsTitlesSize - 4))
	for i := 0; i < len(titleBs); i++ {
		bs[i] = titleBs[i]
	}
	counterBs := utils.Int32ToByteArray(int32(counter))
	for j := 0; j < 4; j++ {
		bs[globals.LabelsTitlesSize - 4 + j] = counterBs[j]
	}
	err := globals.FileHandler.Write(globals.LabelsTitlesStore, offset, bs)
	utils.CheckError(err)
}

func DecreaseLabelTitleCounter(title string)  {
	value := globals.LabelTitleMap[title]
	value.Counter--
	globals.LabelTitleMap[title] = value
	WriteLabelTitle(globals.LabelTitleMap[title].Id, title, globals.LabelTitleMap[title].Counter)
	if globals.LabelTitleMap[title].Counter == 0 {
		delete(globals.LabelTitleMap, title)
	}
}


func AddLabelTitle(title string) *LabelTitle {
	_, present := globals.LabelTitleMap[title]
	if present {
		value := globals.LabelTitleMap[title]
		value.Counter++
		globals.LabelTitleMap[title] = value
		WriteLabelTitle(globals.LabelTitleMap[title].Id, title, globals.LabelTitleMap[title].Counter)

		return &LabelTitle{id: globals.LabelTitleMap[title].Id, title: title, counter: globals.LabelTitleMap[title].Counter}
	} else {
		var value globals.MapValue
		value.Counter = 1
		id, err := globals.FileHandler.ReadId(globals.LabelsTitlesId)
		utils.CheckError(err)
		value.Id = id
		globals.LabelTitleMap[title] = value
		WriteLabelTitle(globals.LabelTitleMap[title].Id, title, globals.LabelTitleMap[title].Counter)
		return &LabelTitle{id: id, title: title, counter: 1}
	}
}
