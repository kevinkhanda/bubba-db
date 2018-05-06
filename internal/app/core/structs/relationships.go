package structs

import (
	"graph-db/internal/app/core/globals"
	"graph-db/internal/pkg/utils"
	"fmt"
)

type Relationship struct {
	id int
	isUsed bool
	isWritten bool
	node1 *Node
	node2 *Node
	title *RelationshipTitle
	previousRelationship1 *Relationship
	previousRelationship2 *Relationship
	nextRelationship1 *Relationship
	nextRelationship2 *Relationship
	property *Property
	isFirst bool

	byteString []byte
}

type RelationshipTitle struct {
	id int
	title string
	counter int
}

func (r *Relationship) Create() {
	id, err := globals.FileHandler.ReadId(globals.RelationshipsId)
	utils.CheckError(err)
	r.id = id
	r.isUsed = true
	r.isWritten = false
	r.isFirst = false
	r.write()
}

func (r *Relationship) Delete(id int) (err error) {
	bs := make([]byte, globals.RelationshipsSize)
	bs[0] = utils.BoolToByteArray(false)[0]
	err = globals.FileHandler.FreeId(globals.RelationshipsId, id)
	if err != nil {
		return err
	}

	offset := globals.RelationshipsSize * id
	err = globals.FileHandler.Write(globals.RelationshipsStore, offset, bs)
	return err
}

func (r *Relationship) Get(id int) Relationship {
	r.id = id
	r.read()
	r.isWritten = true
	return *r
}

func (r *Relationship) GetId() int {
	return r.id
}

func (r *Relationship) getNode(start int, end int) *Node {
	var (
		nodeId int32
		bs = make([]byte, globals.RelationshipsSize)
	)
	if len(r.byteString) < 0 {
		offset := r.id * globals.RelationshipsSize
		err = globals.FileHandler.Read(globals.RelationshipsStore, offset, bs)
		utils.CheckError(err)
		r.byteString = bs
	}
	nodeId, err = utils.ByteArrayToInt32(r.byteString[start:end])
	utils.CheckError(err)
	if nodeId == -1 {
		return nil
	} else {
		var node Node
		node.id = int(nodeId)
		return &node
	}
}

func (r *Relationship) getRelationship(start int, end int) *Relationship {
	var (
		relationshipId int32
		bs = make([]byte, globals.RelationshipsSize)
	)
	if len(r.byteString) < 0 {
		offset := r.id * globals.RelationshipsSize
		err = globals.FileHandler.Read(globals.RelationshipsStore, offset, bs)
		utils.CheckError(err)
		r.byteString = bs
	}
	relationshipId, err = utils.ByteArrayToInt32(r.byteString[start:end])
	utils.CheckError(err)
	if relationshipId == -1 {
		return nil
	} else {
		var relationship Relationship
		relationship.id = int(relationshipId)
		return &relationship
	}
}

func (r *Relationship) GetFirstNode() *Node {
	if r.node1 != nil {
		return r.node1
	} else if !r.isWritten {
		return nil
	} else {
		node := r.getNode(1, 5)
		r.node1 = node
		return node
	}
}

func (r *Relationship) GetSecondNode() *Node {
	if r.node2 != nil {
		return r.node2
	} else if !r.isWritten {
		return nil
	} else {
		node := r.getNode(5, 9)
		r.node2 = node
		return node
	}
}

func (r *Relationship) GetTitle() *RelationshipTitle {
	if r.title != nil {
		return r.title
	} else if !r.isWritten {
		return nil
	} else {
		var (
			titleId int32
			err error
			bs = make([]byte, globals.RelationshipsSize)
		)
		if len(r.byteString) < 0 {
			offset := r.id * globals.RelationshipsSize
			err = globals.FileHandler.Read(globals.RelationshipsTitlesStore, offset, bs)
			utils.CheckError(err)
			r.byteString = bs
		}
		titleId, err = utils.ByteArrayToInt32(r.byteString[9:13])
		utils.CheckError(err)
		if titleId == -1 {
			return nil
		} else {
			var relationshipTitle RelationshipTitle
			relationshipTitle.id = int(titleId)
			r.title = &relationshipTitle
			return r.title
		}
	}
}

