package parser

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
	"time"

	"github.com/elissonalvesilva/eng-zap-challenge-golang/domain/entity"
	viva "github.com/elissonalvesilva/eng-zap-challenge-golang/domain/model/vivareal"
	zap "github.com/elissonalvesilva/eng-zap-challenge-golang/domain/model/zap"

	elapsed "github.com/elissonalvesilva/eng-zap-challenge-golang/shared/time-track"
	consts "github.com/elissonalvesilva/eng-zap-challenge-golang/utils"
)

type PlatformType struct {
	Zap      []entity.Imovel `json:"zap"`
	VivaReal []entity.Imovel `json:"vivareal"`
}

type Response struct {
	Type        string
	imovel      entity.Imovel
	parsedError error
}

func Run() {
	var parsedZapImoveis []entity.Imovel
	var parsedVivaImoveis []entity.Imovel
	var channel = make(chan Response)
	var wg sync.WaitGroup
	path_catalog := os.Getenv("PATH_DADOS") + os.Getenv("FILENAME_CATALOG")
	jsonFile, err := os.Open(path_catalog)

	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	data, _ := ioutil.ReadAll(jsonFile)
	var imoveis []entity.Imovel
	if err := json.Unmarshal(data, &imoveis); err != nil {
		panic(err)
	}

	for _, imovel := range imoveis {
		wg.Add(1)
		go parser(imovel, &wg, channel)
	}
	go func() {
		wg.Wait()
		close(channel)
	}()

	for response := range channel {
		if response.Type == "zap" {
			parsedZapImoveis = append(parsedZapImoveis, response.imovel)
		} else if response.Type == "viva" {
			parsedVivaImoveis = append(parsedVivaImoveis, response.imovel)
		} else if response.Type == "Error" {
			fmt.Println(response.parsedError)
		}
	}

	PlatformTypeStruct := PlatformType{
		Zap:      parsedZapImoveis,
		VivaReal: parsedVivaImoveis,
	}

	f, err := os.OpenFile(os.Getenv("FILENAME_PARSED_CATALOG"), os.O_RDWR|os.O_CREATE, 0644)
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

	defer elapsed.TimeTrack(time.Now(), "parser")
	fmt.Println(len(parsedZapImoveis), len(parsedVivaImoveis))
}

func parser(imovel entity.Imovel, wg *sync.WaitGroup, channel chan Response) {
	defer wg.Done()
	if validateLongAndLat(imovel) {
		channel <- Response{Type: "Error", imovel: imovel, parsedError: errors.New("Invalid imovel")}
		return
	}

	if imovel.Pricinginfos.Businesstype == consts.SALE {
		parsedImovel, err := zap.NewImovel(imovel)
		if err != nil {
			channel <- Response{Type: "Error", imovel: imovel, parsedError: err}
			return
		}
		channel <- Response{Type: "zap", imovel: parsedImovel, parsedError: nil}
	}

	if imovel.Pricinginfos.Businesstype == consts.RENTAL {
		parsedImovel, err := viva.NewImovel(imovel)
		if err != nil {
			channel <- Response{Type: "Error", imovel: imovel, parsedError: err}
			return
		}
		channel <- Response{Type: "viva", imovel: parsedImovel, parsedError: nil}
	}

}

func validateLongAndLat(imovel entity.Imovel) bool {
	isEmptyLatAndLong := false
	if imovel.Address.Geolocation.Location.Lat == 0 &&
		imovel.Address.Geolocation.Location.Lon == 0 {
		isEmptyLatAndLong = true
	}
	return isEmptyLatAndLong
}
