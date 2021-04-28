package protocols

import "github.com/elissonalvesilva/eng-zap-challenge-golang/domain/entity"

type ReturnPlatformResult struct {
	PageNumber int             `json:"pageNumber"`
	PageSize   int             `json:"pageSize"`
	TotalCount int             `json:"totalCount"`
	Listings   []entity.Imovel `json:"listings"`
}
