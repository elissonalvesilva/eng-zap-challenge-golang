package server

import (
	"encoding/json"
	"fmt"
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
}

func NewApp(port int) *App {
	router := mux.NewRouter()

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

	router.HandleFunc("/search/{platform}", controller.GetPropertiesByPlatform).Methods("GET")

	return &App{
		httpServer: &http.Server{
			Addr:    fmt.Sprintf(":%v", port),
			Handler: router,
		},
	}
}

func (a *App) Run(application string) {
	fmt.Println(application+" Is running in port", a.httpServer.Addr)
	a.httpServer.ListenAndServe()
}
