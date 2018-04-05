package core

func InitDatabase(dbTitle string) {
	initFileSystem()
	initDatabaseStructure(dbTitle)
}
