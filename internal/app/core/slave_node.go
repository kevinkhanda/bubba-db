package core

func (entity *Entity) Ping(request *RPCRequest, reply *Reply) error {
	reply.Message = "Pong"
	return err
}

func (entity *Entity) SendStatus(request *RPCRequest, reply *Reply) error {
	reply.Message = "success"
	return err
}

func (entity *Entity) Deploy(request *RPCRequest, reply *Reply) error  {
	var fileHandler = new (FileHandler)
	fileHandler.InitFileSystem()
	reply.Message = "success"
	return nil
}

func (entity *Entity) InitDatabaseStructure(request *RPCRequest, reply *Reply) error {
	var fileHandler  = new(FileHandler)
	fileHandler.InitDatabaseStructure(string(request.Data.Payload))
	return nil
}

func (entity *Entity) SwitchDatabaseStructure(request *RPCRequest, reply *Reply) error  {
	fh := new(FileHandler)
	err = fh.SwitchDatabaseStructure(request.Data.Payload)
	if err == nil {
		reply.Message = "success"
	}
	return nil
}

func (entity *Entity) DropDatabase(request *RPCRequest, reply *Reply) error  {
	var fileHandler = new(FileHandler)
	err = fileHandler.DropDatabase(string(request.Data.Payload))
	reply.Message = "success"
	return err
}

func (entity *Entity) Read(request *RPCRequest, reply *Reply) error  {
	fh := new(FileHandler)
	fh.Read(request.Data.File, request.Data.Offset, reply.Data, request.Data.Id)
	reply.Message = "success"
	return nil
}

func (entity *Entity) Write(request *RPCRequest, reply *Reply) error  {
	fh := new(FileHandler)
	err = fh.Write(request.Data.File, request.Data.Offset, request.Data.Bs, request.Data.Id)
	if err == nil {
		reply.Message = "success"
	}
	return err
}

func (entity *Entity) FreeId(request *RPCRequest, reply *Reply) error {
	return nil
}