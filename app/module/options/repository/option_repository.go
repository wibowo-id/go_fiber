package repository

import (
	"go_fiber_wibowo/app/module/options/response"
	"go_fiber_wibowo/internal/bootstrap/database"
	"go_fiber_wibowo/utils/config"
	"strconv"
)

type optionRepository struct {
	DB  *database.Database
	Cfg *config.Config
}

//go:generate mockgen -destination=option_repository_mock.go -package=repository . OptionRepository
type OptionRepository interface {
	GetWilayah(kode string) (res []*response.OptionsResponse, err error)
	GetRole() (res []*response.OptionsResponse, err error)
	GetPermission() (res []*response.OptionsResponse, err error)
}

func NewOptionRepository(db *database.Database) OptionRepository {
	return &optionRepository{
		DB: db,
	}
}

func (_i *optionRepository) GetWilayah(kode string) (res []*response.OptionsResponse, err error) {
	length := len(kode)

	qState := "SELECT * FROM wilayahs "
	if kode == "" {
		qState = qState + "WHERE length(kode)=2"
	} else {
		qState = qState + "WHERE kode like '" + kode + "%' "

		notLength := 0
		if length < 13 {
			length = length + 3
			notLength = length + 3

			newNotLength := strconv.Itoa(notLength)
			qState = qState + "AND kode not like '%" + kode + "' "
			qState = qState + "AND length(kode) < '" + newNotLength + "' "
		}
	}

	err = _i.DB.DB.Raw(qState).Scan(&res).Error
	if err != nil {
		return nil, err
	}

	return
}

func (_i *optionRepository) GetRole() (res []*response.OptionsResponse, err error) {
	qState := "SELECT id as kode, role as nama FROM roles "
	qState = qState + "WHERE is_active = true"

	err = _i.DB.DB.Raw(qState).Scan(&res).Error
	if err != nil {
		return nil, err
	}

	return
}

func (_i *optionRepository) GetPermission() (res []*response.OptionsResponse, err error) {
	qState := "SELECT id as kode, judul as nama FROM permissions "
	qState = qState + "WHERE is_active = true"

	err = _i.DB.DB.Raw(qState).Scan(&res).Error
	if err != nil {
		return nil, err
	}

	return
}
