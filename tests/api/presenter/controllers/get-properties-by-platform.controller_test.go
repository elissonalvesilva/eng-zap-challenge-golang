package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	usecases "github.com/elissonalvesilva/eng-zap-challenge-golang/api/application/use-cases"
	database "github.com/elissonalvesilva/eng-zap-challenge-golang/api/infra/db/in-memory"
	controllers "github.com/elissonalvesilva/eng-zap-challenge-golang/api/presenter/controllers"
	"github.com/elissonalvesilva/eng-zap-challenge-golang/domain/protocols"
	mockImovel "github.com/elissonalvesilva/eng-zap-challenge-golang/tests/mock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type GetPropertiesByPlatformTypeMockSuccess struct {
	mock.Mock
}

type GetPropertiesByPlatformTypeMockNotFound struct {
	mock.Mock
}

func (m *GetPropertiesByPlatformTypeMockSuccess) GetPropertiesByPlatformType(platform string, page int) (protocols.ReturnPlatformResult, protocols.ErrorResponse) {

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

func (m *GetPropertiesByPlatformTypeMockNotFound) GetPropertiesByPlatformType(platform string, page int) (protocols.ReturnPlatformResult, protocols.ErrorResponse) {

	m.Called(platform, page)
	return protocols.ReturnPlatformResult{}, protocols.ErrorResponse{
		Message: "Error",
	}
}

type ControllerStub struct {
	sut                         *controllers.GetPropertiesByPlatformHandler
	getPropertiesByPlatformStub *usecases.GetPropertiesByPlatform
}

func MakeDBGetProperties() *usecases.GetPropertiesByPlatform {
	imoveis := protocols.PlatformType{
		Zap:      []protocols.Imovel{},
		VivaReal: []protocols.Imovel{},
	}
	db := database.NewDatabasePlatformLocalStorageRepository(imoveis)
	getPropertiesUseCase := usecases.NewGetPropertiesByPlatformHandler(db)
	return getPropertiesUseCase
}

func makeSut() *ControllerStub {
	getPropertiesByPlatformStub := MakeDBGetProperties()
	sut := controllers.NewGetPropertiesByPlatformHandler(getPropertiesByPlatformStub)

	return &ControllerStub{sut, getPropertiesByPlatformStub}
}

func TestGetPropertiesByPlatformPlatformParam(t *testing.T) {
	t.Logf("Should return bad request if platform is not provided")

	r, _ := http.NewRequest("GET", "/search/{platform}", nil)
	w := httptest.NewRecorder()

	makeSut().sut.GetPropertiesByPlatform(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetPropertiesByPlatformPageParam(t *testing.T) {
	t.Logf("Should return internal server error if page is invalid")

	r, _ := http.NewRequest("GET", "/search/{platform}", nil)
	w := httptest.NewRecorder()

	platformParam := map[string]string{
		"platform": "zap",
	}

	r = mux.SetURLVars(r, platformParam)
	params := r.URL.Query()
	params.Set("page", "aaa")
	r.URL.RawQuery = params.Encode()

	makeSut().sut.GetPropertiesByPlatform(w, r)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestGetPropertiesByPlatformNotFoundPlatform(t *testing.T) {
	t.Logf("Should return not found if platform is not in database")

	r, _ := http.NewRequest("GET", "/search/{platform}", nil)
	w := httptest.NewRecorder()

	platformParam := map[string]string{
		"platform": "test_platform",
	}

	r = mux.SetURLVars(r, platformParam)
	params := r.URL.Query()
	params.Set("page", "1")
	r.URL.RawQuery = params.Encode()

	testObj := new(GetPropertiesByPlatformTypeMockNotFound)

	testObj.On("GetPropertiesByPlatformType", "test_platform", 1).Return(protocols.ReturnPlatformResult{}, protocols.ErrorResponse{})

	sut := controllers.NewGetPropertiesByPlatformHandler(testObj)

	sut.GetPropertiesByPlatform(w, r)
	testObj.AssertExpectations(t)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestGetPropertiesByPlatformSuccessToGetPlatform(t *testing.T) {
	t.Logf("Should return success to get platform")

	r, _ := http.NewRequest("GET", "/search/{platform}", nil)
	w := httptest.NewRecorder()

	platformParam := map[string]string{
		"platform": "zap",
	}

	r = mux.SetURLVars(r, platformParam)
	params := r.URL.Query()
	params.Set("page", "1")
	r.URL.RawQuery = params.Encode()

	testObj := new(GetPropertiesByPlatformTypeMockSuccess)

	testObj.On("GetPropertiesByPlatformType", "zap", 1)

	sut := controllers.NewGetPropertiesByPlatformHandler(testObj)

	sut.GetPropertiesByPlatform(w, r)
	testObj.AssertExpectations(t)
	assert.Equal(t, http.StatusOK, w.Code)
}
