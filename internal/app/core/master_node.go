package core

import (
	"log"
	"errors"
	"os"
)

var master Entity

func SendReadData(entity *Entity, file *os.File, offset int, id int) ([]byte, error)  {
	var reply  string
	var attempts = 0
	requestedData := RequestedData{
		File: file,
		Offset: offset,
		Id: id,
	}
	for attempts < 5 {
		err = nil
		request := RPCRequest {requestedData }
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
	return nil, nil // TODO
}

func SendWriteData(entity *Entity, file *os.File, offset int, id int, bs []byte) error {
	reply := new(Reply)
	var attempts = 0
	requestedData := RequestedData{
		File: file,
		Offset: offset,
		Id: id,
		Bs: bs,
	}
	for attempts < 5 {
		err = nil
		request := RPCRequest{ requestedData }
		err = entity.connector.Call("Entity.Write", &request, &reply)
		if err != nil {
			log.Fatal("Problems in requestSlaveStatus ", err)
			err = errors.New("problems in requestSlaveStatus")
			attempts++
			continue
		}
		if reply.Message == "success" {
			attempts = 5
		}
	}
	return err
}

func SendSwitchDatabaseStructure(entity *Entity, newStructure *string) error {
	reply := new(Reply)
	var attempts = 0
	for attempts < 5 {
		err = nil
		request := RPCRequest{ RequestedData{ Payload: *newStructure } }
		err = entity.connector.Call("Entity.SwitchDatabaseStructure", &request, &reply)
		if err != nil {
			log.Fatal("Problems in requestSlaveStatus ", err)
			err = errors.New("problems in requestSlaveStatus")
			attempts++
			continue
		}
		if reply.Message == "success" {
			println("Slave " + entity.ip + ":" + entity.port + " switched status on " + *newStructure)
			attempts = 5
		}
	}
	return nil
}

func RequestSlaveStatus(entity *Entity) error {
	reply := new(Reply)
	var attempts = 0
	for attempts < 5 {
		err = nil
		request := RPCRequest{ *new(RequestedData) }
		err = entity.connector.Call("Entity.SendStatus", &request, &reply)
		if err != nil {
			log.Fatal("Problems in requestSlaveStatus ", err)
			err = errors.New("problems in requestSlaveStatus")
			attempts++
			continue
		}
		if reply.Message == "success" {
			println("Slave " + entity.ip + ":" + entity.port + " is available")
			attempts = 5
		}
	}
	return err
}

func SendDeploy(entity *Entity) error {
	reply := new(Reply)
	var attempts = 0
	for attempts < 5 {
		err = nil
		request := RPCRequest{*new(RequestedData) }
		err = entity.connector.Call("Entity.Deploy", &request, &reply)
		if err != nil {
			log.Fatal("Problems in requestSlaveStatus ", err)
			err = errors.New("problems in requestSlaveStatus")
			attempts++
			continue
		}
		if reply.Message == "success" {
			attempts = 5
		}
	}
	return err
}

func SendInitDatabaseStructure(entity *Entity, dbName *string) error {
	reply := new(Reply)
	var attempts = 0
	for attempts < 5 {
		err = nil
		request := RPCRequest{RequestedData{ Payload: *dbName } }
		err = entity.connector.Call("Entity.InitDatabaseStructure", &request, &reply)
		if err != nil {
			log.Fatal("Problems in requestSlaveStatus ", err)
			err = errors.New("problems in requestSlaveStatus")
			attempts++
			continue
		}
		if reply.Message == "success" {
			attempts = 5
		}
	}
	return err
}

func SendDropDatabase(entity *Entity, dbName *string) error {
	reply := new(Reply)
	var attempts = 0
	for attempts < 5 {
		err = nil
		request := RPCRequest{ RequestedData{ Payload: *dbName } }
		err = entity.connector.Call("Entity.DropDatabase", &request, &reply)
		if err != nil {
			log.Fatal("Problems in requestSlaveStatus ", err)
			err = errors.New("problems in requestSlaveStatus")
			attempts++
			continue
		}
		if reply.Message == "success" {
			attempts = 5
		}
	}
	return err
}
