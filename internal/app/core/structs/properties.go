package structs

import (
	"graph-db/internal/app/core/globals"
	"graph-db/internal/pkg/utils"
	"fmt"
	"errors"
)

type Property struct {
	id int
	isUsed bool
	isWritten bool
	nextProperty *Property
	title *PropertyTitle
	valueType int8
	value *Value

	byteString []byte
}

type PropertyTitle struct {
	id int
	title string
	counter int
}

func CreateProperty() *Property {
	var p Property
	id, err := globals.FileHandler.ReadId(globals.PropertiesId)
	utils.CheckError(err)
	p.id = id
	p.isUsed = true
	p.isWritten = false
	p.write()
	return &p
}

func (p *Property) Delete(id int) (err error) {
	bs := make([]byte, globals.PropertiesSize)
	bs[0] = utils.BoolToByteArray(false)[0]
	err = globals.FileHandler.FreeId(globals.PropertiesId, id)
	if err != nil {
		return err
	}
	offset := globals.PropertiesSize * id
	err = globals.FileHandler.Write(globals.PropertiesStore, offset, bs, id)
	return err
}

func (p *Property) Get(id int) Property {
	p.id = id
	p.read()
	p.isWritten = true
	return *p
}

func (p Property) GetId() int {
	return p.id
}

func (p Property) GetNextProperty() *Property {
	if p.nextProperty != nil {
		return p.nextProperty
	} else if !p.isWritten {
		return nil
	} else {
		var (
			propertyId int32
			err error
			bs = make([]byte, globals.PropertiesSize)
		)
		if len(p.byteString) < 0 {
			offset := p.id * globals.PropertiesSize
			err = globals.FileHandler.Read(globals.PropertiesStore, offset, bs, p.id)
			utils.CheckError(err)
			p.byteString = bs
		}
		propertyId, err = utils.ByteArrayToInt32(p.byteString[1:5])
		utils.CheckError(err)
		if propertyId == -1 {
			return nil
		} else {
			var property Property
			property.id = int(propertyId)
			p.nextProperty = &property
			return p.nextProperty
		}
	}
}

func (p *Property) SetNextProperty(nextProperty *Property) {
	p.nextProperty = nextProperty
	p.write()
}

func (p *Property) GetTitle() *PropertyTitle {
	if p.title != nil {
		return p.title
	} else if !p.isWritten {
		return nil
	} else {
		var (
			titleId int32
			err error
			bs = make([]byte, globals.PropertiesSize)
		)
		if len(p.byteString) < 0 {
			offset := p.id * globals.PropertiesSize
			err = globals.FileHandler.Read(globals.PropertiesStore, offset, bs, p.id)
			utils.CheckError(err)
			p.byteString = bs
		}
		titleId, err = utils.ByteArrayToInt32(p.byteString[5:9])
		utils.CheckError(err)
		if titleId == -1 {
			return nil
		} else {
			var title PropertyTitle
			title.id = int(titleId)
			p.title = &title
			return p.title
		}
	}
}

func (p *Property) SetTitle(title *PropertyTitle) {
	p.title = title
	p.write()
}

func (p *Property) GetValueType() int8 {
	return p.valueType
}

func (p *Property) SetValueType(valueType int8) {
	p.valueType = valueType
	p.write()
}

func (p *Property) GetValue() *Value {
	if p.value != nil {
		return p.value
	} else if !p.isWritten {
		return nil
	} else {
		var (
			err error
			bs = make([]byte, globals.PropertiesSize)
			value Value
		)
		if len(p.byteString) < 0 {
			offset := p.id * globals.PropertiesSize
			err = globals.FileHandler.Read(globals.PropertiesStore, offset, bs, p.id)
			utils.CheckError(err)
			p.byteString = bs
		}
		val, err := utils.ByteArrayToInt32(p.byteString[10:14])
		utils.CheckError(err)
		store := GetValueFile(p.valueType)
		if val == -1 {
			return nil
		} else {
			if p.valueType == 0 {
				value = &IntegerValue{value: int(val)}
			} else {
				if p.valueType == 1 {
					bs := make([]byte, globals.DoubleSize)
					offset := int(val * globals.DoubleSize)
					err := globals.FileHandler.Read(store, offset, bs, int(val))
					utils.CheckError(err)
					fileValue, err := utils.ByteArrayToFloat64(bs)
					utils.CheckError(err)
					value = &DoubleValue{id: int(val), value: fileValue}
				} else {
					bs := make([]byte, globals.StringSize)
					offset := int(val * globals.StringSize)
					err := globals.FileHandler.Read(store, offset, bs, int(val))
					utils.CheckError(err)
					fileValue := utils.ByteArrayToString(bs)
					utils.CheckError(err)
					value = &StringValue{id: int(val), value: fileValue}
				}
			}
		}
		p.value = &value
		return &value
	}
}

func (p *Property) SetValue(value *Value) {
	p.value = value
	p.write()
}

