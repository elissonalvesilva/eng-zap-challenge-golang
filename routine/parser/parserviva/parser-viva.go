package parserviva

import (
	"errors"
	"strconv"

	"github.com/elissonalvesilva/eng-zap-challenge-golang/domain/entity"

	consts "github.com/elissonalvesilva/eng-zap-challenge-golang/utils"
)

type Imovel entity.Imovel

func ParseToVivaReal(imovel Imovel) (Imovel, error) {
	parsedRentalTotalPrice, err := strconv.ParseFloat(imovel.Pricinginfos.RentalTotalPrice, 64)
	if err != nil {
		return imovel, errors.New("Invalid Price Value")
	}

	parsedMontlyCondoFee, err := strconv.ParseFloat(imovel.Pricinginfos.Monthlycondofee, 64)
	if err != nil {
		return imovel, errors.New("Invalid Monthlycondofee value")
	}

	rentalValueAdd30Percent := parsedRentalTotalPrice + (parsedRentalTotalPrice * 0.3)
	if parsedMontlyCondoFee >= rentalValueAdd30Percent {
		return imovel, errors.New("Not Elegible Monthlycondofee")
	}

	if isInBoundingBoxZap(imovel) {
		maxValueRental := consts.VIVAREAL_RENTAL_MAX_VALUE + (consts.VIVAREAL_RENTAL_MAX_VALUE * 0.5)
		if maxValueRental < parsedRentalTotalPrice {
			return imovel, errors.New("Not Elegible max rental rental")
		}
	}
	return imovel, nil
}

func isInBoundingBoxZap(data Imovel) bool {

	minlon := -46.693419 // left
	minlat := -23.568704 // top
	maxlon := -46.641146 // right
	maxlat := -23.546686 // botom
	point := data.Address.Geolocation.Location
	/* Check latitude bounds first. */
	if minlat >= point.Lat && point.Lat >= maxlat {
		if minlon <= maxlon && minlon <= point.Lon && point.Lon <= maxlon {
			return true
		} else if minlon > maxlon && (minlon <= point.Lon || point.Lon <= maxlon) {
			return true
		}
	}

	return false
}
