package globals

import "os"

type FileHandlerInterface interface {
	InitFileSystem()
	InitDatabaseStructure(dbTitle string)
	SwitchDatabaseStructure(dbTitle string) (err error)
	DropDatabase(dbTitle string) (err error)
	Write(file *os.File, offset int, bs []byte) (err error)
	Read(file *os.File, offset int, bs []byte) (err error)
	ReadId(file *os.File) (id int, err error)
	FreeId(file *os.File, id int) (err error)
}

const (
	INTEGER = 0
	DOUBLE = 1
	STRING = 2
)

const (
	LabelsTitlesSize = 36
	RelationshipsTitlesSize = 36
	PropertiesTitlesSize = 36
	LabelsSize = 34
	NodesSize = 13
	RelationshipsSize = 34
	PropertiesSize = 14
	StringSize = 36
	DoubleSize = 9
)

var (
	CurrentDb string
	// nodes/id
	NodesId, LabelsId, LabelsTitlesId,
	// nodes/store
	NodesStore, LabelsStore, LabelsTitlesStore,
	// relationships/id
	RelationshipsId, RelationshipsTypesId, RelationshipsTitlesId,
	// relationships/store
	RelationshipsStore, RelationshipsTypesStore, RelationshipsTitlesStore,
	// properties/id
	PropertiesId, PropertiesTitlesId, StringId, DoubleId,
	// properties/store
	PropertiesStore, PropertiesTitlesStore, StringStore, DoubleStore * os.File
	// file handler
	FileHandler FileHandlerInterface
)
