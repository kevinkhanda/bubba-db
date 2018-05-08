package main

import (
	"graph-db/internal/app/core"
	"graph-db/internal/pkg/utils"
	"graph-db/internal/app/core/globals"
	"log"
	"graph-db/api"
	"graph-db/internal/app/core/structs"
)

func printNode(node structs.Node) {
	var str string
	str += "Node labels: "
	for _, label := range node.GetLabel().GetLabelNames() {
		if label.GetTitle() != "" {
			str += label.GetTitle()
			str += ", "
		}
	}
	str += "\n"
	str += "Relationships: \n"
	relationship := node.GetRelationship()
	i := 1
	for true {
		if relationship != nil {
			str += "\n\t"
			str += string(i)
			str += "\t Title: "
			str += relationship.GetTitle().GetTitle()
			i++
		}
	}
}

func main() {
	//err := core.InitDb("asd", "local")
	//err = core.SwitchDb("asd")
	//utils.CheckError(err)

	dbTitle := "asd"
	var dfh core.DistributedFileHandler
	dfh.InitFileSystem()
	err := core.InitDb(dbTitle, "distributed")
	dfh.InitDatabaseStructure(dbTitle)
	if err != nil {
		log.Fatal("Error in initialization of database")
	}

	bs := utils.StringToByteArray("Test")
	dfh.Write(globals.NodesStore, 20, bs, 0)
	newBs := make([]byte, 4)
	dfh.Read(globals.NodesStore, 20, &newBs, 0)

	//if string(newBs) != string(bs) {
	//	log.Fatal("Byte arrays are not the same!")
	//} else {
	//	println("Congratulations!")
	//}

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

	printNode(*node1)
	printNode(*node2)
	printNode(*node3)
	printNode(*node4)
}