package fias

import (
	"encoding/binary"
	"os"
	"github.com/davidmz/go-charset"
	"strings"
	"log"
)

type DbfHead struct {
	Version      byte
	Updatedate   [3]byte
	RecordsCount uint32
	Headerlen    uint16
	Recordlen    uint16
	Other        [20]byte
}

type DBase7Field struct {
	Name         [32]byte
	Type         byte
	FullLen      byte
	DecimalCount byte
	Reserved1    [2]byte
	TagMDX       byte
	Reserved2    [2]byte
	Incr         uint32
	Reserved3    uint32
}

type DBaseField struct {
	Name         [11]byte
	Type         byte
	Reserved1    [4]byte
	FullLen      byte
	DecimalCount byte
	Reserved2    [13]byte
	TagMDX       byte
}

func (field DBaseField) Print() {
	log.Println("Name: ", string(field.Name[:]), "    Type: ", string(field.Type))
}

func readHeader(file *os.File) (DbfHead, error) {
	var header DbfHead
	err := binary.Read(file, binary.LittleEndian, &header)
	if err != nil {
		return DbfHead{}, err
	}
	return header, nil
}

func readFields(file *os.File) ([]DBaseField, error) {
	var fields []DBaseField
	for {
		var term byte
		err := binary.Read(file, binary.LittleEndian, &term)
		if err != nil {
			return nil, err
		}
		if term == 13 {
			break
		} else {
			file.Seek(-1, os.SEEK_CUR)
		}

		var field DBaseField
		err = binary.Read(file, binary.LittleEndian, &field)
		if err != nil {
			return nil, err
		}
		fields = append(fields, field)
	}
	return fields, nil
}

func readRecords(file *os.File, count uint32, fields []DBaseField) ([][]string, error) {

	var records [][]string
	for record_id := uint32(0); record_id < count; record_id++ {
		var head_byte byte
		err := binary.Read(file, binary.LittleEndian, &head_byte)
		if err != nil {
			log.Println("Ну удалось прочитать head_byte: record_id=", record_id, "   ", err)
			return nil, err
		}
		var strs []string
		for field_id := 0; field_id < len(fields); field_id++ {
			field := fields[field_id]
			bytes := make([]byte, field.FullLen)
			err = binary.Read(file, binary.LittleEndian, &bytes)
			if err != nil {
				log.Println("Read bytes failed: rec_i=", record_id, "  filed_id=", field_id, "   ", err)
				return nil, err
			}
			var str = charset.CP866.Decode(bytes)
			str = strings.TrimSpace(str)
			strs = append(strs, str)
		}
		records = append(records, strs)
	}
	return records, nil
}

func getRecordsFromDBF(file_path string) ([][]string, error) {
	log.Println("Считываем " + file_path)
	file, err := os.Open(file_path)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer file.Close()

	header, err := readHeader(file)
	if err != nil {
		log.Println("Не могу прочитать заголовок:", err)
		return nil, err
	}

	fields, err := readFields(file)
	if err != nil {
		log.Println("Не могу прочитать дескрипторы полей:", err)
		return nil, err
	}

	records, err := readRecords(file, header.RecordsCount, fields)
	if err != nil {
		log.Println("Не могу прочитать записи:", err)
		return nil, err
	}

	return records, nil
}
