package inmemory

import (
	"errors"
	"testing"

	database "github.com/elissonalvesilva/eng-zap-challenge-golang/api/infra/db/in-memory"
	"github.com/elissonalvesilva/eng-zap-challenge-golang/domain/protocols"
	mockImovel "github.com/elissonalvesilva/eng-zap-challenge-golang/tests/mock"
	"github.com/stretchr/testify/assert"
)

var imoveisData = protocols.PlatformType{
	Zap: []protocols.Imovel{
		mockImovel.MockImovelSuccess,
	},
	VivaReal: []protocols.Imovel{
		mockImovel.MockImovelSuccess,
	},
}

type RequestParams struct {
	Platform string
	Page     int
}

func TestDatabasePlatformLocalStorageRepository_GetPropertiesZap(t *testing.T) {
	t.Log("Should return response with Zap data")
	sut := database.NewDatabasePlatformLocalStorageRepository(imoveisData)
	requestParams := RequestParams{
		Platform: "zap",
		Page:     1,
	}

	expectedResponse := protocols.ReturnPlatformResult{
		PageNumber: requestParams.Page,
		PageSize:   10,
		TotalCount: 1,
		Listings: []protocols.Imovel{
			imoveisData.Zap[0],
		},
	}

	response, _ := sut.GetProperties(requestParams.Platform, requestParams.Page)
	assert.Equal(t, expectedResponse, response)
}

func TestDatabasePlatformLocalStorageRepository_GetPropertiesViva(t *testing.T) {
	t.Log("Should return response with VivaReal data")
	sut := database.NewDatabasePlatformLocalStorageRepository(imoveisData)
	requestParams := RequestParams{
		Platform: "vivareal",
		Page:     1,
	}

	expectedResponse := protocols.ReturnPlatformResult{
		PageNumber: requestParams.Page,
		PageSize:   10,
		TotalCount: 1,
		Listings: []protocols.Imovel{
			imoveisData.VivaReal[0],
		},
	}

	response, _ := sut.GetProperties(requestParams.Platform, requestParams.Page)
	assert.Equal(t, expectedResponse, response)
}

func TestDatabasePlatformLocalStorageRepository_GetPropertiesError(t *testing.T) {
	t.Log("Should return a error if platform is not found")
	sut := database.NewDatabasePlatformLocalStorageRepository(imoveisData)
	requestParams := RequestParams{
		Platform: "not_found",
		Page:     1,
	}

	expectedResponse := protocols.ReturnPlatformResult{}

	expectedErrorResponse := protocols.ErrorResponse{
		Message: errors.New("Not found platform").Error(),
	}

	response, err := sut.GetProperties(requestParams.Platform, requestParams.Page)
	assert.Equal(t, expectedResponse, response)
	assert.Equal(t, err, expectedErrorResponse)
}
