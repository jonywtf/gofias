package fias

import (
	"log"
	"errors"
)

type KladrRecod struct {
	NAME   string
	SOCR   string
	CODE   string
	INDEX  string
	GNINMB string
	UNO    string
	OCATD  string
	STATUS  string
}

func (rec *KladrRecod) From(strs []string) error {
	if len(strs) != 8 {
		return errors.New("Некорректное количество полей в KladrRecod")
	}

	rec.NAME = strs[0]
	rec.SOCR = strs[1]
	rec.CODE = strs[2]
	rec.INDEX = strs[3]
	rec.GNINMB = strs[4]
	rec.UNO = strs[5]
	rec.OCATD = strs[6]
	rec.STATUS = strs[7]

	return nil
}

func (rec KladrRecod) Print() {
	log.Print("NAME: '", rec.NAME, "', ")
	log.Print("SOCR: '", rec.SOCR, "', ")
	log.Print("CODE: '", rec.CODE, "', ")
	log.Print("INDEX: '", rec.INDEX, "', ")
	log.Print("GNINMB: '", rec.GNINMB, "', ")
	log.Print("UNO: '", rec.UNO, "', ")
	log.Print("OCATD: '", rec.OCATD, "', ")
	log.Println("STATUS: '", rec.STATUS, "'")
}

func ImportKladr(file string) ([]KladrRecod, error) {
	records, err := getRecordsFromDBF(file)
	if err != nil {
		log.Println("Не могу прочитать записи:", err)
		return nil, err
	}

	var res []KladrRecod
	for _, record := range records {
		var rec KladrRecod
		err := rec.From(record)
		if err != nil {
			log.Println("Не могу получить KladrRecod из массива строк:", err)
			return nil, err
		}
		res = append(res, rec)
	}

	return res, err
}
