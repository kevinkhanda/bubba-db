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
	return node.Get(id)
}


func CreateRelationship() (relationship structs.Relationship) {
	//relationship.Create()
	return relationship
}