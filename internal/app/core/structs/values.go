package structs

import (
	"graph-db/internal/app/core/globals"
	"graph-db/internal/pkg/utils"
	"fmt"
)

type Value interface {
	get() interface{}
	set(value interface{})
}

// Integer value
type IntegerValue struct {
	value int
}

func (i IntegerValue) get() interface{} {
	return i.value
}

func (i *IntegerValue) set(value interface{}) {
	i.value = value.(int)
}

// String value
type StringValue struct {
	id int
	isUsed bool
	value string
	nextChunk *StringValue
}

func (s StringValue) get() interface{} {
	return s.value
}

func (s *StringValue) set(value interface{}) {
	s.value = value.(string)
}

func (s *StringValue) GetValue() string {
	return s.value
}

func (s *StringValue) GetNextChunk() *StringValue {
	if s.nextChunk != nil {
		return s.nextChunk
	} else {
		offset := s.id * globals.NodesSize
		bs := make([]byte, globals.NodesSize)
		globals.FileHandler.Read(globals.NodesStore, offset, bs)
		nextChunkId, err := utils.ByteArrayToInt32(bs[1:5])
		utils.CheckError(err)
		if nextChunkId == -1 {
			return nil
		} else {
			var nextChunk StringValue
			nextChunk.id = int(nextChunkId)
			s.nextChunk = &nextChunk
			return s.nextChunk
		}
	}
}

func (s *StringValue) toBytes() (bs []byte) {
	bs = append(utils.BoolToByteArray(s.isUsed),
		utils.StringToByteArray(utils.AddStopCharacter(s.value, 31))...)
	bs = append(bs, utils.Int32ToByteArray(int32(IfNilAssignMinusOne(s.nextChunk)))...)
	return bs
}

func (s *StringValue) fromBytes(bs []byte) {
	if len(bs) != globals.StringSize {
		errorMessage := fmt.Sprintf("Converter: wrong string value byte array length, expected 36, given %d", len(bs))
		panic(errorMessage)
	}
	isUsed, err := utils.ByteArrayToBool(bs[0:1])
	utils.CheckError(err)
	s.isUsed = isUsed
	s.value = utils.RemoveStopCharacter(utils.ByteArrayToString(bs[1:32]))
	id, err := utils.ByteArrayToInt32(bs[32:36])
	utils.CheckError(err)
	if id != -1 {
		var nextChunk StringValue
		nextChunk.id = int(id)
		s.nextChunk = &nextChunk
	} else {
		s.nextChunk = nil
	}
}

func (s *StringValue) read() {
	bs :=  make([]byte, globals.StringSize)
	offset := globals.StringSize * s.id
	err = globals.FileHandler.Read(globals.StringStore, offset, bs)
	utils.CheckError(err)
	s.fromBytes(bs)
}

func (s *StringValue) write() {
	offset := globals.StringSize * s.id
	bs := s.toBytes()
	err := globals.FileHandler.Write(globals.StringStore, offset, bs)
	utils.CheckError(err)
}

// Double value
type DoubleValue struct {
	id int
	value float64
}

func (d DoubleValue) get() interface{} {
	return d.value
}

func (d *DoubleValue) set(value interface{}) {
	d.value = value.(float64)
}
