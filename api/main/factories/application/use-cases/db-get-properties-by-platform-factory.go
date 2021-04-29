package usecases

import (
	"encoding/json"
	"os"

	database "github.com/elissonalvesilva/eng-zap-challenge-golang/api/infra/db/in-memory"

	usecases "github.com/elissonalvesilva/eng-zap-challenge-golang/api/application/use-cases"
	"github.com/elissonalvesilva/eng-zap-challenge-golang/domain/protocols"
	"github.com/elissonalvesilva/eng-zap-challenge-golang/shared/file-json"
)

func getDataPlatform() protocols.PlatformType {
	path_catalog := os.Getenv("PATH_DADOS") + os.Getenv("FILENAME_PARSED_CATALOG")
	data, err := file.Read(path_catalog)

	if err != nil {
		panic(err)
	}

	var imoveis protocols.PlatformType
	if err := json.Unmarshal(data, &imoveis); err != nil {
		panic(err)
	}
	return imoveis
}

func MakeDBGetProperties() *usecases.GetPropertiesByPlatform {
	imoveis := getDataPlatform()
	db := database.NewDatabasePlatformLocalStorageRepository(imoveis)
	getPropertiesUseCase := usecases.NewGetPropertiesByPlatformHandler(db)
	return getPropertiesUseCase
}