func (r *Relationship) GetFirstPreviousRelationship() *Relationship {
	if r.previousRelationship1 != nil {
		return r.previousRelationship1
	} else if !r.isWritten {
		return nil
	} else {
		relationship := r.getRelationship(13, 17)
		r.previousRelationship1 = relationship
		return relationship
	}
}

func (r *Relationship) GetSecondPreviousRelationship() *Relationship {
	if r.previousRelationship2 != nil {
		return r.previousRelationship2
	} else if !r.isWritten {
		return nil
	} else {
		relationship := r.getRelationship(17, 21)
		r.previousRelationship2 = relationship
		return relationship
	}
}

func (r *Relationship) GetFirstNextRelationship() *Relationship {
	if r.nextRelationship1 != nil {
		return r.nextRelationship1
	} else if !r.isWritten {
		return nil
	} else {
		relationship := r.getRelationship(21, 25)
		r.nextRelationship1 = relationship
		return relationship
	}
}

func (r *Relationship) GetSecondNextRelationship() *Relationship {
	if r.nextRelationship2 != nil {
		return r.nextRelationship2
	} else if !r.isWritten {
		return nil
	} else {
		relationship := r.getRelationship(25, 29)
		r.nextRelationship2 = relationship
		return relationship
	}
}

func (r *Relationship) GetProperty() *Property {
	if r.property != nil {
		return r.property
	} else if !r.isWritten {
		return nil
	} else {
		var (
			propertyId int32
			err error
			bs = make([]byte, globals.RelationshipsSize)
		)
		if len(r.byteString) < 0 {
			offset := r.id * globals.RelationshipsSize
			err = globals.FileHandler.Read(globals.RelationshipsTitlesStore, offset, bs)
			utils.CheckError(err)
			r.byteString = bs
		}
		propertyId, err = utils.ByteArrayToInt32(r.byteString[29:33])
		utils.CheckError(err)
		if propertyId == -1 {
			return nil
		} else {
			var property Property
			property.id = int(propertyId)
			r.property = &property
			return r.property
		}
	}
}

func (r *Relationship) toBytes() (bs []byte) {
	var (
		isUsed []byte
		isFirst []byte
		node1 *Node
		node2 *Node
		title *RelationshipTitle
		previousRelationship1 *Relationship
		previousRelationship2 *Relationship
		nextRelationship1 *Relationship
		nextRelationship2 *Relationship
		property *Property
	)
	node1 = r.GetFirstNode()
	node2 = r.GetSecondNode()
	title = r.GetTitle()
	previousRelationship1 = r.GetFirstPreviousRelationship()
	previousRelationship2 = r.GetSecondPreviousRelationship()
	nextRelationship1 = r.GetFirstNextRelationship()
	nextRelationship2 = r.GetSecondNextRelationship()
	property = r.GetProperty()

	isUsed = utils.BoolToByteArray(r.isUsed)
	node1Bs := utils.Int32ToByteArray(int32(IfNilAssignMinusOne(node1)))
	node2Bs := utils.Int32ToByteArray(int32(IfNilAssignMinusOne(node2)))
	titleBs := utils.Int32ToByteArray(int32(IfNilAssignMinusOne(title)))
	previousRelationship1Bs := utils.Int32ToByteArray(int32(IfNilAssignMinusOne(previousRelationship1)))
	previousRelationship2Bs := utils.Int32ToByteArray(int32(IfNilAssignMinusOne(previousRelationship2)))
	nextRelationship1Bs := utils.Int32ToByteArray(int32(IfNilAssignMinusOne(nextRelationship1)))
	nextRelationship2Bs := utils.Int32ToByteArray(int32(IfNilAssignMinusOne(nextRelationship2)))
	propertyBs := utils.Int32ToByteArray(int32(IfNilAssignMinusOne(property)))
	isFirst = utils.BoolToByteArray(r.isFirst)

	bs = append(isUsed, node1Bs...)
	bs = append(bs, node2Bs...)
	bs = append(bs, titleBs...)
	bs = append(bs, previousRelationship1Bs...)
	bs = append(bs, previousRelationship2Bs...)
	bs = append(bs, nextRelationship1Bs...)
	bs = append(bs, nextRelationship2Bs...)
	bs = append(bs, propertyBs...)
	bs = append(bs, isFirst...)

	return bs
}

