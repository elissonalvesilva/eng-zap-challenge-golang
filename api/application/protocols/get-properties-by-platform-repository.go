package protocols

import "github.com/elissonalvesilva/eng-zap-challenge-golang/domain/entity"

type GetPropertiesByPlatformRepository interface {
	GetProperties(platform string) ([]entity.Imovel, error)
}
