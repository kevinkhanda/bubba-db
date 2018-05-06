package main

import (
	"graph-db/internal/app/core"
	"graph-db/internal/pkg/utils"
)

func main() {
	//err := core.InitDb("asd", "local")
	err := core.SwitchDb("asd")
	utils.CheckError(err)

	//arith := new(core.Entity)
	//rpc.Register(arith)
	//rpc.HandleHTTP()
	//l, e := net.Listen("tcp", ":7000")
	//if e != nil {
	//	log.Fatal("listen error:", e)
	//}
	//go http.Serve(l, nil)

	core.Test()
}
