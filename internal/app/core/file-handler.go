package core

import (
	"fmt"
	"os"
	"path/filepath"
	"errors"
	"graph-db/internal/pkg/utils"
	"io/ioutil"
	"strings"
	"strconv"
)

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

func initFileSystem() {
	if _, err := os.Stat(rootPath); os.IsNotExist(err) {
		os.Mkdir(rootPath, os.ModePerm)
	}
}

func initDatabaseStructure(dbTitle string) {
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
	utils.CheckError(err)
	labelsId, err = os.Create(filepath.Join(nodesIdPath, "labels.id"))
	utils.CheckError(err)
	labelsTitlesId, err = os.Create(filepath.Join(nodesIdPath, "labelsTitles.id"))
	utils.CheckError(err)
	// nodes/store
	nodesStore, err = os.Create(filepath.Join(nodesStorePath, "nodes.store"))
	utils.CheckError(err)
	labelsStore, err = os.Create(filepath.Join(nodesStorePath, "labels.store"))
	utils.CheckError(err)
	labelsTitlesStore, err = os.Create(filepath.Join(nodesStorePath, "labelsTitles.store"))
	utils.CheckError(err)

	// relationships/id
	relationshipsId, err = os.Create(filepath.Join(relationshipsIdPath, "relationships.id"))
	utils.CheckError(err)
	relationshipsTypesId, err = os.Create(filepath.Join(relationshipsIdPath, "relationshipsTypes.id"))
	utils.CheckError(err)
	// relationships/store
	relationshipsStore, err = os.Create(filepath.Join(relationshipsStorePath, "relationships.store"))
	utils.CheckError(err)
	relationshipsTypesStore, err = os.Create(filepath.Join(relationshipsStorePath, "relationshipsTypes.store"))
	utils.CheckError(err)

	// properties/id
	propertiesId, err = os.Create(filepath.Join(propertiesIdPath, "properties.id"))
	utils.CheckError(err)
	propertiesTitlesId, err = os.Create(filepath.Join(propertiesIdPath, "propertiesTitles.id"))
	utils.CheckError(err)
	stringId, err = os.Create(filepath.Join(propertiesIdPath, "string.id"))
	utils.CheckError(err)
	doubleId, err = os.Create(filepath.Join(propertiesIdPath, "double.id"))
	utils.CheckError(err)
	// properties/store
	propertiesStore, err = os.Create(filepath.Join(propertiesStorePath, "properties.store"))
	utils.CheckError(err)
	propertiesTitlesStore, err = os.Create(filepath.Join(propertiesStorePath, "propertiesTitles.store"))
	utils.CheckError(err)
	stringStore, err = os.Create(filepath.Join(propertiesStorePath, "string.store"))
	utils.CheckError(err)
	doubleStore, err = os.Create(filepath.Join(propertiesStorePath, "double.store"))
	utils.CheckError(err)

	nodesId.WriteString(fmt.Sprintf("%d", 0))
	labelsId.WriteString(fmt.Sprintf("%d", 0))
	labelsTitlesId.WriteString(fmt.Sprintf("%d", 0))

	relationshipsId.WriteString(fmt.Sprintf("%d", 0))
	relationshipsTypesId.WriteString(fmt.Sprintf("%d", 0))

	propertiesId.WriteString(fmt.Sprintf("%d", 0))
	propertiesTitlesId.WriteString(fmt.Sprintf("%d", 0))
	stringId.WriteString(fmt.Sprintf("%d", 0))
	doubleId.WriteString(fmt.Sprintf("%d", 0))
}

func write(file *os.File, offset int, bs []byte) (err error) {
	offset = offset * len(bs)
	bytesWritten, err := file.WriteAt(bs, int64(offset))
	if bytesWritten != len(bs) {
		err = errors.New("write: wrote less bytes than expected")
	}
	return err
}

func read(file *os.File, offset int, bs []byte) (err error) {
	offset = offset * len(bs)
	bytesRead, err := file.ReadAt(bs, int64(offset))
	if bytesRead != len(bs) {
		err = errors.New("read: read less bytes than expected")
	}
	return err
}

func readId(file *os.File) (id int, err error) {
	fileData, err := ioutil.ReadFile(file.Name())
	if err == nil {
		ids := strings.Split(string(fileData), "\n")
		id, err := strconv.Atoi(ids[0])
		if err == nil {
			if len(ids) == 1 {
				str := strconv.Itoa(id + 1)
				err := ioutil.WriteFile(file.Name(), []byte(str), os.ModePerm)
				if err == nil {
					return id, err
				}
			} else {
				str := strings.Join(ids[1:], "\n")
				err := ioutil.WriteFile(file.Name(), []byte(str), os.ModePerm)
				if err == nil {
					return id, err
				}
			}
		}
	}
	return 0, err
}