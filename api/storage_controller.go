package api

import (
	"graph-db/internal/app/core/structs"
	"graph-db/internal/pkg/utils"
	"graph-db/internal/app/core"
)

type Storage interface {}

func CreateDatabase(dbTitle string, storageMode string) {
	err := core.InitDb(dbTitle, storageMode)
	utils.CheckError(err)
}

func SwitchDatabase(dbTitle string) {
	err := core.SwitchDb(dbTitle)
	utils.CheckError(err)
}

func DropDatabase(dbTitle string) {
	err := core.DropDb(dbTitle)
	utils.CheckError(err)
}

func CreateNode() (node structs.Node) {
	node = CreateNode()
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