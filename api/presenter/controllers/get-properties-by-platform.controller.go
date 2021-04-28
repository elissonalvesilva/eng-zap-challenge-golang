package controllers

import (
	"encoding/json"
	"net/http"

	usecases "github.com/elissonalvesilva/eng-zap-challenge-golang/domain/use-cases"
	"github.com/gorilla/mux"
)

type GetPropertiesByPlatformHandler struct {
	useCase usecases.GetPropertiesByPlatform
}

func NewGetPropertiesByPlatformHandler(useCase usecases.GetPropertiesByPlatform) *GetPropertiesByPlatformHandler {
	return &GetPropertiesByPlatformHandler{
		useCase: useCase,
	}
}

func (h *GetPropertiesByPlatformHandler) GetPropertiesByPlatform(w http.ResponseWriter, r *http.Request) {
	platform := mux.Vars(r)

	if platform["platform"] == "" {
		json.NewEncoder(w).Encode("'platform' param must be pass")
	}

	response, errorResponse := h.useCase.GetPropertiesByPlatformType(platform["platform"])
	if errorResponse != nil {
		json.NewEncoder(w).Encode(errorResponse)
	}

	json.NewEncoder(w).Encode(response)
}