func (p *Property) toBytes() (bs []byte) {

	var (
		isUsed []byte
		nextProperty *Property
		title *PropertyTitle
		valueType int8
		value *Value
	)

	nextProperty = p.GetNextProperty()
	title = p.GetTitle()
	valueType = p.GetValueType()
	value = p.GetValue()

	isUsed = utils.BoolToByteArray(p.isUsed)
	nextPropertyBs := utils.Int32ToByteArray(int32(IfNilAssignMinusOne(nextProperty)))
	titleBs := utils.Int32ToByteArray(int32(IfNilAssignMinusOne(title)))
	valueTypeBs := utils.Int8ToByteArray(valueType)
	valueBs := utils.Int32ToByteArray(int32(IfNilAssignMinusOne(value)))

	bs = append(isUsed, nextPropertyBs...)
	bs = append(bs, titleBs...)
	bs = append(bs, valueTypeBs...)
	bs = append(bs, valueBs...)

	return bs
}

func (p *Property) fromBytes(bs []byte) {
	var (
		id int32
		nextProperty Property
		title PropertyTitle
		valueType int8
		value Value
	)

	if len(bs) != globals.PropertiesSize {
		errorMessage := fmt.Sprintf("Converter: wrong properties byte array length, expected 14, given %d", len(bs))
		panic(errorMessage)
	}
	p.byteString = bs
	p.isUsed, err = utils.ByteArrayToBool(bs[0:1])
	utils.CheckError(err)

	id, err = utils.ByteArrayToInt32(bs[1:5])
	utils.CheckError(err)
	nextProperty.id = int(id)
	p.nextProperty = &nextProperty

	id, err = utils.ByteArrayToInt32(bs[5:9])
	utils.CheckError(err)
	title.id = int(id)
	p.title = &title

	vType := utils.ByteArrayToInt8(bs[9:10])
	valueType = int8(vType)
	p.valueType = valueType

	val, err := utils.ByteArrayToInt32(bs[10:14])
	utils.CheckError(err)
	store := GetValueFile(p.valueType)
	if val == 0 {
		value = &IntegerValue{value: int(val)}
	} else if p.valueType == 1 {
		bs := make([]byte, globals.DoubleSize)
		offset := int(val * globals.DoubleSize)
		err := globals.FileHandler.Read(store, offset, bs, int(val))
		utils.CheckError(err)
		fileValue, err := utils.ByteArrayToFloat64(bs)
		utils.CheckError(err)
		value = &DoubleValue{id: int(val), value: fileValue}
	} else {
		bs := make([]byte, globals.StringSize)
		offset := int(val * globals.StringSize)
		err := globals.FileHandler.Read(store, offset, bs, int(val))
		utils.CheckError(err)
		fileValue := utils.ByteArrayToString(bs)
		utils.CheckError(err)
		value = &StringValue{id: int(val), value: fileValue}
	}
	p.value = &value
}

func (p *Property) read() {
	bs := make([]byte, globals.PropertiesSize)
	offset := globals.PropertiesSize * p.id
	err = globals.FileHandler.Read(globals.PropertiesStore, offset, bs, p.id)
	utils.CheckError(err)
	p.fromBytes(bs)
}

func (p *Property) write() {
	offset := globals.PropertiesSize * p.id
	bs := p.toBytes()
	err = globals.FileHandler.Write(globals.PropertiesStore, offset, bs, p.id)
	utils.CheckError(err)
	p.isWritten = true
}

func WritePropertyTitle(id int, title string, counter int) {
	offset := id * globals.PropertiesTitlesSize
	bs := make([]byte, globals.PropertiesTitlesSize)
	titleBs := utils.StringToByteArray(utils.AddStopCharacter(title, globals.PropertiesTitlesSize - 4))
	for i := 0; i < len(titleBs); i++ {
		bs[i] = titleBs[i]
	}
	counterBs := utils.Int32ToByteArray(int32(counter))
	for i := 0; i < 4; i++ {
		bs[globals.PropertiesTitlesSize - 4 + i] = counterBs[i]
	}
	err := globals.FileHandler.Write(globals.PropertiesTitlesStore, offset, bs, id)
	utils.CheckError(err)
}

func DecreasePropertyTitleCounter(title string) {
	value := globals.PropertyTitleMap[title]
	value.Counter--
	globals.PropertyTitleMap[title] = value
	WritePropertyTitle(value.Id, title, value.Counter)
	if globals.PropertyTitleMap[title].Counter == 0 {
		delete(globals.PropertyTitleMap, title)
		globals.FileHandler.FreeId(globals.PropertiesId, value.Id)
	}
}

func AddPropertyTitle(title string) *PropertyTitle {
	if len(title) > globals.PropertiesTitlesSize - 4 {
		err = errors.New("property title is too big")
	} else {
		value, present := globals.PropertyTitleMap[title]
		if present {
			value.Counter++
			globals.PropertyTitleMap[title] = value
		} else {
			id, err := globals.FileHandler.ReadId(globals.PropertiesTitlesId)
			utils.CheckError(err)
			value = globals.MapValue{Counter: 1, Id: id}
			globals.PropertyTitleMap[title] = value
		}
		WritePropertyTitle(value.Id, title, value.Counter)
		return &PropertyTitle{id: value.Id, title: title, counter: value.Counter}
	}
}