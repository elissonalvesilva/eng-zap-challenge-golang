package usecases

import (
	"testing"

	usecases "github.com/elissonalvesilva/eng-zap-challenge-golang/api/application/use-cases"
	"github.com/elissonalvesilva/eng-zap-challenge-golang/domain/protocols"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	mockImovel "github.com/elissonalvesilva/eng-zap-challenge-golang/__tests__/mock"
)

type GetPropertiesSuccess struct {
	mock.Mock
}

func (m *GetPropertiesSuccess) GetProperties(platform string, page int) (protocols.ReturnPlatformResult, protocols.ErrorResponse) {

	m.Called(platform, page)
	return protocols.ReturnPlatformResult{
		PageNumber: 1,
		PageSize:   10,
		TotalCount: 1,
		Listings: []protocols.Imovel{
			mockImovel.MockImovelSuccess,
		},
	}, protocols.ErrorResponse{}
}

type GetPropertiesError struct {
	mock.Mock
}

func (m *GetPropertiesError) GetProperties(platform string, page int) (protocols.ReturnPlatformResult, protocols.ErrorResponse) {

	m.Called(platform, page)
	return protocols.ReturnPlatformResult{}, protocols.ErrorResponse{
		Message: "Error to get platform",
	}
}

type RequestParams struct {
	Platform string
	Page     int
}

func TestGetPropertiesByPlatformGetPropertiesCorrectParams(t *testing.T) {
	t.Logf("Should return call GetProperties with correct params")
	requestParams := RequestParams{
		Platform: "zap",
		Page:     1,
	}

	testObj := new(GetPropertiesSuccess)

	testObj.On("GetProperties", requestParams.Platform, requestParams.Page)
	sut := usecases.NewGetPropertiesByPlatformHandler(testObj)
	sut.GetPropertiesByPlatformType(requestParams.Platform, requestParams.Page)

	testObj.AssertExpectations(t)
	testObj.AssertCalled(t, "GetProperties", requestParams.Platform, requestParams.Page)
}

func TestGetPropertiesByPlatformOnErrorGetProperties(t *testing.T) {
	t.Logf("Should return error when GetProperties return a error response")
	requestParams := RequestParams{
		Platform: "zap",
		Page:     1,
	}

	testObj := new(GetPropertiesError)

	testObj.On("GetProperties", requestParams.Platform, requestParams.Page)
	sut := usecases.NewGetPropertiesByPlatformHandler(testObj)
	_, err := sut.GetPropertiesByPlatformType(requestParams.Platform, requestParams.Page)

	testObj.AssertExpectations(t)
	testObj.AssertCalled(t, "GetProperties", requestParams.Platform, requestParams.Page)
	_, errorResponse := testObj.GetProperties(requestParams.Platform, requestParams.Page)
	assert.Equal(t, err, errorResponse)
}

func TestGetPropertiesByPlatformOnSuccessGetProperties(t *testing.T) {
	t.Logf("Should return success when GetProperties return a success response")
	requestParams := RequestParams{
		Platform: "zap",
		Page:     1,
	}

	testObj := new(GetPropertiesSuccess)

	testObj.On("GetProperties", requestParams.Platform, requestParams.Page)
	sut := usecases.NewGetPropertiesByPlatformHandler(testObj)
	getPropertiesByPlatformTypeResponse, err := sut.GetPropertiesByPlatformType(requestParams.Platform, requestParams.Page)

	testObj.AssertExpectations(t)
	testObj.AssertCalled(t, "GetProperties", requestParams.Platform, requestParams.Page)
	getPropertiesResponse, errorResponse := testObj.GetProperties(requestParams.Platform, requestParams.Page)
	assert.Equal(t, err, errorResponse)
	assert.Equal(t, getPropertiesByPlatformTypeResponse, getPropertiesResponse)

}
