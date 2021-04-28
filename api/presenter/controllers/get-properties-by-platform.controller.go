package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	usecases "github.com/elissonalvesilva/eng-zap-challenge-golang/domain/use-cases"
	timetrack "github.com/elissonalvesilva/eng-zap-challenge-golang/shared/time-track"
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
	defer timetrack.TimeTrack(time.Now(), r.RequestURI+" Finished in ")

	platform := mux.Vars(r)
	var page int = 1
	queryPage := r.URL.Query().Get("page")
	if queryPage != "" {
		convertedPage, err := strconv.ParseInt(queryPage, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		page = int(convertedPage)
	}

	if platform["platform"] == "" {
		json.NewEncoder(w).Encode("'platform' param must be pass")
	}

	response, errorResponse := h.useCase.GetPropertiesByPlatformType(platform["platform"], page)
	if errorResponse != nil {
		http.Error(w, errorResponse.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(response)
}
