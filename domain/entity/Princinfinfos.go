package entity

type PricingInfos struct {
	Yearlyiptu       string `json:"yearlyIptu"`
	Price            string `json:"price"`
	RentalTotalPrice string `json:"rentalTotalPrice"`
	Businesstype     string `json:"businessType"`
	Monthlycondofee  string `json:"monthlyCondoFee"`
}
