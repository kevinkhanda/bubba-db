package core

func initDatabase(dbTitle string) {
	initFileSystem()
	initDatabaseStructure(dbTitle)
}
