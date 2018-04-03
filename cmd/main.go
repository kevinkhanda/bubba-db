package main

import (
	"github.com/kevinkhanda/graph-db/cmd/core"
)

func main() {
	core.InitFileSystem()
	core.InitDatabaseStructure("test-db")
}
