package core

import (
	"errors"
	"graph-db/internal/app/core/globals"
)

func InitDatabase(dbTitle string, storageMode string) error {
	if storageMode == "local" {
		var fh FileHandler
		fh.InitFileSystem()
		fh.InitDatabaseStructure(dbTitle)
		globals.FileHandler = fh
		return nil
	} else if storageMode == "distributed" {
		return errors.New("not implemented yet")
	} else {
		return errors.New("storageMode should be local or distributed")
	}
}