func (r *Relationship) fromBytes(bs []byte) {
	var (
		id int32
		node1 Node
		node2 Node
		title RelationshipTitle
		previousRelationship1 Relationship
		previousRelationship2 Relationship
		nextRelationship1 Relationship
		nextRelationship2 Relationship
		property Property
	)
	if len(bs) != globals.RelationshipsSize {
		errorMessage := fmt.Sprintf("Converter: wrong relationships byte array length, expected 34, given %d", len(bs))
		panic(errorMessage)
	}
	r.isUsed, err = utils.ByteArrayToBool(bs[0:1])
	utils.CheckError(err)

	id, err = utils.ByteArrayToInt32(bs[1:5])
	utils.CheckError(err)
	node1.id = int(id)
	r.node1 = &node1
	id, err = utils.ByteArrayToInt32(bs[5:9])
	utils.CheckError(err)
	node2.id = int(id)
	r.node2 = &node2

	id, err = utils.ByteArrayToInt32(bs[9:13])
	utils.CheckError(err)
	title.id = int(id)
	r.title = &title

	id, err = utils.ByteArrayToInt32(bs[13:17])
	utils.CheckError(err)
	previousRelationship1.id = int(id)
	r.previousRelationship1 = &previousRelationship1
	id, err = utils.ByteArrayToInt32(bs[17:21])
	utils.CheckError(err)
	previousRelationship2.id = int(id)
	r.previousRelationship2 = &previousRelationship2

	id, err = utils.ByteArrayToInt32(bs[21:25])
	utils.CheckError(err)
	nextRelationship1.id = int(id)
	r.nextRelationship1 = &nextRelationship1
	id, err = utils.ByteArrayToInt32(bs[25:29])
	utils.CheckError(err)
	nextRelationship2.id = int(id)
	r.nextRelationship2 = &nextRelationship2

	id, err = utils.ByteArrayToInt32(bs[29:33])
	utils.CheckError(err)
	property.id = int(id)
	r.property = &property

	r.isFirst, err = utils.ByteArrayToBool(bs[33:34])
	utils.CheckError(err)
}

func (r *Relationship) read() {
	bs := make([]byte, globals.RelationshipsSize)
	offset := globals.RelationshipsSize * r.id
	err = globals.FileHandler.Read(globals.RelationshipsStore, offset, bs)
	utils.CheckError(err)
	r.fromBytes(bs)
}

func (r *Relationship) write() {
	offset := globals.RelationshipsSize * r.id
	bs := r.toBytes()
	err = globals.FileHandler.Write(globals.RelationshipsStore, offset, bs)
	utils.CheckError(err)
	r.isWritten = true
}

// Relationships Title
func (title *RelationshipTitle) GetId() int {
	return title.id
}

func WriteRelationshipsTitle(id int, title string, counter int) {
	offset := id * globals.RelationshipsTitlesSize
	bs := make([]byte, globals.RelationshipsTitlesSize)
	titleBs := utils.StringToByteArray(utils.AddStopCharacter(title, globals.RelationshipsTitlesSize - 4))
	for i := 0; i < len(titleBs); i++ {
		bs[i] = titleBs[i]
	}
	counterBs := utils.Int32ToByteArray(int32(counter))
	for i := 0; i < 4; i++ {
		bs[globals.RelationshipsTitlesSize - 4 + i] = counterBs[i]
	}
	err := globals.FileHandler.Write(globals.RelationshipsTitlesStore, offset, bs)
	utils.CheckError(err)
}

func DecreaseRelationshipTitleCounter(title string) {
	value := globals.RelationshipTitleMap[title]
	value.Counter--
	globals.RelationshipTitleMap[title] = value
	WriteRelationshipsTitle(value.Id, title, value.Counter)
	if globals.RelationshipTitleMap[title].Counter == 0 {
		delete(globals.RelationshipTitleMap, title)
	}
}

func AddRelationshipTitle(title string) *RelationshipTitle {
	value, present := globals.RelationshipTitleMap[title]
	if present {
		value.Counter++
		globals.RelationshipTitleMap[title] = value
	} else {
		id, err := globals.FileHandler.ReadId(globals.RelationshipsTitlesId)
		utils.CheckError(err)
		value = globals.MapValue{Counter: 1, Id: id}
		globals.RelationshipTitleMap[title] = value
	}
	WriteRelationshipsTitle(value.Id, title, value.Counter)
	return &RelationshipTitle{id: value.Id, title: title, counter: value.Counter}
}
