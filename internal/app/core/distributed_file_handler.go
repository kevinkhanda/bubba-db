package core

import "os"

type DistributedFileHandler struct {
}

var exceptionFileNames = [6]string{"","","","","",""} // TODO: fill files names

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

func (dfh DistributedFileHandler) InitDatabaseStructure(dbIdentifier string)  {
	for i := range master.slaves {
		SendInitDatabaseStructure(&master.slaves[i], &dbIdentifier)
	}
}

func (dfh DistributedFileHandler) SwitchDatabaseStructure(dbTitle string) (err error) {
	for i := range master.slaves {
		SendSwitchDatabaseStructure(&master.slaves[i], &dbTitle)
	}
	return nil
}

func (dfh DistributedFileHandler) DropDatabase(dbIdentifier string)  {
	for i := range master.slaves {
		SendDropDatabase(&master.slaves[i], &dbIdentifier)
	}
}

func (dfh DistributedFileHandler) Read(file *os.File, offset int, bs []byte, id int) (err error) {
	if inArray(file.Name()) {
		// kek
	} else {
		slaveIndex := id % len(master.slaves) + 1
		bs, err = SendReadData(&master.slaves[slaveIndex], file, offset, id)
	}
	return nil
}

func (dfh DistributedFileHandler) Write(file *os.File, offset int, bs []byte) (err error) {
	if inArray(file.Name()) {
		// kek
	} else {
		//SendWriteData()
	}
	//TODO: implement
	return nil
}

func (dfh DistributedFileHandler) ReadId(file *os.File) (id int, err error) {
	var fh = new(FileHandler)
	id, err = fh.ReadId(file)
	return id, err
}

func (dfh DistributedFileHandler) FreeId(file *os.File, id int) (err error) {
	var fh = new(FileHandler)
	err = fh.FreeId(file, id)
	return err
}