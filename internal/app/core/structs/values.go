package structs

type Value interface {
	get() interface{}
	set(value interface{})
}

type StringValue struct {
	id int
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
	id int
	value float64
}

func (doubleValue DoubleValue) get() interface{} {
	return doubleValue.value
}

func (doubleValue DoubleValue) set(value interface{}) {
	doubleValue.value = value.(float64)
}
