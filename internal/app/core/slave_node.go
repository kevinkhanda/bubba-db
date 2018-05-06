package core

func (entity *Entity) Ping(request *RPCRequest, reply *string) error {
	*reply = "Pong"
	return err
}

func (entity *Entity) SendStatus(request *RPCRequest, reply *string) error {
	*reply = "success"
	return err
}

