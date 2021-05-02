package mock

import (
	"time"

	"github.com/elissonalvesilva/eng-zap-challenge-golang/domain/protocols"
)

var MockImovelErrorUsableareas = protocols.Imovel{
	Usableareas:   0,
	Listingtype:   "list1",
	Createdat:     time.Now(),
	Listingstatus: "active",
	ID:            "123",
	Parkingspaces: 1,
	Updatedat:     time.Now(),
	Owner:         false,
	Images:        []string{"image1"},
	Address: protocols.Address{
		City:         "city",
		Neighborhood: "aa",
		Geolocation: protocols.Geolocation{
			Precision: "GEO",
			Location: protocols.Location{
				Lat: 1231313,
				Lon: 1213131,
			},
		},
	},
	Bathrooms: 1,
	Bedrooms:  1,
	Pricinginfos: protocols.PricingInfos{
		Price:            45000,
		Yearlyiptu:       11233,
		RentalTotalPrice: 12311313,
		Businesstype:     "SALE",
		Monthlycondofee:  123131,
	},
}

var MockImovelErrorInvalidPrice = protocols.Imovel{
	Usableareas:   40,
	Listingtype:   "list1",
	Createdat:     time.Now(),
	Listingstatus: "active",
	ID:            "123",
	Parkingspaces: 1,
	Updatedat:     time.Now(),
	Owner:         false,
	Images:        []string{"image1"},
	Address: protocols.Address{
		City:         "city",
		Neighborhood: "aa",
		Geolocation: protocols.Geolocation{
			Precision: "GEO",
			Location: protocols.Location{
				Lat: 1231313,
				Lon: 1213131,
			},
		},
	},
	Bathrooms: 1,
	Bedrooms:  1,
	Pricinginfos: protocols.PricingInfos{
		Price:            0,
		Yearlyiptu:       11233,
		RentalTotalPrice: 12311313,
		Businesstype:     "SALE",
		Monthlycondofee:  123131,
	},
}

var MockImovelErrorNotElegiblePricePerMeterValue = protocols.Imovel{
	Usableareas:   150,
	Listingtype:   "list1",
	Createdat:     time.Now(),
	Listingstatus: "active",
	ID:            "123",
	Parkingspaces: 1,
	Updatedat:     time.Now(),
	Owner:         false,
	Images:        []string{"image1"},
	Address: protocols.Address{
		City:         "city",
		Neighborhood: "aa",
		Geolocation: protocols.Geolocation{
			Precision: "GEO",
			Location: protocols.Location{
				Lat: 1231313,
				Lon: 1213131,
			},
		},
	},
	Bathrooms: 1,
	Bedrooms:  1,
	Pricinginfos: protocols.PricingInfos{
		Price:            45000,
		Yearlyiptu:       11233,
		RentalTotalPrice: 12311313,
		Businesstype:     "SALE",
		Monthlycondofee:  123131,
	},
}

var MockImovelErrorNotElegibleMinSaleValue = protocols.Imovel{
	Usableareas:   150,
	Listingtype:   "list1",
	Createdat:     time.Now(),
	Listingstatus: "active",
	ID:            "123",
	Parkingspaces: 1,
	Updatedat:     time.Now(),
	Owner:         false,
	Images:        []string{"image1"},
	Address: protocols.Address{
		City:         "city",
		Neighborhood: "aa",
		Geolocation: protocols.Geolocation{
			Precision: "GEO",
			Location: protocols.Location{
				Lat: -23.553723,
				Lon: -46.631146,
			},
		},
	},
	Bathrooms: 1,
	Bedrooms:  1,
	Pricinginfos: protocols.PricingInfos{
		Price:            45000,
		Yearlyiptu:       11233,
		RentalTotalPrice: 12311313,
		Businesstype:     "SALE",
		Monthlycondofee:  123131,
	},
}

var MockImovelSuccess = protocols.Imovel{
	Usableareas:   4,
	Listingtype:   "list1",
	Createdat:     time.Now(),
	Listingstatus: "active",
	ID:            "123",
	Parkingspaces: 1,
	Updatedat:     time.Now(),
	Owner:         false,
	Images:        []string{"image1"},
	Address: protocols.Address{
		City:         "city",
		Neighborhood: "aa",
		Geolocation: protocols.Geolocation{
			Precision: "GEO",
			Location: protocols.Location{
				Lat: -23.553723,
				Lon: -46.631146,
			},
		},
	},
	Bathrooms: 1,
	Bedrooms:  1,
	Pricinginfos: protocols.PricingInfos{
		Price:            45000,
		Yearlyiptu:       11233,
		RentalTotalPrice: 12311313,
		Businesstype:     "SALE",
		Monthlycondofee:  123131,
	},
}
