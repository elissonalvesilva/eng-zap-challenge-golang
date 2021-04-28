package usecases

import "github.com/elissonalvesilva/eng-zap-challenge-golang/domain/entity"

type GetPropertiesByPlatform interface {
	GetPropertiesByPlatformType(platform string) ([]entity.Imovel, error)
}
