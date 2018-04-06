package main

import (
	"graph-db/internal/app/core"
	"graph-db/internal/pkg/utils"
)

func main() {
	err := core.InitDatabase("asd", "local")
	utils.CheckError(err)
}
