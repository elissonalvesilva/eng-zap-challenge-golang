package parser

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/elissonalvesilva/eng-zap-challenge-golang/domain/entity"
	consts "github.com/elissonalvesilva/eng-zap-challenge-golang/utils"
)

type Imovel entity.Imovel
type PlatformType struct {
	Zap      []Imovel `json:"zap"`
	VivaReal []Imovel `json:"vivareal"`
}

func Run() {
	var parsedZapImoveis []Imovel
	var parsedVivaImoveis []Imovel

	path_catalog := os.Getenv("PATH_DADOS") + os.Getenv("FILENAME_CATALOG")
	jsonFile, err := os.Open(path_catalog)

	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	data, _ := ioutil.ReadAll(jsonFile)
	var imoveis []Imovel
	if err := json.Unmarshal(data, &imoveis); err != nil {
		panic(err)
	}

	for _, imovel := range imoveis {
		parsedImovel, platformType, err := parser(imovel)

		if err != nil {
			fmt.Println(imovel.ID, err)
			continue
		}

		if platformType == "zap" {
			parsedZapImoveis = append(parsedZapImoveis, parsedImovel)
		} else {
			parsedVivaImoveis = append(parsedVivaImoveis, parsedImovel)
		}

	}

	PlatformTypeStruct := PlatformType{
		Zap:      parsedZapImoveis,
		VivaReal: parsedVivaImoveis,
	}

	f, err := os.OpenFile(os.Getenv("FILENAME_PARSED_CATALOG"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	bytes, err := json.Marshal(PlatformTypeStruct)
	if err != nil {
		fmt.Println(err)
		return
	}

	if _, err := f.Write(bytes); err != nil {
		fmt.Println(err)
	}

	fmt.Println(len(parsedZapImoveis), len(parsedVivaImoveis))
}

func parser(imovel Imovel) (Imovel, string, error) {
	if validateLongAndLat(imovel) {
		return imovel, "", errors.New("Invalid imovel")
	}

	if imovel.Pricinginfos.Businesstype == consts.SALE {
		parsedImovel, err := parserToZap(imovel)
		if err != nil {
			return parsedImovel, "", err
		}
		return parsedImovel, "zap", nil
	}

	if imovel.Pricinginfos.Businesstype == consts.RENTAL {
		parsedImovel, err := parseToVivaReal(imovel)
		if err != nil {
			return parsedImovel, "", err
		}
		return parsedImovel, "viva", nil
	}
	return imovel, "", nil
}

func parserToZap(imovel Imovel) (Imovel, error) {
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

func parseToVivaReal(imovel Imovel) (Imovel, error) {
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

func validateLongAndLat(imovel Imovel) bool {
	isEmptyLatAndLong := false
	if imovel.Address.Geolocation.Location.Lat == 0 &&
		imovel.Address.Geolocation.Location.Lon == 0 {
		isEmptyLatAndLong = true
	}
	return isEmptyLatAndLong
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
