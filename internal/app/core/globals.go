package core

import "os"

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
