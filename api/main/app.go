package mainly

import (
	"fmt"
	"net/http"

	controllers "github.com/elissonalvesilva/eng-zap-challenge-golang/api/main/factories/controllers"

	"github.com/gorilla/mux"
)

type App struct {
	httpServer *http.Server
}

func NewApp(port int) *App {
	router := mux.NewRouter()

	controller := controllers.MakeGetPropertiesByPlatformController()
	router.HandleFunc("/search/{platform}", controller.GetPropertiesByPlatformController).Methods("GET")

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
