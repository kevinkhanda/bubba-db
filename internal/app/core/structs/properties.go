package structs

type Property struct {
	id int
	isUsed bool
	nextProperty * Property
	title PropertyTitle
	valueType int
	value Value
}

func (p Property) GetId() int {
	return p.id
}

type PropertyTitle struct {
	id int
	title string
	counter int
}
