package fias

import (
	"log"
	"errors"
)

type SockBaseRecod struct {
	LEVEL   string
	SCNAME   string
	SOCRNAME  string
	KOD_T_ST string
}

func (rec *SockBaseRecod) From(strs []string) error {
	if len(strs) != 4 {
		return errors.New("Некорректное количество полей в SockBaseRecod")
	}

	rec.LEVEL = strs[0]
	rec.SCNAME = strs[1]
	rec.SOCRNAME = strs[2]
	rec.KOD_T_ST = strs[3]

	return nil
}

func (rec SockBaseRecod) Print() {
	log.Print("LEVEL: '", rec.LEVEL, "', ")
	log.Print("SCNAME: '", rec.SCNAME, "', ")
	log.Print("SOCRNAME: '", rec.SOCRNAME, "', ")
	log.Println("KOD_T_ST: '", rec.KOD_T_ST, "'")
}

func ImportSockBase(file string) ([]SockBaseRecod, error) {
	records, err := getRecordsFromDBF(file)
	if err != nil {
		log.Println("Не могу прочитать записи:", err)
		return nil, err
	}

	var res []SockBaseRecod
	for _, record := range records {
		var rec SockBaseRecod
		err := rec.From(record)
		if err != nil {
			log.Println("Не могу получить SockBaseRecod из массива строк:", err)
			return nil, err
		}
		res = append(res, rec)
	}

	return res, err
}
