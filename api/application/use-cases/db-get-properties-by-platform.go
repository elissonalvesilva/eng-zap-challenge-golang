package usecases

import (
	repository "github.com/elissonalvesilva/eng-zap-challenge-golang/api/application/protocols"
	"github.com/elissonalvesilva/eng-zap-challenge-golang/domain/protocols"
)

type GetPropertiesByPlatform struct {
	repo repository.GetPropertiesByPlatformRepository
}

func NewGetPropertiesByPlatformHandler(propertiesRepository repository.GetPropertiesByPlatformRepository) *GetPropertiesByPlatform {
	return &GetPropertiesByPlatform{
		repo: propertiesRepository,
	}
}

func (r *GetPropertiesByPlatform) GetPropertiesByPlatformType(platform string, page int) (protocols.ReturnPlatformResult, error) {
	response, err := r.repo.GetProperties(platform, page)

	if err != nil {
		return response, err
	}

	return response, nil
}
