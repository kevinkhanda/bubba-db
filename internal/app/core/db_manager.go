package core

import (
	"errors"
	"graph-db/internal/app/core/globals"
)

func InitDb(dbTitle string, storageMode string) (err error) {
	if storageMode == "local" {
		var fh FileHandler
		fh.InitFileSystem()
		fh.InitDatabaseStructure(dbTitle)
		globals.FileHandler = fh
		return err
	} else if storageMode == "distributed" {
		return errors.New("not implemented yet")
	} else {
		return errors.New("storageMode should be local or distributed")
	}
}

func SwitchDb(dbTitle string) (err error){
	var fh FileHandler
	err = fh.SwitchDatabaseStructure(dbTitle)
	globals.FileHandler = fh
	return err
}

func DropDb(dbTitle string) (err error) {
	err = globals.FileHandler.DropDatabase(dbTitle)
	return err
}
