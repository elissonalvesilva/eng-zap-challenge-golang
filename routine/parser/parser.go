package parser

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	errorsType "github.com/elissonalvesilva/eng-zap-challenge-golang/routine/errors"

	viva "github.com/elissonalvesilva/eng-zap-challenge-golang/domain/model/vivareal"
	zap "github.com/elissonalvesilva/eng-zap-challenge-golang/domain/model/zap"
	"github.com/elissonalvesilva/eng-zap-challenge-golang/domain/protocols"

	file "github.com/elissonalvesilva/eng-zap-challenge-golang/shared/file-json"
	elapsed "github.com/elissonalvesilva/eng-zap-challenge-golang/shared/time-track"

	consts "github.com/elissonalvesilva/eng-zap-challenge-golang/utils"
)

type Response struct {
	Type        string
	imovel      protocols.Imovel
	parsedError error
}

func removeLockRoutine() {
	filename := os.Getenv("PATH_DADOS") + "lock"
	e := os.Remove(filename)
	if e != nil {
		log.Fatal(e)
	}
}

func Run() {
	var parsedZapImoveis []protocols.Imovel
	var parsedVivaImoveis []protocols.Imovel
	var channel = make(chan Response)
	var wg sync.WaitGroup
	path_catalog := os.Getenv("PATH_DADOS") + os.Getenv("FILENAME_CATALOG")
	data, err := file.Read(path_catalog)

	if err != nil {
		panic(err)
	}

	var imoveis []protocols.Imovel
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

	PlatformTypeStruct := protocols.PlatformType{
		Zap:      parsedZapImoveis,
		VivaReal: parsedVivaImoveis,
	}

	errToWriteFile := file.Write(os.Getenv("PATH_DADOS")+os.Getenv("FILENAME_PARSED_CATALOG"), PlatformTypeStruct)
	if errToWriteFile != nil {
		panic(errToWriteFile)
	}

	defer elapsed.TimeTrack(time.Now(), "parser")
	defer removeLockRoutine()
	fmt.Println("ZAP: ", len(parsedZapImoveis), "Viva:", len(parsedVivaImoveis))
}

func parser(imovel protocols.Imovel, wg *sync.WaitGroup, channel chan Response) {
	defer wg.Done()
	if validateLongAndLat(imovel) {
		channel <- Response{Type: "Error", imovel: imovel, parsedError: errorsType.InvalidLonAndLat(imovel.ID)}
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

func validateLongAndLat(imovel protocols.Imovel) bool {
	isEmptyLatAndLong := false
	if imovel.Address.Geolocation.Location.Lat == 0 &&
		imovel.Address.Geolocation.Location.Lon == 0 {
		isEmptyLatAndLong = true
	}
	return isEmptyLatAndLong
}
