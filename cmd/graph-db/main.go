package main

import (
	"graph-db/internal/app/core"
	"graph-db/internal/pkg/utils"
)

func main() {
	err := core.InitDb("asd", "local")
	utils.CheckError(err)
}
