package structs

import "fmt"

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
	isUsed := boolToByteArray(n.isUsed)
	rel := int32ToByteArray(int32(n.getRelationship().id))
	prop := int32ToByteArray(int32(n.getProperty().id))
	label := int32ToByteArray(int32(n.getLabel().id))
	bs = append(isUsed, rel...)
	bs = append(bs, prop...)
	bs = append(bs, label...)
	return bs
}

func (n Node) fromBytes(bs []byte) {
	//todo
	if len(bs) != NodesSize {
		errorMessage := fmt.Sprintf("converter: wrong byte array length. expected array length is 13, actual length is %d", len(bs))
		panic(errorMessage)
	}
	n.isUsed, err = byteArrayToBool(bs[0:1])
	checkError(err)
	var id int32
	var rel Relationship
	id, err = byteArrayToInt32(bs[1:5])
	checkError(err)
	rel.id = int(id)
	n.relationship = &rel
	var prop Property
	id, err = byteArrayToInt32(bs[5:9])
	checkError(err)
	prop.id = int(id)
	n.property = prop
	var label Label
	id, err = byteArrayToInt32(bs[9:13])
	checkError(err)
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
