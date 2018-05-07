package api

import (
	"graph-db/internal/app/core"
	"graph-db/internal/app/core/globals"
	"testing"
)

var fh core.FileHandler

func init() {
	globals.FileHandler = fh
}

func TestCreateNode(t *testing.T) {
	core.InitDb("test_db", "local")
	node := CreateNode("test_node")
	if node.GetId() != 0 {
		t.Errorf("Node id mismatch")
	}
	if node.GetProperty() != nil {
		t.Errorf("Node is not nil")
	}
	if len(node.GetLabel().GetLabelNames()) != 5 {
		t.Errorf("Labels length is not 5")
	}
	if node.GetLabel().GetLabelNames()[0].GetTitle() != "test_node" {
		t.Errorf("Label title mismatch")
	}
	DropDatabase("test_db")
}

func TestCreateRelationship(t *testing.T) {
	core.InitDb("test_db", "local")
	node1 := CreateNode("node1")
	node2 := CreateNode("node2")
	if node1.GetId() != 0 {
		t.Errorf("Node1 id mismatch")
	}
	if node2.GetId() != 1 {
		t.Errorf("Node2 id mismatch")
	}
	relationship1 := CreateRelationship(&node1, &node2, "tests")
	if relationship1.GetId() != 0 {
		t.Errorf("Relationship 1 id mismatch")
	}
	if relationship1.GetTitle().GetTitle() != "tests" {
		t.Errorf("Relationship 1 title mismatch")
	}
	if relationship1.GetFirstNode() != &node1 {
		t.Errorf("First node mismatch")
	}
	if relationship1.GetSecondNode() != &node2 {
		t.Errorf("Second node mismatch")
	}
	node3 := CreateNode("node3")
	//node4 := CreateNode("node4")
	relationship2 := CreateRelationship(&node1, &node3, "stores")
	if relationship2.GetId() != 1 {
		t.Errorf("Relationship 2 id mismatch")
	}
	if relationship2.GetTitle().GetTitle() != "stores" {
		t.Errorf("Relationship 2 title mismatch")
	}
	if relationship2.GetFirstNode() != &node1 {
		t.Errorf("First node mismatch")
	}
	if relationship2.GetSecondNode() != &node3 {
		t.Errorf("Second node mismatch")
	}
	if relationship1.GetFirstNextRelationship().GetId() != relationship2.GetId() {
		t.Errorf("First node previous relationship mismatch")
	}
	if relationship2.GetFirstPreviousRelationship().GetId() != relationship1.GetId() {
		t.Errorf("First node previous relationship mismatch")
	}
	DropDatabase("test_db")
}