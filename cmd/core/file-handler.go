package core

import "os"
import "path/filepath"

var (
	// nodes/id
	nodesId, labelsId, labelsTitlesId,
	// nodes/store
 	nodesStore, labelsStore, labelsTitlesStore,
 	// relationships/id
	relationshipsId, relationshipsTypesId,
	// relationships/store
	relationshipsStore, relationshipsTypesStore,
	// properties/id
	propertiesId, propertiesTitlesId, stringId, doubleId,
	// properties/store
	propertiesStore, propertiesTitlesStore, stringStore, doubleStore * os.File
	err error
)

var rootPath = "databases"

func InitFileSystem() {
	if _, err := os.Stat(rootPath); os.IsNotExist(err) {
		os.Mkdir(rootPath, os.ModePerm)
	}
}

func InitDatabaseStructure(dbTitle string) {
	var storagePath = filepath.Join(rootPath, dbTitle, "storage")
	var nodesPath = filepath.Join(storagePath, "nodes")
	var nodesIdPath = filepath.Join(nodesPath, "id")
	var nodesStorePath = filepath.Join(nodesPath, "store")
	var relationshipsPath = filepath.Join(storagePath, "relationships")
	var relationshipsIdPath = filepath.Join(relationshipsPath, "id")
	var relationshipsStorePath = filepath.Join(relationshipsPath, "store")
	var propertiesPath = filepath.Join(storagePath, "properties")
	var propertiesIdPath = filepath.Join(propertiesPath, "id")
	var propertiesStorePath = filepath.Join(propertiesPath, "store")

	os.MkdirAll(nodesIdPath, os.ModePerm)
	os.MkdirAll(nodesStorePath, os.ModePerm)
	os.MkdirAll(relationshipsIdPath, os.ModePerm)
	os.MkdirAll(relationshipsStorePath, os.ModePerm)
	os.MkdirAll(propertiesIdPath, os.ModePerm)
	os.MkdirAll(propertiesStorePath, os.ModePerm)

	// nodes/id
	nodesId, err = os.Create(filepath.Join(nodesIdPath, "nodes.id"))
	checkError(err)
	labelsId, err = os.Create(filepath.Join(nodesIdPath, "labels.id"))
	checkError(err)
	labelsTitlesId, err = os.Create(filepath.Join(nodesIdPath, "labelsTitles.id"))
	checkError(err)
	// nodes/store
	nodesStore, err = os.Create(filepath.Join(nodesStorePath, "nodes.store"))
	checkError(err)
	labelsStore, err = os.Create(filepath.Join(nodesStorePath, "labels.store"))
	checkError(err)
	labelsTitlesStore, err = os.Create(filepath.Join(nodesStorePath, "labelsTitles.store"))
	checkError(err)

	// relationships/id
	relationshipsId, err = os.Create(filepath.Join(relationshipsIdPath, "relationships.id"))
	checkError(err)
	relationshipsTypesId, err = os.Create(filepath.Join(relationshipsIdPath, "relationshipsTypes.id"))
	checkError(err)
	// relationships/store
	relationshipsStore, err = os.Create(filepath.Join(relationshipsStorePath, "relationships.store"))
	checkError(err)
	relationshipsTypesStore, err = os.Create(filepath.Join(relationshipsStorePath, "relationshipsTypes.store"))
	checkError(err)

	// properties/id
	propertiesId, err = os.Create(filepath.Join(propertiesIdPath, "properties.id"))
	checkError(err)
	propertiesTitlesId, err = os.Create(filepath.Join(propertiesIdPath, "propertiesTitles.id"))
	checkError(err)
	stringId, err = os.Create(filepath.Join(propertiesIdPath, "string.id"))
	checkError(err)
	doubleId, err = os.Create(filepath.Join(propertiesIdPath, "double.id"))
	checkError(err)
	// properties/store
	propertiesStore, err = os.Create(filepath.Join(propertiesStorePath, "properties.store"))
	checkError(err)
	propertiesTitlesStore, err = os.Create(filepath.Join(propertiesStorePath, "propertiesTitles.store"))
	checkError(err)
	stringStore, err = os.Create(filepath.Join(propertiesStorePath, "string.store"))
	checkError(err)
	doubleStore, err = os.Create(filepath.Join(propertiesStorePath, "double.store"))
	checkError(err)
}