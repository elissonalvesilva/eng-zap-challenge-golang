package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/elissonalvesilva/eng-zap-challenge-golang/domain/protocols"
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

	param := mux.Vars(r)
	var page int = 1
	queryPage := r.URL.Query().Get("page")

	if param["platform"] == "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(protocols.ErrorResponse{
			Message: "'platform' param must be pass",
		})
		return
	}

	if queryPage != "" && queryPage != "0" {
		convertedPage, err := strconv.ParseInt(queryPage, 10, 64)
		if err != nil {
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(protocols.ErrorResponse{
				Message: "Internal server error",
			})
			fmt.Println(err)
			return
		}
		page = int(convertedPage)
	}

	platform := strings.ToLower(param["platform"])

	response, errorResponse := h.useCase.GetPropertiesByPlatformType(platform, page)
	if errorResponse != (protocols.ErrorResponse{}) {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	json.NewEncoder(w).Encode(response)
}
