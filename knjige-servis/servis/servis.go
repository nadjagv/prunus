package servis

import (
	model "knjige-servis/model"
	repozitorijum "knjige-servis/repozitorijum"
)

func PreuzmiSve() []model.Knjiga {
	return repozitorijum.PreuzmiSve()
}
