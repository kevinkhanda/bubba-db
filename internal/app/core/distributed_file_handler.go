package core

type DistributedFileHandler struct {
}

func (dfh DistributedFileHandler) DropDatabase(dbIdentifier string)  {
	for i, slave := range master.slaves{
		if slave.identifier == dbIdentifier {
			SendDropDatabase(&master.slaves[i])
		}
	}
}