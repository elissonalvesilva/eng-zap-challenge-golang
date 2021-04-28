package protocols

type ReturnPlatformResult struct {
	PageNumber int      `json:"pageNumber"`
	PageSize   int      `json:"pageSize"`
	TotalCount int      `json:"totalCount"`
	Listings   []Imovel `json:"listings"`
}
