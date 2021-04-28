package usecases

import "github.com/elissonalvesilva/eng-zap-challenge-golang/domain/protocols"

type GetPropertiesByPlatform interface {
	GetPropertiesByPlatformType(platform string) (protocols.ReturnPlatformResult, error)
}
