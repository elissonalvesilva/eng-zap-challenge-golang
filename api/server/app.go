package server

import (
	"log"
	"net/http"

	usecases "github.com/elissonalvesilva/eng-zap-challenge-golang/api/application/use-cases"
	controllers "github.com/elissonalvesilva/eng-zap-challenge-golang/api/presenter/controllers"

	database "github.com/elissonalvesilva/eng-zap-challenge-golang/api/infra/db/in-memory"

	"github.com/gorilla/mux"
	// controllers "github.com/elissonalvesilva/eng-zap-challenge-golang/api/presenter/controllers"
)

type App struct {
	httpServer *http.Server

	handleController func(http.ResponseWriter, *http.Request)
}

func NewApp() *App {

	db := database.NewDatabaseLocalStorage()
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
