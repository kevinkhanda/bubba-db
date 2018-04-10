package structs

import "reflect"

func IfNilAssignMinusOne(value interface{}) int {
	if value != nil {
		if reflect.TypeOf(value) == reflect.TypeOf(Node{}) {
			value := value.(Node)
			return value.id
		}
		if reflect.TypeOf(value) == reflect.TypeOf(Relationship{}) {
			value := value.(Relationship)
			return value.id
		}
		if reflect.TypeOf(value) == reflect.TypeOf(Property{}) {
			value := value.(Property)
			return value.id
		}
		if reflect.TypeOf(value) == reflect.TypeOf(Label{}) {
			value := value.(Label)
			return value.id
		} else {
			return -1
		}
	} else {
		return -1
	}
}
