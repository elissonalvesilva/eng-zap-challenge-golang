package vivamodel

import (
	"github.com/elissonalvesilva/eng-zap-challenge-golang/domain/protocols"
	boudingbox "github.com/elissonalvesilva/eng-zap-challenge-golang/shared/bounding-box"
	modelError "github.com/elissonalvesilva/eng-zap-challenge-golang/shared/errors"

	consts "github.com/elissonalvesilva/eng-zap-challenge-golang/utils"
)

func NewImovel(imovel protocols.Imovel) (protocols.Imovel, error) {
	rentalTotalPrice := imovel.Pricinginfos.RentalTotalPrice
	if rentalTotalPrice <= 10 {
		return imovel, modelError.InvalidPriceValue(imovel.ID)
	}

	parsedMontlyCondoFee := imovel.Pricinginfos.Monthlycondofee
	if parsedMontlyCondoFee == 0 {
		return imovel, modelError.InvalidMonthlycondofeeValue(imovel.ID)
	}

	rentalValueAdd30Percent := rentalTotalPrice + (rentalTotalPrice * 0.3)
	if parsedMontlyCondoFee >= rentalValueAdd30Percent {
		return imovel, modelError.NotElegibleMonthlycondofeeValue(imovel.ID)
	}

	if boudingbox.IsInBoundingBoxZap(imovel) {
		maxValueRental := consts.VIVAREAL_RENTAL_MAX_VALUE + (consts.VIVAREAL_RENTAL_MAX_VALUE * 0.5)
		if maxValueRental < rentalTotalPrice {
			return imovel, modelError.NotElegibleMaxRentalValue(imovel.ID)
		}
	}
	return imovel, nil
}
