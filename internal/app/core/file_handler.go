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
	"graph-db/internal/app/core/globals"
)

var (
	rootPath = "databases"
	err error
)

type FileHandler struct {
}

func (fh FileHandler) InitFileSystem() {
	if _, err := os.Stat(rootPath); os.IsNotExist(err) {
		os.Mkdir(rootPath, os.ModePerm)
	}
}

func (fh FileHandler) InitDatabaseStructure(dbTitle string) {
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
	globals.NodesId, err = os.Create(filepath.Join(nodesIdPath, "nodes.id"))
	utils.CheckError(err)
	globals.LabelsId, err = os.Create(filepath.Join(nodesIdPath, "labels.id"))
	utils.CheckError(err)
	globals.LabelsTitlesId, err = os.Create(filepath.Join(nodesIdPath, "labelsTitles.id"))
	utils.CheckError(err)
	// nodes/store
	globals.NodesStore, err = os.Create(filepath.Join(nodesStorePath, "nodes.store"))
	utils.CheckError(err)
	globals.LabelsStore, err = os.Create(filepath.Join(nodesStorePath, "labels.store"))
	utils.CheckError(err)
	globals.LabelsTitlesStore, err = os.Create(filepath.Join(nodesStorePath, "labelsTitles.store"))
	utils.CheckError(err)

	// relationships/id
	globals.RelationshipsId, err = os.Create(filepath.Join(relationshipsIdPath, "relationships.id"))
	utils.CheckError(err)
	globals.RelationshipsTypesId, err = os.Create(filepath.Join(relationshipsIdPath, "relationshipsTypes.id"))
	utils.CheckError(err)
	// relationships/store
	globals.RelationshipsStore, err = os.Create(filepath.Join(relationshipsStorePath, "relationships.store"))
	utils.CheckError(err)
	globals.RelationshipsTypesStore, err = os.Create(filepath.Join(relationshipsStorePath, "relationshipsTypes.store"))
	utils.CheckError(err)

	// properties/id
	globals.PropertiesId, err = os.Create(filepath.Join(propertiesIdPath, "properties.id"))
	utils.CheckError(err)
	globals.PropertiesTitlesId, err = os.Create(filepath.Join(propertiesIdPath, "propertiesTitles.id"))
	utils.CheckError(err)
	globals.StringId, err = os.Create(filepath.Join(propertiesIdPath, "string.id"))
	utils.CheckError(err)
	globals.DoubleId, err = os.Create(filepath.Join(propertiesIdPath, "double.id"))
	utils.CheckError(err)
	// properties/store
	globals.PropertiesStore, err = os.Create(filepath.Join(propertiesStorePath, "properties.store"))
	utils.CheckError(err)
	globals.PropertiesTitlesStore, err = os.Create(filepath.Join(propertiesStorePath, "propertiesTitles.store"))
	utils.CheckError(err)
	globals.StringStore, err = os.Create(filepath.Join(propertiesStorePath, "string.store"))
	utils.CheckError(err)
	globals.DoubleStore, err = os.Create(filepath.Join(propertiesStorePath, "double.store"))
	utils.CheckError(err)

	globals.NodesId.WriteString(fmt.Sprintf("%d", 0))
	globals.LabelsId.WriteString(fmt.Sprintf("%d", 0))
	globals.LabelsTitlesId.WriteString(fmt.Sprintf("%d", 0))

	globals.RelationshipsId.WriteString(fmt.Sprintf("%d", 0))
	globals.RelationshipsTypesId.WriteString(fmt.Sprintf("%d", 0))

	globals.PropertiesId.WriteString(fmt.Sprintf("%d", 0))
	globals.PropertiesTitlesId.WriteString(fmt.Sprintf("%d", 0))
	globals.StringId.WriteString(fmt.Sprintf("%d", 0))
	globals.DoubleId.WriteString(fmt.Sprintf("%d", 0))
}

