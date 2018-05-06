package core

import (
	"errors"
	"graph-db/internal/app/core/globals"
	"graph-db/internal/pkg/utils"
	"os"
)

func InitDb(dbTitle string, storageMode string) (err error) {
	if storageMode == "local" {
		var fh FileHandler
		fh.InitFileSystem()
		fh.InitDatabaseStructure(dbTitle)
		globals.LabelTitleMap = make(map[string]globals.MapValue)
		globals.RelationshipTitleMap = make(map[string]globals.MapValue)
		globals.PropertyTitleMap = make(map[string]globals.MapValue)
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
	globals.LabelTitleMap = make(map[string]globals.MapValue)
	globals.RelationshipTitleMap = make(map[string]globals.MapValue)
	globals.PropertyTitleMap = make(map[string]globals.MapValue)
	fillMap(globals.PropertyTitleMap, globals.PropertiesTitlesStore, globals.PropertiesTitlesSize)
	fillMap(globals.RelationshipTitleMap, globals.RelationshipsTitlesStore, globals.RelationshipsTitlesSize)
	fillMap(globals.LabelTitleMap, globals.LabelsTitlesStore, globals.LabelsTitlesSize)
	globals.CurrentDb = dbTitle
	return err
}

func DropDb(dbTitle string) (err error) {
	err = globals.FileHandler.DropDatabase(dbTitle)
	return err
}

func fillMap(m map[string]globals.MapValue, file *os.File, recordSize int) {
	var (
		i int
		counter int32
		str string
		err, conversionError error
		bs []byte
	)
	bs = make([]byte, recordSize)
	i = 0
	for true {
		err = globals.FileHandler.Read(file, i * recordSize, bs)
		if err != nil {
			break
		}
		counter, conversionError = utils.ByteArrayToInt32(bs[recordSize - 4:])
		utils.CheckError(conversionError)
		if counter != 0 {
			str = utils.RemoveStopCharacter(utils.ByteArrayToString(bs[0 : recordSize-4]))
			utils.CheckError(conversionError)
			m[str] = globals.MapValue{Id: i, Counter: int(counter)}
		}
		i++
	}

}

