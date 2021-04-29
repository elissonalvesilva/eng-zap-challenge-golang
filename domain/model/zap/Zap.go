package zapmodel

import (
	modelError "github.com/elissonalvesilva/eng-zap-challenge-golang/shared/errors"

	"github.com/elissonalvesilva/eng-zap-challenge-golang/domain/protocols"
	boudingbox "github.com/elissonalvesilva/eng-zap-challenge-golang/shared/bounding-box"
	consts "github.com/elissonalvesilva/eng-zap-challenge-golang/utils"
)

func NewImovel(imovel protocols.Imovel) (protocols.Imovel, error) {
	if imovel.Usableareas == 0 {
		return imovel, modelError.InvalidUsableareasValue(imovel.ID)
	}

	price := imovel.Pricinginfos.Price
	if price <= 1 {
		return imovel, modelError.InvalidPriceValue(imovel.ID)
	}

	valuePerMeter := price / float64(imovel.Usableareas)
	if valuePerMeter <= consts.ZAP_SALE_MIN_VALUE_BY_METER {
		return imovel, modelError.NotElegiblePricePerMeterValue(imovel.ID)
	}

	if boudingbox.IsInBoundingBoxZap(imovel) {
		minValueSale := consts.ZAP_SALE_MIN_VALUE - (consts.ZAP_SALE_MIN_VALUE * 0.1)
		if minValueSale > price {
			return imovel, modelError.NotElegibleMinSaleValue(imovel.ID)
		}
	}
	return imovel, nil
}
