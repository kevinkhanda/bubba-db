package structs

import (
	"reflect"
	"os"
	"graph-db/internal/app/core/globals"
	"fmt"
)

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
		}
		if reflect.TypeOf(value) == reflect.TypeOf(RelationshipTitle{}) {
			value := value.(RelationshipTitle)
			return value.id
		}
		return -1
	} else {
		return -1
	}
}

func GetValueFile(valueType int8) *os.File {
	switch valueType {
		case globals.INTEGER:
			return nil
		case globals.DOUBLE:
			return globals.DoubleStore
		case globals.STRING:
			return globals.StringStore
		default:
			errorMessage := fmt.Sprintf("Such type does not exist. This should never happen.")
			panic(errorMessage)
	}
}
