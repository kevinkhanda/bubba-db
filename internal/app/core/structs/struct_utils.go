package structs

import (
	"reflect"
	"os"
	"graph-db/internal/app/core/globals"
	"fmt"
)

func IfNilAssignMinusOne(value interface{}) int {
	if !reflect.ValueOf(value).IsNil() {
		if reflect.TypeOf(value).String() == reflect.TypeOf(&Node{}).String() {
			value := value.(*Node)
			return value.id
		}
		if reflect.TypeOf(value).String() == reflect.TypeOf(&Relationship{}).String() {
			println("Came here")
			value := value.(*Relationship)
			return value.id
		}
		if reflect.TypeOf(value).String() == reflect.TypeOf(&Property{}).String() {
			value := value.(*Property)
			return value.id
		}
		if reflect.TypeOf(value).String() == reflect.TypeOf(&Label{}).String() {
			value := value.(*Label)
			return value.id
		}
		if reflect.TypeOf(value).String() == reflect.TypeOf(&RelationshipTitle{}).String() {
			value := value.(*RelationshipTitle)
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
