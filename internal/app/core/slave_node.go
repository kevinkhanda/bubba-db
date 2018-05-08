package core

import (
	"graph-db/internal/app/core/globals"
	"os"
	"log"
)

func getFilePointerByName(filePath string) *os.File {
	pwd, _ := os.Getwd()
	filePath = pwd + filePath
	file, err := os.OpenFile(filePath, os.O_RDWR, 0)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func (entity *Entity) Ping(request *RPCRequest, reply *Reply) error {
	reply.Message = "Pong"
	return err
}

func (entity *Entity) SendStatus(request *RPCRequest, reply *Reply) error {
	reply.Message = "success"
	return err
}

func (entity *Entity) Deploy(request *RPCRequest, reply *Reply) error  {
	var fileHandler FileHandler
	fileHandler.InitFileSystem()
	globals.FileHandler = fileHandler
	reply.Message = "success"
	return nil
}

func (entity *Entity) InitDatabaseStructure(request *RPCRequest, reply *Reply) error {
	var fileHandler  FileHandler
	fileHandler.InitDatabaseStructure(request.Data.Payload)
	globals.Config.WriteAt([]byte("[\"10.240.22.31:7000\"]"), 0)
	reply.Message = "success"
	return nil
}

func (entity *Entity) SwitchDatabaseStructure(request *RPCRequest, reply *Reply) error  {
	var fh FileHandler
	err = fh.SwitchDatabaseStructure(request.Data.Payload)
	if err == nil {
		reply.Message = "success"
	}
	return nil
}

func (entity *Entity) DropDatabase(request *RPCRequest, reply *Reply) error  {
	var fileHandler FileHandler
	err = fileHandler.DropDatabase(string(request.Data.Payload))
	reply.Message = "success"
	return err
}

func (entity *Entity) Read(request *RPCRequest, reply *Reply) error  {
	var fh FileHandler
	file := getFilePointerByName(request.Data.File)
	fh.Read(file, request.Data.Offset, &request.Data.Bs, request.Data.Id)
	reply.Message = "success"
	reply.Data = request.Data.Bs
	return nil
}

func (entity *Entity) Write(request *RPCRequest, reply *Reply) error  {
	var fh FileHandler
	file := getFilePointerByName(request.Data.File)
	err = fh.Write(file, request.Data.Offset, request.Data.Bs, request.Data.Id)
	if err == nil {
		reply.Message = "success"
	}
	return err
}

func (entity *Entity) FreeId(request *RPCRequest, reply *Reply) error {
	return nil
}