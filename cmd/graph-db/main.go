package main

import (
	"graph-db/internal/app/core"
	"log"
	"graph-db/internal/pkg/utils"
	"graph-db/internal/app/core/globals"
)

func main() {
	//err := core.InitDb("asd", "local")
	//err := core.SwitchDb("asd")
	//utils.CheckError(err)
	dbTitle := "asd"
	err := core.InitDb(dbTitle, "distributed")
	if err != nil {
		log.Fatal("Error in initialization of database")
	}
	dfh := new(core.DistributedFileHandler)
	dfh.InitDatabaseStructure(dbTitle)
	bs := utils.StringToByteArray("Test")
	dfh.Write(globals.NodesStore, 20, bs, 0)
	newBs := make([]byte, 4)
	dfh.Read(globals.NodesStore, 20, newBs, 0)
	if string(newBs) != string(bs) {
		log.Fatal("Byte arrays are not the same!")
	}
}
