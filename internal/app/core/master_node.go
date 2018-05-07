package core

import (
	"log"
	"errors"
)

var master Entity

func SendReadData(entity *Entity) error  {
	var reply  string
	var attempts = 0
	for attempts < 5 {
		err = nil
		request := RPCRequest{nil}
		err = entity.connector.Call("Entity.Read", &request, &reply)
		if err != nil {
			log.Fatal("Problems in requestSlaveStatus ", err)
			err = errors.New("problems in requestSlaveStatus")
			attempts++
			continue
		}
		if reply == "success" {
			attempts = 5
		}
	}
	return err
}

func RequestSlaveStatus(entity *Entity) error {
	var reply string
	var attempts = 0
	for attempts < 5 {
		err = nil
		request := RPCRequest{nil}
		err = entity.connector.Call("Entity.SendStatus", &request, &reply)
		if err != nil {
			log.Fatal("Problems in requestSlaveStatus ", err)
			err = errors.New("problems in requestSlaveStatus")
			attempts++
			continue
		}
		if reply == "success" {
			println("Slave " + entity.ip + ":" + entity.port + " is available")
			attempts = 5
		}
	}
	return err
}

func SendDeploy(entity *Entity) error {
	var reply string
	var attempts = 0
	for attempts < 5 {
		err = nil
		request := RPCRequest{[]byte(string(entity.identifier))}
		err = entity.connector.Call("Entity.Deploy", &request, &reply)
		if err != nil {
			log.Fatal("Problems in requestSlaveStatus ", err)
			err = errors.New("problems in requestSlaveStatus")
			attempts++
			continue
		}
		if reply == "success" {
			attempts = 5
		}
	}
	return err
}

func SendInitDatabaseStructure(entity *Entity, dbName string) error {
	var reply string
	var attempts = 0
	for attempts < 5 {
		err = nil
		request := RPCRequest{[]byte(dbName)}
		err = entity.connector.Call("Entity.InitDatabaseStructure", &request, &reply)
		if err != nil {
			log.Fatal("Problems in requestSlaveStatus ", err)
			err = errors.New("problems in requestSlaveStatus")
			attempts++
			continue
		}
		if reply == "success" {
			attempts = 5
		}
	}
	return err
}

func SendWriteData(entity *Entity) error {
	var reply  string
	var attempts = 0
	for attempts < 5 {
		err = nil
		request := RPCRequest{nil}
		err = entity.connector.Call("Entity.Write", &request, &reply)
		if err != nil {
			log.Fatal("Problems in requestSlaveStatus ", err)
			err = errors.New("problems in requestSlaveStatus")
			attempts++
			continue
		}
		if reply == "success" {
			attempts = 5
		}
	}
	return err
}

func SendDropDatabase(entity *Entity) error {
	var reply string
	var attempts = 0
	for attempts < 5 {
		err = nil
		request := RPCRequest{[]byte(string(entity.identifier))}
		err = entity.connector.Call("Entity.DropDatabase", &request, &reply)
		if err != nil {
			log.Fatal("Problems in requestSlaveStatus ", err)
			err = errors.New("problems in requestSlaveStatus")
			attempts++
			continue
		}
		if reply == "success" {
			attempts = 5
		}
	}
	return err
}
