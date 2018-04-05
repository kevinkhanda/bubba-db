package structs

type Relationship struct {
	id int
	isUsed bool
	node1 Node
	node2 Node
	title RelationshipTitle
	previousRelationship1 * Relationship
	previousRelationship2 * Relationship
	nextRelationship1 * Relationship
	nextRelationship2 * Relationship
	property Property
	isFirst bool
}

type RelationshipTitle struct {
	id int
	title string
	counter int
}
