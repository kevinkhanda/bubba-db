package core

import "net/rpc"

type Slave struct {
	ip			string
	port		string
	identifier	string
	connector	*rpc.Client
}


func (entity *Entity) Ping(request *RPCRequest, reply *string) error {
	*reply = "Pong"
	return err
}

