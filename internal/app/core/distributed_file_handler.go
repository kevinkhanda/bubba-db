package core

import "os"

type DistributedFileHandler struct {
}

var exceptionFileNames = [6]string{"","","","","",""}

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

func (dfh DistributedFileHandler) DropDatabase(dbIdentifier string)  {
	for i, slave := range master.slaves{
		if slave.identifier == dbIdentifier {
			SendDropDatabase(&master.slaves[i])
		}
	}
}

func (dfh DistributedFileHandler) Read(file *os.File, offset int, bs []byte) (err error) {
	if inArray(file.Name()) {

	} else {

	}
	//TODO: implement
	return nil
}

func (dfh DistributedFileHandler) Write(file *os.File, offset int, bs []byte) (err error) {
	if inArray(file.Name()) {

	} else {

	}
	//TODO: implement
	return nil
}

func (dfh DistributedFileHandler) ReadId(file *os.File) (id int, err error) {
	if inArray(file.Name()) {

	} else {

	}
	//TODO: implement
	return 0, err
}

func (dfh DistributedFileHandler) FreeId(file *os.File, id int) (err error) {
	if inArray(file.Name()) {

	} else {

	}
	//TODO: implement
	return nil
}