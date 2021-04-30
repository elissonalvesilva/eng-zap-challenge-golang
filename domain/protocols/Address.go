package protocols

type Address struct {
	City         string      `json:"city"`
	Neighborhood string      `json:"neighborhood"`
	Geolocation  Geolocation `json:"geoLocation"`
}

type Geolocation struct {
	Precision string   `json:"precision"`
	Location  Location `json:"location"`
}

type Location struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}
