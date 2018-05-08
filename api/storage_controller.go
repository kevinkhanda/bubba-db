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

func CreateNode(title string) (node *structs.Node) {
	node = structs.CreateNode()
	label := *structs.CreateLabel()
	label.AddLabelName(title)
	node.SetLabel(&label)

	return node
}

func CreateRelationship(firstNode *structs.Node, secondNode *structs.Node, title string) (relationship *structs.Relationship) {
	flag := firstNode.GetRelationship() == nil
	relationship = structs.CreateRelationship(flag)
	relationship.SetNode1(firstNode)
	relationship.SetNode2(secondNode)
	if flag {
		firstNode.SetRelationship(relationship)
		relationship.SetPreviousRelationship1(nil)
	} else {
		lastRelationship1 := firstNode.GetRelationship()
		for true {
			if lastRelationship1.GetFirstNode().GetId() == firstNode.GetId() {
				if lastRelationship1.GetFirstNextRelationship() == nil {
					lastRelationship1.SetNextRelationship1(relationship)
					relationship.SetPreviousRelationship1(lastRelationship1)
					break
				} else {
					lastRelationship1 = lastRelationship1.GetFirstNextRelationship()
				}
			} else {
				if lastRelationship1.GetSecondNextRelationship() == nil {
					lastRelationship1.SetNextRelationship2(relationship)
					relationship.SetPreviousRelationship1(lastRelationship1)
					break
				} else {
					lastRelationship1 = lastRelationship1.GetSecondNextRelationship()
				}
			}
		}
	}

	if secondNode.GetRelationship() == nil {
		secondNode.SetRelationship(relationship)
		relationship.SetPreviousRelationship2(nil)
	} else {
		lastRelationship2 := secondNode.GetRelationship()
		for true {
			if lastRelationship2.GetFirstNode().GetId() == secondNode.GetId() {
				if lastRelationship2.GetFirstNextRelationship() == nil {
					lastRelationship2.SetNextRelationship1(relationship)
					relationship.SetPreviousRelationship2(lastRelationship2)
					break
				} else {
					lastRelationship2 = lastRelationship2.GetFirstNextRelationship()
				}
			} else {
				if lastRelationship2.GetSecondNextRelationship() == nil {
					lastRelationship2.SetNextRelationship2(relationship)
					relationship.SetPreviousRelationship2(lastRelationship2)
					break
				} else {
					lastRelationship2 = lastRelationship2.GetSecondNextRelationship()
				}
			}
		}
	}

	relTitle, err := structs.AddRelationshipTitle(title)
	utils.CheckError(err)
	relationship.SetTitle(relTitle)

	return relationship
}

func CreatePropertyForNode(node *structs.Node, title string, valueType int, value interface{}) (property *structs.Property){
	property = structs.CreateProperty()
	property.SetValueType(int8(valueType))
	property.SetValue(int8(valueType), value)
	propertyTitle, err := structs.AddPropertyTitle(title)
	utils.CheckError(err)
	property.SetTitle(propertyTitle)

	if node.GetProperty() == nil {
		node.SetProperty(property)
	} else {
		lastProperty := node.GetProperty()
		for true {
			if lastProperty.GetNextProperty() == nil {
				lastProperty.SetNextProperty(property)
				break
			} else {
				lastProperty = lastProperty.GetNextProperty()
			}
		}
	}

	return property
}

func CreatePropertyForRelationship(relationship *structs.Relationship, title string, valueType int, value interface{}) (property *structs.Property){
	property = structs.CreateProperty()
	property.SetValueType(int8(valueType))
	property.SetValue(int8(valueType), value)
	propertyTitle, err := structs.AddPropertyTitle(title)
	utils.CheckError(err)
	property.SetTitle(propertyTitle)

	if relationship.GetProperty() == nil {
		relationship.SetProperty(property)
	} else {
		lastProperty := relationship.GetProperty()
		for true {
			if lastProperty.GetNextProperty() == nil {
				lastProperty.SetNextProperty(property)
				break
			} else {
				lastProperty = lastProperty.GetNextProperty()
			}
		}
	}

	return property
}

func GetNode(id int) (node *structs.Node) {
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