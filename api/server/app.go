package server

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/elissonalvesilva/eng-zap-challenge-golang/domain/protocols"
	file "github.com/elissonalvesilva/eng-zap-challenge-golang/shared/file-json"

	usecases "github.com/elissonalvesilva/eng-zap-challenge-golang/api/application/use-cases"
	controllers "github.com/elissonalvesilva/eng-zap-challenge-golang/api/presenter/controllers"

	database "github.com/elissonalvesilva/eng-zap-challenge-golang/api/infra/db/in-memory"

	"github.com/gorilla/mux"
)

type App struct {
	httpServer *http.Server

	handleController func(http.ResponseWriter, *http.Request)
}

func NewApp() *App {
	path_catalog := os.Getenv("PATH_DADOS") + os.Getenv("FILENAME_PARSED_CATALOG")
	data, err := file.Read(path_catalog)

	if err != nil {
		panic(err)
	}

	var imoveis protocols.PlatformType
	if err := json.Unmarshal(data, &imoveis); err != nil {
		panic(err)
	}

	db := database.NewDatabasePlatformLocalStorageRepository(imoveis)
	getPropertiesUseCase := usecases.NewGetPropertiesByPlatformHandler(db)
	controller := controllers.NewGetPropertiesByPlatformHandler(getPropertiesUseCase)

	return &App{
		handleController: controller.GetPropertiesByPlatform,
	}
}

func (a *App) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/search/{platform}", a.handleController).Methods("GET")
	log.Fatal(http.ListenAndServe(":4513", router))

}
