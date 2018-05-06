package core

import "net/rpc"

type Slavef struct {
	ip			string
	port		string
	identifier	string
	connector	*rpc.Client
}


type RPCCall struct {
	Ip		string
	Port	string
	Method 	string
	Data 	[]byte
}

func (entity *Entity) Ping(request *RPCCall, reply string) error {
	reply = "Pong"
	return nil
}

