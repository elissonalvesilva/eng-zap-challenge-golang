package protocols

import "github.com/elissonalvesilva/eng-zap-challenge-golang/domain/protocols"

type GetPropertiesByPlatformRepository interface {
	GetProperties(platform string) (protocols.ReturnPlatformResult, error)
}
