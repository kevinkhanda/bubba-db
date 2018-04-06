package core

import "graph-db/internal/app/core/structs"

type Storage interface {

}

func createNode() structs.Node {
	var node structs.Node
	node.Create()
	return node
}