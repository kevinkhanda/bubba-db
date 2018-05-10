package main

import (
	"graph-db/api"
	"graph-db/internal/app/core/structs"
	"strings"
	"strconv"
)

func printNode(node structs.Node) {
	var str string
	str += strings.Join([]string{"Node id: ", strconv.Itoa(node.GetId()), ", labels: "}, "")
	for index, label := range node.GetLabel().GetLabelNames() {
		if label != nil && label.GetTitle() != "" {
			str += label.GetTitle()
			if index != len(node.GetLabel().GetLabelNames()) - 1 {
				if node.GetLabel().GetLabelNames()[index+1] != nil {
					str += ", "
				} else {
					str += "."
				}
			}
		}
	}
	if node.GetProperty() != nil {
		str += "\nNode properties: \n"
		property := node.GetProperty()
		var prop string
		index := 1
		for property != nil {
			prop += strings.Join([]string{strconv.Itoa(index), ". Property title: \"", property.GetTitle().String(),
				"\"; value: ", (*property.GetValue()).String(), "\n"}, "")
			property = property.GetNextProperty()
			index++
		}
		str += prop[:len(prop) - 2]
	}
	if node.GetRelationship() != nil {
		str += "\nRelationships: \n"
		relationship := node.GetRelationship()
		var rel string
		index := 0
		for relationship != nil {
			rel += strings.Join([]string{"\t", strconv.Itoa(index + 1), ". Node with id ",
				strconv.Itoa(relationship.GetFirstNode().GetId()), " ",
				relationship.GetTitle().GetTitle(), " node with id ",
				strconv.Itoa(relationship.GetSecondNode().GetId()), "\n"}, "")
			property := relationship.GetProperty()
			index2 := 0
			if property != nil {
				rel += "\t\tRelationship properties:\n"
			}
			for property != nil {
				rel += strings.Join([]string{"\t\t\t", strconv.Itoa(index2 + 1), ". Property title: \"", property.GetTitle().String(),
					"\"; value: ", (*property.GetValue()).String(), "\n"}, "")
				property = property.GetNextProperty()
				index2++
			}
			if relationship.GetFirstNode().GetId() == node.GetId() {
				relationship = relationship.GetFirstNextRelationship()
			} else {
				relationship = relationship.GetSecondNextRelationship()
			}
			index++
		}
		str += rel
	}
	println(str)
}

func main() {
	api.CreateDatabase("asd", "local") // flag "distributed" for distributed

	node1 := api.CreateNode("Kevin")
	node2 := api.CreateNode("Sergey")
	node3 := api.CreateNode("Oleg")
	node4 := api.CreateNode("Car")

	_ = api.CreateRelationship(node1, node2, "is friends with")
	_ = api.CreateRelationship(node2, node3, "is friends with")
	_ = api.CreateRelationship(node1, node3, "is friends with")
	relationship4 := api.CreateRelationship(node1, node4, "owning a")

	_ = api.CreatePropertyForNode(node1, "age", 0, 21)
	_ = api.CreatePropertyForNode(node1, "city", 2, "Moscow")
	_ = api.CreatePropertyForNode(node1, "height", 1, 185.3)
	_ = api.CreatePropertyForNode(node2, "age", 0, 20)
	_ = api.CreatePropertyForNode(node2, "city", 2, "Tolyatti")
	_ = api.CreatePropertyForNode(node2, "height", 1, 179.4)
	_ = api.CreatePropertyForNode(node3, "age", 0, 20)
	_ = api.CreatePropertyForNode(node3, "city", 2, "Tolyatti")
	_ = api.CreatePropertyForNode(node3, "height", 1, 179.4)
	_ = api.CreatePropertyForNode(node4, "brand", 2, "BMW")
	_ = api.CreatePropertyForNode(node4, "year", 0, 2014)

	_ = api.CreatePropertyForRelationship(relationship4, "how many years", 0, 2)

	printNode(*node1)  // use print node method to retrieve information
	printNode(*node2)
	printNode(*node3)
	printNode(*node4)
}