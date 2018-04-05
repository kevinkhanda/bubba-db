package structs

import (
	"fmt"
	"graph-db/internal/pkg/utils"
	"graph-db/internal/app/core/globals"
	"graph-db/internal/app/core"
)

type Node struct {
	id       int
	isUsed   bool
	relationship * Relationship
	property Property
	label    Label
}

func (n Node) toBytes() []byte {
	//todo
	var bs []byte
	isUsed := utils.BoolToByteArray(n.isUsed)
	rel := utils.Int32ToByteArray(int32(n.getRelationship().id))
	prop := utils.Int32ToByteArray(int32(n.getProperty().id))
	label := utils.Int32ToByteArray(int32(n.getLabel().id))
	bs = append(isUsed, rel...)
	bs = append(bs, prop...)
	bs = append(bs, label...)
	return bs
}

func (n Node) fromBytes(bs []byte) {
	//todo
	var (
		id int32
		rel Relationship
		prop Property
		label Label
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
	n.property = prop
	id, err = utils.ByteArrayToInt32(bs[9:13])
	utils.CheckError(err)
	label.id = int(id)
	n.label = label
}

func (n Node) getRelationship() Relationship {
	//todo
	return Relationship{}
}

func (n Node) getProperty() Property {
	//todo
	return Property{}
}

func (n Node) getLabel() Label {
	//todo
	return Label{}
}

func (n Node) write()  {
	//todo
	offset := globals.NodesSize * n.id
	bs := n.toBytes()
	core.Write(globals.NodesStore, offset, bs)
}

func (n Node) read() {
	//todo
}

func (n Node) Create() {
	//todo
	var id int
	id, err = core.ReadId(globals.NodesId)
	utils.CheckError(err)
	n.id = id
	n.write()
}

type Label struct {
	id int
	isUsed bool
	numberOfLabels int
	labelNames [5]LabelTitle
}

type LabelTitle struct {
	id int
	title string
	counter int
}
