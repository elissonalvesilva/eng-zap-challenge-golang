package entity

type PricingInfos struct {
	Yearlyiptu       float64 `json:"yearlyIptu,string"`
	Price            float64 `json:"price,string"`
	RentalTotalPrice float64 `json:"rentalTotalPrice,string"`
	Businesstype     string  `json:"businessType"`
	Monthlycondofee  float64 `json:"monthlyCondoFee,string"`
}
