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
	str := s.value
	if s.nextChunk != nil {
		str += s.nextChunk.get().(string)
	}
	return str
}

func (s *StringValue) set(value interface{}) {
	str := value.(string)
	if len(str) > 31 {
		s.value = str[0:31]
		nextChunk := CreateStringValue(str[31:])
		s.nextChunk = nextChunk
	} else {
		s.value = str
	}
	s.write()
}

func CreateStringValue(value string) *StringValue {
	var s StringValue
	id, err := globals.FileHandler.ReadId(globals.StringId)
	utils.CheckError(err)
	s.id = id
	s.isUsed = true
	s.set(value)
	s.write()
	return &s
}

func (s *StringValue) GetValue() string {
	return s.get().(string)
}

func (s *StringValue) SetValue(value string) {
	s.set(value)
}

func (s *StringValue) GetNextChunk() *StringValue {
	if s.nextChunk != nil {
		return s.nextChunk
	} else {
		offset := s.id * globals.NodesSize
		bs := make([]byte, globals.NodesSize)
		globals.FileHandler.Read(globals.NodesStore, offset, bs, s.id)
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
	bs = append(utils.BoolToByteArray(s.isUsed), utils.Int32ToByteArray(int32(IfNilAssignMinusOne(s.value)))...)
	bs = append(bs, utils.Int32ToByteArray(int32(IfNilAssignMinusOne(s.nextChunk)))...)
	return bs
}

func (s *StringValue) fromBytes(bs []byte) {
	if len(bs) != globals.StringSize {
		errorMessage := fmt.Sprintf("Converter: wrong string value byte array length, " +
			"expected 36, given %d", len(bs))
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
	err = globals.FileHandler.Read(globals.StringStore, offset, bs, s.id)
	utils.CheckError(err)
	s.fromBytes(bs)
}

func (s *StringValue) write() {
	offset := globals.StringSize * s.id
	bs := s.toBytes()
	err := globals.FileHandler.Write(globals.StringStore, offset, bs, s.id)
	utils.CheckError(err)
}

// Double value
type DoubleValue struct {
	id int
	isUsed bool
	value float64
}

func (d DoubleValue) get() interface{} {
	return d.value
}

func (d *DoubleValue) set(value interface{}) {
	d.value = value.(float64)
}

func CreateDoubleValue(value float64) *DoubleValue {
	var d DoubleValue
	id, err := globals.FileHandler.ReadId(globals.DoubleId)
	utils.CheckError(err)
	d.id = id
	d.isUsed = true
	d.set(value)
	d.write()
	return &d
}

func (d *DoubleValue) GetValue() float64 {
	return d.get().(float64)
}

func (d *DoubleValue) SetValue(value float64) {
	d.set(value)
}

func (d *DoubleValue) toBytes() (bs []byte) {
	bs = append(utils.BoolToByteArray(d.isUsed), utils.Int32ToByteArray(int32(IfNilAssignMinusOne(d.value)))...)
	return bs
}

func (d *DoubleValue) fromBytes(bs []byte) {
	if len(bs) != globals.DoubleSize {
		errorMessage := fmt.Sprintf("Converter: wrong double value byte array length, " +
			"expected 9, given %d", len(bs))
		panic(errorMessage)
	}
	d.isUsed, err = utils.ByteArrayToBool(bs[0:1])
	utils.CheckError(err)

	d.value, err = utils.ByteArrayToFloat64(bs[1:9])
	utils.CheckError(err)
}

func (d *DoubleValue) read() {
	bs :=  make([]byte, globals.DoubleSize)
	offset := globals.DoubleSize * d.id
	err = globals.FileHandler.Read(globals.DoubleStore, offset, bs, d.id)
	utils.CheckError(err)
	d.fromBytes(bs)
}

func (d *DoubleValue) write() {
	offset := globals.DoubleSize * d.id
	bs := d.toBytes()
	err := globals.FileHandler.Write(globals.DoubleStore, offset, bs, d.id)
	utils.CheckError(err)
}
