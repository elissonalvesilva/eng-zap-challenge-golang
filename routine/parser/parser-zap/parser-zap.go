package parserzap

import (
	"errors"
	"strconv"

	"github.com/elissonalvesilva/eng-zap-challenge-golang/domain/entity"

	consts "github.com/elissonalvesilva/eng-zap-challenge-golang/utils"
)

type Imovel entity.Imovel

func ParserZap(imovel entity.Imovel) (entity.Imovel, error) {
	if imovel.Usableareas == 0 {
		return imovel, errors.New("Invalid Usableareas")
	}

	parsedPrice, err := strconv.ParseFloat(imovel.Pricinginfos.Price, 64)
	if err != nil {
		return imovel, errors.New("Invalid Price value")
	}

	valuePerMeter := parsedPrice / float64(imovel.Usableareas)
	if valuePerMeter <= consts.ZAP_SALE_MIN_VALUE_BY_METER {
		return imovel, errors.New("Invalid Value by Meter")
	}

	if isInBoundingBoxZap(imovel) {
		minValueSale := consts.ZAP_SALE_MIN_VALUE - (consts.ZAP_SALE_MIN_VALUE * 0.1)
		if minValueSale > parsedPrice {
			return imovel, errors.New("Invalid min sale")
		}
	}
	return imovel, nil
}

func isInBoundingBoxZap(data entity.Imovel) bool {

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
