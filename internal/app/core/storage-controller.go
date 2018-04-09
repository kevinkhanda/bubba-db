package core

import "graph-db/internal/app/core/structs"

type Storage interface {

}

func CreateNode() (node structs.Node) {
	node.Create()
	return node
}

func GetNode(id int) (node structs.Node) {
	return node.Get(id)
}
