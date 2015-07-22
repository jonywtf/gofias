package fias

import (
	"errors"
)

type Kladr struct {
	Streets []StreetRecod
	Doma []DomRecod
	Kladr []KladrRecod
	Flat []FlatRecod
	SockBase []SockBaseRecod
	AltNames []AltNamesRecod
}

func Import(base_path string) (Kladr, error) {

	var kladr = Kladr{}
	var err error
	kladr.AltNames, err = ImportAltNames(base_path + "ALTNAMES.DBF")
	if err != nil {
		return kladr, errors.New("Не могу импортировать КЛАДР ALTNAMES.DBF:" + err.Error())
	}
	kladr.SockBase, err = ImportSockBase(base_path + "SOCRBASE.DBF")
	if err != nil {
		return kladr, errors.New("Не могу импортировать КЛАДР SOCRBASE.DBF:" + err.Error())
	}
	kladr.Flat, err = ImportFlat(base_path + "FLAT.DBF")
	if err != nil {
		return kladr, errors.New("Не могу импортировать КЛАДР FLAT.DBF:" + err.Error())
	}
	kladr.Kladr, err = ImportKladr(base_path + "KLADR.DBF")
	if err != nil {
		return kladr, errors.New("Не могу импортировать КЛАДР KLADR.DBF:" + err.Error())
	}
	kladr.Doma, err = ImportDoma(base_path + "DOMA.DBF")
	if err != nil {
		return kladr, errors.New("Не могу импортировать КЛАДР DOMA.DBF:" + err.Error())
	}
	kladr.Streets, err = ImportStreet(base_path + "STREET.DBF")
	if err != nil {
		return kladr, errors.New("Не могу импортировать КЛАДР STREET.DBF:" + err.Error())
	}
	return kladr, nil
}
