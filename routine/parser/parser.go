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

func readCatalog() []Imovel {
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
	return imoveis
}

func Run() {
	// var parsedZapImoveis []Imovel
	// var parsedVivaImoveis []Imovel
	imoveis := readCatalog()

	zap := make(chan Imovel)
	viva := make(chan Imovel)
	err := make(chan error)
	quit := make(chan bool)

	go parser(imoveis, zap, viva, quit, err)
	receive(zap, viva, quit, err)
	// fmt.Println(len(parsedZapImoveis), len(parsedVivaImoveis))
}

func receive(zap chan Imovel, viva chan Imovel, quit chan bool, err chan error) {
	var parsedZapImoveis []Imovel
	var parsedVivaImoveis []Imovel
	var errors []error
	for {
		select {
		case v := <-zap:
			parsedZapImoveis = append(parsedZapImoveis, v)
		case v := <-viva:
			parsedVivaImoveis = append(parsedVivaImoveis, v)
		case v := <-err:
			errors = append(errors, v)
		case v, ok := <-quit:
			if !ok {
				fmt.Println("Deu zebra: ", v)
			} else {
				fmt.Println("Encerrando. Recebemos: ", v)
				fmt.Println(len(parsedZapImoveis), len(parsedVivaImoveis))
			}
			return
		}
	}

}

func parser(imoveis []Imovel, zap chan Imovel, viva chan Imovel, quit chan bool, err chan error) {
	for _, imovel := range imoveis {
		if validateLongAndLat(imovel) {
			err <- errors.New("EERROR")
			continue
		}

		if imovel.Pricinginfos.Businesstype == consts.SALE {
			parsedImovel, errr := parserToZap(imovel)

			if errr != nil {
				err <- errr
				continue

			}

			zap <- parsedImovel
			continue
		}

		if imovel.Pricinginfos.Businesstype == consts.RENTAL {
			parsedImovel, errr := parseToVivaReal(imovel)
			if errr != nil {
				err <- errr
				continue

			}
			viva <- parsedImovel
			continue
		}
	}
	close(zap)
	close(viva)
	close(err)
	quit <- true
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
