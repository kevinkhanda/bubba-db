package core

func (entity *Entity) Ping(request *RPCRequest, reply *string) error {
	*reply = "Pong"
	return err
}

func (entity *Entity) SendStatus(request *RPCRequest, reply *string) error {
	*reply = "success"
	return err
}

func (entity *Entity) Deploy(request *RPCRequest, reply *string) error  {
	var fileHandler = new (FileHandler)
	fileHandler.InitFileSystem()
	fileHandler.InitDatabaseStructure(string(request.Data))
	*reply = "success"
	return nil
}

func (entity *Entity) SwitchDatabaseStructure(request *RPCRequest, reply *string) error  {
	return nil
}

func (entity *Entity) DropDatabase(request *RPCRequest, reply *string) error  {
	var fileHandler = new(FileHandler)
	fileHandler.DropDatabase(string(request.Data))
	*reply = "success"
	return nil
}

func (entity *Entity) Read(request *RPCRequest, reply *string) error  {
	return nil
}

func (entity *Entity) Write(request *RPCRequest, reply *string) error  {
	return nil
}

func (entity *Entity) FreeId(request *RPCRequest, reply *string) error {
	return nil
}