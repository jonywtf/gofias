package fias

import (
	"log"
	"errors"
)

type FlatRecod struct {
	NAME   string
	CODE   string
	INDEX  string
	GNINMB string
	UNO    string
	NP     string
}

func (rec *FlatRecod) From(strs []string) error {
	if len(strs) != 6 {
		return errors.New("Некорректное количество полей в FlatRecod")
	}

	rec.NAME = strs[0]
	rec.CODE = strs[1]
	rec.INDEX = strs[2]
	rec.GNINMB = strs[3]
	rec.UNO = strs[4]
	rec.NP = strs[5]

	return nil
}

func (rec FlatRecod) Print() {
	log.Print("NAME: '", rec.NAME, "', ")
	log.Print("CODE: '", rec.CODE, "', ")
	log.Print("INDEX: '", rec.INDEX, "', ")
	log.Print("GNINMB: '", rec.GNINMB, "', ")
	log.Print("UNO: '", rec.UNO, "', ")
	log.Println("NP: '", rec.NP, "'")
}

func ImportFlat(file string) ([]FlatRecod, error) {
	records, err := getRecordsFromDBF(file)
	if err != nil {
		log.Println("Не могу прочитать записи:", err)
		return nil, err
	}

	var res []FlatRecod
	for _, record := range records {
		var rec FlatRecod
		err := rec.From(record)
		if err != nil {
			log.Println("Не могу получить FlatRecod из массива строк:", err)
			return nil, err
		}
		res = append(res, rec)
	}

	return res, err
}
