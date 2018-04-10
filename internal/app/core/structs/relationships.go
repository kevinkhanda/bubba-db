package structs

type Relationship struct {
	id int
	isUsed bool
	node1 *Node
	node2 *Node
	title *RelationshipTitle
	previousRelationship1 *Relationship
	previousRelationship2 *Relationship
	nextRelationship1 *Relationship
	nextRelationship2 *Relationship
	property *Property
	isFirst bool
}

func (r Relationship) GetId() int {
	return r.id
}

type RelationshipTitle struct {
	id int
	title string
	counter int
}

//func (r Relationship) toBytes() (bs []byte) {
//
//}
//
//func (r Relationship) fromBytes() (bs []byte) {
//
//}
//
//func (r Relationship) write() {
//	offset := globals.RelationshipsSize * r.id
//	bs := r.toBytes()
//}
//
//func (r Relationship) Create() {
//	id, err := globals.FileHandler.ReadId(globals.RelationshipsId)
//	utils.CheckError(err)
//	r.id = id
//	r.write()
//}
