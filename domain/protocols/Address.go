package protocols

type Address struct {
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Geolocation  struct {
		Precision string `json:"precision"`
		Location  struct {
			Lon float64 `json:"lon"`
			Lat float64 `json:"lat"`
		} `json:"location"`
	} `json:"geoLocation"`
}
