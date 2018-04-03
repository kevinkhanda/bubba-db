package core

const (
	INTEGER = 0
	DOUBLE = 1
	STRING = 2
)

type Store interface {
	read()
	write()
}

type Node struct {
	id       int
	isUsed   bool
	relation Relationship
	property Property
	label    Label
}

type Label struct {
	id int
	isUsed bool
	numberOfLabels int
	labelNames [5]LabelTitle
}

type LabelTitle struct {
	title string
	counter int
}

type Relationship struct {
	id int
	isUsed bool
	node1 Node
	node2 Node
	title RelationshipTitle
	previousRelationship1 Relationship
	previousRelationship2 Relationship
	nextRelationship1 Relationship
	nextRelationship2 Relationship
	property Property
	isFirst bool
}

type RelationshipTitle struct {
	title string
	counter int
}

type Property struct {
	id int
	isUsed bool
	nextProperty Property
	title PropertyTitle
	valueType int
	value Value
}

type PropertyTitle struct {
	title string
	counter int
}

type Value interface {
	get() interface{}
	set(value interface{})
}

type StringValue struct {
	value string
}

func (stringValue StringValue) get() interface{} {
	return stringValue.value
}

func (stringValue StringValue) set(value interface{}) {
	stringValue.value = value.(string)
}

type IntegerValue struct {
	value int
}

func (integerValue IntegerValue) get() interface{} {
	return integerValue.value
}

func (integerValue IntegerValue) set(value interface{}) {
	integerValue.value = value.(int)
}

type DoubleValue struct {
	value float64
}

func (doubleValue DoubleValue) get() interface{} {
	return doubleValue.value
}

func (doubleValue DoubleValue) set(value interface{}) {
	doubleValue.value = value.(float64)
}
