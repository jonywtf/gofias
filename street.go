package fias

import (
	"log"
	"errors"
)

type StreetRecod struct {
	NAME   string
	SOCR   string
	CODE   string
	INDEX  string
	GNINMB string
	UNO    string
	OCATD  string
}

func (rec *StreetRecod) From(strs []string) error {
	if len(strs) != 7 {
		return errors.New("Некорректное количество полей в StreetRecod")
	}

	rec.NAME = strs[0]
	rec.SOCR = strs[1]
	rec.CODE = strs[2]
	rec.INDEX = strs[3]
	rec.GNINMB = strs[4]
	rec.UNO = strs[5]
	rec.OCATD = strs[6]

	return nil
}

func (rec StreetRecod) Print() {
	log.Print("NAME: '", rec.NAME, "', ")
	log.Print("SOCR: '", rec.SOCR, "', ")
	log.Print("CODE: '", rec.CODE, "', ")
	log.Print("INDEX: '", rec.INDEX, "', ")
	log.Print("GNINMB: '", rec.GNINMB, "', ")
	log.Print("UNO: '", rec.UNO, "', ")
	log.Println("OCATD: '", rec.OCATD, "'")
}


func ImportStreet(file string) ([]StreetRecod, error) {
	records, err := getRecordsFromDBF(file)
	if err != nil {
		log.Println("Не могу прочитать записи:", err)
		return nil, err
	}

	var res []StreetRecod
	for _, record := range records {
		var rec StreetRecod
		err := rec.From(record)
		if err != nil {
			log.Println("Не могу получить StreetRecod из массива строк:", err)
			return nil, err
		}
		res = append(res, rec)
	}

	return res, err
}