func (fh FileHandler) SwitchDatabaseStructure(dbTitle string) (err error) {
	if _, err := os.Stat(filepath.Join(rootPath, dbTitle)); err == nil {
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

		// nodes/id
		globals.NodesId, err = os.Open(filepath.Join(nodesIdPath, "nodes.id"))
		utils.CheckError(err)
		globals.LabelsId, err = os.Open(filepath.Join(nodesIdPath, "labels.id"))
		utils.CheckError(err)
		globals.LabelsTitlesId, err = os.Open(filepath.Join(nodesIdPath, "labelsTitles.id"))
		utils.CheckError(err)
		// nodes/store
		globals.NodesStore, err = os.Open(filepath.Join(nodesStorePath, "nodes.store"))
		utils.CheckError(err)
		globals.LabelsStore, err = os.Open(filepath.Join(nodesStorePath, "labels.store"))
		utils.CheckError(err)
		globals.LabelsTitlesStore, err = os.Open(filepath.Join(nodesStorePath, "labelsTitles.store"))
		utils.CheckError(err)

		// relationships/id
		globals.RelationshipsId, err = os.Open(filepath.Join(relationshipsIdPath, "relationships.id"))
		utils.CheckError(err)
		globals.RelationshipsTypesId, err = os.Open(filepath.Join(relationshipsIdPath, "relationshipsTypes.id"))
		utils.CheckError(err)
		// relationships/store
		globals.RelationshipsStore, err = os.Open(filepath.Join(relationshipsStorePath, "relationships.store"))
		utils.CheckError(err)
		globals.RelationshipsTypesStore, err = os.Open(filepath.Join(relationshipsStorePath, "relationshipsTypes.store"))
		utils.CheckError(err)

		// properties/id
		globals.PropertiesId, err = os.Open(filepath.Join(propertiesIdPath, "properties.id"))
		utils.CheckError(err)
		globals.PropertiesTitlesId, err = os.Open(filepath.Join(propertiesIdPath, "propertiesTitles.id"))
		utils.CheckError(err)
		globals.StringId, err = os.Open(filepath.Join(propertiesIdPath, "string.id"))
		utils.CheckError(err)
		globals.DoubleId, err = os.Open(filepath.Join(propertiesIdPath, "double.id"))
		utils.CheckError(err)
		// properties/store
		globals.PropertiesStore, err = os.Open(filepath.Join(propertiesStorePath, "properties.store"))
		utils.CheckError(err)
		globals.PropertiesTitlesStore, err = os.Open(filepath.Join(propertiesStorePath, "propertiesTitles.store"))
		utils.CheckError(err)
		globals.StringStore, err = os.Open(filepath.Join(propertiesStorePath, "string.store"))
		utils.CheckError(err)
		globals.DoubleStore, err = os.Open(filepath.Join(propertiesStorePath, "double.store"))
		utils.CheckError(err)

		return err
	} else {
		return errors.New(fmt.Sprintf("Database with title %s does not exist", dbTitle))
	}
}

func (fh FileHandler) DropDatabase(dbTitle string) (err error) {
	if _, err := os.Stat(filepath.Join(rootPath, dbTitle)); err == nil {
		err = os.RemoveAll(filepath.Join(rootPath, dbTitle))
		return err
	} else {
		return errors.New(fmt.Sprintf("Database with title %s does not exist", dbTitle))
	}
}

func (fh FileHandler) Write(file *os.File, offset int, bs []byte) (err error) {
	offset = offset * len(bs)
	bytesWritten, err := file.WriteAt(bs, int64(offset))
	if bytesWritten != len(bs) {
		err = errors.New("write: wrote less bytes than expected")
	}
	return err
}

func (fh FileHandler) Read(file *os.File, offset int, bs []byte) (err error) {
	offset = offset * len(bs)
	bytesRead, err := file.ReadAt(bs, int64(offset))
	if bytesRead != len(bs) {
		err = errors.New("read: read less bytes than expected")
	}
	return err
}

func (fh FileHandler) ReadId(file *os.File) (id int, err error) {
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

func (fh FileHandler) FreeId(file *os.File, id int) (err error) {
	fileData, err := ioutil.ReadFile(file.Name())
	if err == nil {
		ids := strings.Split(string(fileData), "\n")
		str := strconv.Itoa(id)
		var fetchedPrev, fetchedNext, firstId, lastId int
		firstId, err = strconv.Atoi(ids[0])
		utils.CheckError(err)
		lastId, err = strconv.Atoi(ids[len(ids) - 1])
		utils.CheckError(err)
		if id < firstId {
			ids = append([]string{str}, ids[:]...)
		} else if id > lastId {
			return errors.New("Bad id (specified id is out of range)")
		} else if id == lastId {
			return errors.New("Bad id (specified id is already free)")
		} else {
			for i := 0; i < len(ids) - 1; i++ {
				fetchedPrev, err = strconv.Atoi(ids[i])
				utils.CheckError(err)
				fetchedNext, err = strconv.Atoi(ids[i + 1])
				utils.CheckError(err)
				if id == fetchedPrev {
					return errors.New("Bad id (specified id is already free)")
				}
				if id > fetchedPrev && id < fetchedNext {
					ids = append(ids[:i + 1], append([]string{str}, ids[i + 1:]...)...)
					break
				}
			}
		}
		str = strings.Join(ids, "\n")
		err = ioutil.WriteFile(file.Name(), []byte(str), os.ModePerm)
	}
	return err
}