package repozitorijum

import (
	model "knjige-servis/model"
	util "knjige-servis/util"
)

func PreuzmiSve() []model.Knjiga {
	var knjige []model.Knjiga
	util.Database.Find(&knjige)
	return knjige

}
