package usecases

import (
	repository "github.com/elissonalvesilva/eng-zap-challenge-golang/api/application/protocols"
	"github.com/elissonalvesilva/eng-zap-challenge-golang/domain/entity"
)

type GetPropertiesByPlatform struct {
	repo repository.GetPropertiesByPlatformRepository
}

func NewGetPropertiesByPlatformHandler(propertiesRepository repository.GetPropertiesByPlatformRepository) *GetPropertiesByPlatform {
	return &GetPropertiesByPlatform{
		repo: propertiesRepository,
	}
}

func (r *GetPropertiesByPlatform) GetPropertiesByPlatformType(platform string) ([]entity.Imovel, error) {
	response, err := r.repo.GetProperties(platform)

	if err != nil {
		return nil, err
	}

	return response, nil
}
