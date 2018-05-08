package core

import (
	"os"
	"graph-db/internal/app/core/globals"
)

type DistributedFileHandler struct {
}

var exceptionFileNames = [3]string{"LabelsTitlesStore", "RelationshipsTitlesStore", "PropertiesTitlesStore"}

func inArray(fileName string) bool {
	res := false
	for _, exceptionFileName := range exceptionFileNames {
		if exceptionFileName == fileName {
			res = true
			break
		}
	}
	return res
}

func (dfh DistributedFileHandler) InitFileSystem() {
	if _, err := os.Stat(rootPath); os.IsNotExist(err) {
		os.Mkdir(rootPath, os.ModePerm)
	}
}

func (dfh DistributedFileHandler) InitDatabaseStructure(dbIdentifier string) {
	var fh FileHandler
	fh.InitDatabaseStructure(dbIdentifier)
	globals.Config.WriteAt([]byte("[\"10.240.22.31:7000\"]"), 0)
	println(len(master.Slaves))
	for i := range master.Slaves {
		SendInitDatabaseStructure(&master.Slaves[i], &dbIdentifier)
	}
}

func (dfh DistributedFileHandler) SwitchDatabaseStructure(dbTitle string) (err error) {
	for i := range master.Slaves {
		SendSwitchDatabaseStructure(&master.Slaves[i], &dbTitle)
	}
	return nil
}

func (dfh DistributedFileHandler) DropDatabase(dbIdentifier string) (err error) {
	for i := range master.Slaves {
		SendDropDatabase(&master.Slaves[i], &dbIdentifier)
	}
	return nil
}

func (dfh DistributedFileHandler) Read(file *os.File, offset int, bs []byte, id int) (err error) {
	if inArray(file.Name()) {
		var fh FileHandler
		fh.Read(file, offset, bs, id)
	} else {
		slaveIndex := id % len(master.Slaves)
		println(len(bs))
		bs, err = SendReadData(&master.Slaves[slaveIndex], file, offset, id, bs)
	}
	return nil
}

func (dfh DistributedFileHandler) Write(file *os.File, offset int, bs []byte, id int) (err error) {
	if inArray(file.Name()) {
		var fh FileHandler
		fh.Write(file, offset, bs, id)
	} else {
		slaveIndex := id % len(master.Slaves)
		SendWriteData(&master.Slaves[slaveIndex], file, offset, id, bs)
	}
	return nil
}

func (dfh DistributedFileHandler) ReadId(file *os.File) (id int, err error) {
	var fh FileHandler
	id, err = fh.ReadId(file)
	return id, err
}

func (dfh DistributedFileHandler) FreeId(file *os.File, id int) (err error) {
	var fh FileHandler
	err = fh.FreeId(file, id)
	return err
}
