package core

import (
	"graph-db/internal/app/core/structs"
	"graph-db/internal/pkg/utils"
)

type Storage interface {}

func CreateDatabase(dbTitle string, storageMode string) {
	err := InitDb(dbTitle, storageMode)
	utils.CheckError(err)
}

func SwitchDatabase(dbTitle string) {
	err := SwitchDb(dbTitle)
	utils.CheckError(err)
}

func DropDatabase(dbTitle string) {
	err := DropDb(dbTitle)
	utils.CheckError(err)
}

func CreateNode() (node structs.Node) {
	node.Create()
	return node
}

func GetNode(id int) (node structs.Node) {
	//todo Recover from error or make node.Get() return error
	return node.Get(id)
}

func DeleteNode(id int) (err error) {
	var n structs.Node
	return n.Delete(id)
}

func CreateRelationship() (relationship structs.Relationship) {
	//relationship.Create()
	return relationship
}