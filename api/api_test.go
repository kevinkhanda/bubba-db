package api

import (
	"graph-db/internal/app/core"
	"graph-db/internal/app/core/globals"
	"testing"
	"graph-db/internal/app/core/structs"
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
	relationship1 := CreateRelationship(node1, node2, "tests")
	if relationship1.GetId() != 0 {
		t.Errorf("Relationship 1 id mismatch")
	}
	if relationship1.GetTitle().GetTitle() != "tests" {
		t.Errorf("Relationship 1 title mismatch")
	}
	if relationship1.GetFirstNode() != node1 {
		t.Errorf("First node mismatch")
	}
	if relationship1.GetSecondNode() != node2 {
		t.Errorf("Second node mismatch")
	}
	node3 := CreateNode("node3")
	//node4 := CreateNode("node4")
	relationship2 := CreateRelationship(node1, node3, "stores")
	if relationship2.GetId() != 1 {
		t.Errorf("Relationship 2 id mismatch")
	}
	if relationship2.GetTitle().GetTitle() != "stores" {
		t.Errorf("Relationship 2 title mismatch")
	}
	if relationship2.GetFirstNode() != node1 {
		t.Errorf("First node mismatch")
	}
	if relationship2.GetSecondNode() != node3 {
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

func TestCreatePropertyForNode(t *testing.T) {
	core.InitDb("test_db", "local")
	node1 := CreateNode("node1")
	property1 := CreatePropertyForNode(node1, "color", 2, "red")
	if property1.GetId() != 0 {
		t.Errorf("Property id mismatch")
	}
	if property1.GetNextProperty() != nil {
		t.Errorf("Next property is not null")
	}
	if property1.GetValueType() != 2 {
		t.Errorf("Value type mismatch")
	}
	if structs.GetStringValue(*property1.GetValue()) != "red" {
		t.Errorf("String value mismatch")
	}
	property2 := CreatePropertyForNode(node1, "amount", 0, 780)
	if property2.GetId() != 1 {
		t.Errorf("Property id mismatch")
	}
	if property1.GetNextProperty() != property2 {
		t.Errorf("Next property is not first")
	}
	if property2.GetValueType() != 0 {
		t.Errorf("Value type mismatch")
	}
	if structs.GetIntegerValue(*property2.GetValue()) != 780 {
		t.Errorf("Integer value mismatch")
	}
	property3 := CreatePropertyForNode(node1, "price", 1, 70.5)
	if property3.GetId() != 2 {
		t.Errorf("Property id mismatch")
	}
	if property2.GetNextProperty() != property3 {
		t.Errorf("Next property is not first")
	}
	if property3.GetValueType() != 1 {
		t.Errorf("Value type mismatch")
	}
	if structs.GetDoubleValue(*property3.GetValue()) != 70.5 {
		t.Errorf("Double value mismatch")
	}

	if property1.GetId() != 0 {
		t.Errorf("Property id mismatch")
	}
	DropDatabase("test_db")
}

func TestCreatePropertyForRelationship(t *testing.T) {
	core.InitDb("test_db", "local")
	node1 := CreateNode("node1")
	node2 := CreateNode("node2")
	relationship := CreateRelationship(node1, node2, "stores")
	property1 := CreatePropertyForRelationship(relationship, "description", 2, "Property, which " +
		"provides an information to fit more that 31 symbols in order to test the recursive data allocation")
	property2 := CreatePropertyForNode(node1, "name", 2, "Node number one")
	if property1.GetId() != 0 {
		t.Errorf("Property id mismatch")
	}
	if property1.GetNextProperty() != nil {
		t.Errorf("Next property is not null")
	}
	if property1.GetValueType() != 2 {
		t.Errorf("Value type mismatch")
	}
	if structs.GetStringValue(*property1.GetValue()) != "Property, which " +
		"provides an information to fit more that 31 symbols in order to test the recursive data allocation" {
		t.Errorf("String value mismatch")
	}

	if property2.GetId() != 1 {
		t.Errorf("Property id mismatch")
	}
	if property2.GetNextProperty() != nil {
		t.Errorf("Next property is not null")
	}
	if property2.GetValueType() != 2 {
		t.Errorf("Value type mismatch")
	}
	if structs.GetStringValue(*property2.GetValue()) != "Node number one" {
		t.Errorf("String value mismatch")
	}

	DropDatabase("test_db")
}