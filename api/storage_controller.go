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

func CreateNode(title string) (node structs.Node) {
	node = *structs.CreateNode()
	label := *structs.CreateLabel()
	label.AddLabelName(title)
	node.SetLabel(&label)

	return node
}

func CreateRelationship(firstNode *structs.Node, secondNode *structs.Node, title string) (relationship structs.Relationship) {
	flag := firstNode.GetRelationship() == nil
	relationship = *structs.CreateRelationship(flag)
	relationship.SetNode1(firstNode)
	relationship.SetNode2(secondNode)
	if flag {
		firstNode.SetRelationship(&relationship)
		relationship.SetPreviousRelationship1(nil)
	} else {
		lastRelationship1 := firstNode.GetRelationship()
		for true {
			if lastRelationship1.GetFirstNode().GetId() == firstNode.GetId() {
				if lastRelationship1.GetFirstNextRelationship() == nil {
					lastRelationship1.SetNextRelationship1(&relationship)
					relationship.SetPreviousRelationship1(lastRelationship1)
					break
				} else {
					lastRelationship1 = lastRelationship1.GetFirstNextRelationship()
				}
			} else {
				if lastRelationship1.GetSecondNextRelationship() == nil {
					lastRelationship1.SetNextRelationship2(&relationship)
					relationship.SetPreviousRelationship1(lastRelationship1)
					break
				} else {
					lastRelationship1 = lastRelationship1.GetSecondNextRelationship()
				}
			}
		}
	}

	if secondNode.GetRelationship() == nil {
		secondNode.SetRelationship(&relationship)
		relationship.SetPreviousRelationship2(nil)
	} else {
		lastRelationship2 := secondNode.GetRelationship()
		for true {
			if lastRelationship2.GetFirstNode().GetId() == secondNode.GetId() {
				if lastRelationship2.GetFirstNextRelationship() == nil {
					lastRelationship2.SetNextRelationship1(&relationship)
					relationship.SetPreviousRelationship2(lastRelationship2)
					break
				} else {
					lastRelationship2 = lastRelationship2.GetFirstNextRelationship()
				}
			} else {
				if lastRelationship2.GetSecondNextRelationship() == nil {
					lastRelationship2.SetNextRelationship2(&relationship)
					relationship.SetPreviousRelationship2(lastRelationship2)
					break
				} else {
					lastRelationship2 = lastRelationship2.GetSecondNextRelationship()
				}
			}
		}
	}

	relTitle := structs.AddRelationshipTitle(title)
	relationship.SetTitle(relTitle)

	return relationship
}

func GetNode(id int) (node structs.Node) {
	//todo Recover from error or make node.Get() return error
	return node.Get(id)
}

func DeleteNode(id int) (err error) {
	var n structs.Node
	return n.Delete(id)
}

//
//func CreateRelationship() (relationship structs.Relationship) {
//	//relationship.Create()
//	return relationship
//}