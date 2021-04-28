package vivamodel

import (
	"errors"

	"github.com/elissonalvesilva/eng-zap-challenge-golang/domain/protocols"
	boudingbox "github.com/elissonalvesilva/eng-zap-challenge-golang/shared/bounding-box"
	consts "github.com/elissonalvesilva/eng-zap-challenge-golang/utils"
)

func NewImovel(imovel protocols.Imovel) (protocols.Imovel, error) {
	rentalTotalPrice := imovel.Pricinginfos.RentalTotalPrice
	if rentalTotalPrice <= 10 {
		return imovel, errors.New("Invalid Price Value")
	}

	parsedMontlyCondoFee := imovel.Pricinginfos.Monthlycondofee
	if parsedMontlyCondoFee == 0 {
		return imovel, errors.New("Invalid Monthlycondofee value")
	}

	rentalValueAdd30Percent := rentalTotalPrice + (rentalTotalPrice * 0.3)
	if parsedMontlyCondoFee >= rentalValueAdd30Percent {
		return imovel, errors.New("Not Elegible Monthlycondofee")
	}

	if boudingbox.IsInBoundingBoxZap(imovel) {
		maxValueRental := consts.VIVAREAL_RENTAL_MAX_VALUE + (consts.VIVAREAL_RENTAL_MAX_VALUE * 0.5)
		if maxValueRental < rentalTotalPrice {
			return imovel, errors.New("Not Elegible max rental rental")
		}
	}
	return imovel, nil
}
