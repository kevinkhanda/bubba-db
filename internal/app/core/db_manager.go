package core

import (
	"errors"
	"graph-db/internal/app/core/globals"
	"graph-db/internal/pkg/utils"
)

func InitDb(dbTitle string, storageMode string) (err error) {
	if storageMode == "local" {
		var fh FileHandler
		fh.InitFileSystem()
		fh.InitDatabaseStructure(dbTitle)
		globals.FileHandler = fh
		globals.CurrentDb = dbTitle
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
	utils.CheckError(err)
	globals.FileHandler = fh
	globals.CurrentDb = dbTitle
	return err
}

func DropDb(dbTitle string) (err error) {
	err = globals.FileHandler.DropDatabase(dbTitle)
	return err
}
