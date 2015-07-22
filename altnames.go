package fias

import (
	"log"
	"errors"
)

type AltNamesRecod struct {
	OLDCODE   string
	NEWCODE   string
	LEVEL  string
}

func (rec *AltNamesRecod) From(strs []string) error {
	if len(strs) != 3 {
		return errors.New("Некорректное количество полей в AltNamesRecod")
	}

	rec.OLDCODE = strs[0]
	rec.NEWCODE = strs[1]
	rec.LEVEL = strs[2]

	return nil
}

func (rec AltNamesRecod) Print() {
	log.Print("OLDCODE: '", rec.OLDCODE, "', ")
	log.Print("NEWCODE: '", rec.NEWCODE, "', ")
	log.Println("LEVEL: '", rec.LEVEL, "'")
}

func ImportAltNames(file string) ([]AltNamesRecod, error) {
	records, err := getRecordsFromDBF(file)
	if err != nil {
		log.Println("Не могу прочитать записи:", err)
		return nil, err
	}

	var res []AltNamesRecod
	for _, record := range records {
		var rec AltNamesRecod
		err := rec.From(record)
		if err != nil {
			log.Println("Не могу получить AltNamesRecod из массива строк:", err)
			return nil, err
		}
		res = append(res, rec)
	}

	return res, err
}
