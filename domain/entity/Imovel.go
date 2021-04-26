package entity

import "time"

type Imovel struct {
	Usableareas   int          `json:"usableAreas"`
	Listingtype   string       `json:"listingType"`
	Createdat     time.Time    `json:"createdAt"`
	Listingstatus string       `json:"listingStatus"`
	ID            string       `json:"id"`
	Parkingspaces int          `json:"parkingSpaces"`
	Updatedat     time.Time    `json:"updatedAt"`
	Owner         bool         `json:"owner"`
	Images        []string     `json:"images"`
	Address       Address      `json:"address"`
	Bathrooms     int          `json:"bathrooms"`
	Bedrooms      int          `json:"bedrooms"`
	Pricinginfos  PricingInfos `json:"pricingInfos"`
}
