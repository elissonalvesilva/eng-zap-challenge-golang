package controllers

import (
	"encoding/json"
	"fmt"
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
	fmt.Println(errorResponse)
	if errorResponse != nil {
		http.Error(w, errorResponse.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(response)
}
