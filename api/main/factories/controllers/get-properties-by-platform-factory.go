package controllers

import (
	usecases "github.com/elissonalvesilva/eng-zap-challenge-golang/api/main/factories/application/use-cases"
	"github.com/elissonalvesilva/eng-zap-challenge-golang/api/presenter/controllers"
)

func MakeGetPropertiesByPlatformController() *controllers.GetPropertiesByPlatformHandler {
	getPropertiesUseCase := usecases.MakeDBGetProperties()
	controller := controllers.NewGetPropertiesByPlatformHandler(getPropertiesUseCase)
	return controller
}
