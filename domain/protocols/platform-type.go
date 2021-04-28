package protocols

import "github.com/elissonalvesilva/eng-zap-challenge-golang/domain/entity"

type PlatformType struct {
	Zap      []entity.Imovel `json:"zap"`
	VivaReal []entity.Imovel `json:"vivareal"`
}
