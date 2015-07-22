package fias

import (
	"log"
	"errors"
)

type DomRecod struct {
	NAME   string
	KORP   string
	SOCR   string
	CODE   string
	INDEX  string
	GNINMB string
	UNO    string
	OCATD  string
}

func (rec *DomRecod) From(strs []string) error {
	if len(strs) != 8 {
		return errors.New("Некорректное количество полей в DomaRecod")
	}

	rec.NAME = strs[0]
	rec.KORP = strs[1]
	rec.SOCR = strs[2]
	rec.CODE = strs[3]
	rec.INDEX = strs[4]
	rec.GNINMB = strs[5]
	rec.UNO = strs[6]
	rec.OCATD = strs[7]

	return nil
}

func (rec DomRecod) Print() {
	log.Print("NAME: '", rec.NAME, "', ")
	log.Print("KORP: '", rec.KORP, "', ")
	log.Print("SOCR: '", rec.SOCR, "', ")
	log.Print("CODE: '", rec.CODE, "', ")
	log.Print("INDEX: '", rec.INDEX, "', ")
	log.Print("GNINMB: '", rec.GNINMB, "', ")
	log.Print("UNO: '", rec.UNO, "', ")
	log.Println("OCATD: '", rec.OCATD, "'")
}

func ImportDoma(file string) ([]DomRecod, error) {
	records, err := getRecordsFromDBF(file)
	if err != nil {
		log.Println("Не могу прочитать записи:", err)
		return nil, err
	}

	var res []DomRecod
	for _, record := range records {
		var rec DomRecod
		err := rec.From(record)
		if err != nil {
			log.Println("Не могу получить DomRecod из массива строк:", err)
			return nil, err
		}
		res = append(res, rec)
	}

	return res, err
}
