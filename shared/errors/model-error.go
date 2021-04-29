package error

import "errors"

func InvalidPriceValue(id string) error {
	return errors.New("[ERROR] - Invalid Price Value - ID: " + id)
}

func InvalidMonthlycondofeeValue(id string) error {
	return errors.New("[ERROR] - Invalid Monthlycondofee value - ID: " + id)
}

func InvalidUsableareasValue(id string) error {
	return errors.New("[ERROR] - Invalid Usableareas value - ID: " + id)
}

func NotElegiblePricePerMeterValue(id string) error {
	return errors.New("[ERROR] - Not Elegible Price per Meter value - ID: " + id)
}

func NotElegibleMonthlycondofeeValue(id string) error {
	return errors.New("[ERROR] - Not Elegible Monthlycondofee - ID: " + id)
}

func NotElegibleMaxRentalValue(id string) error {
	return errors.New("[ERROR] - Not Elegible max rental value - ID: " + id)
}

func NotElegibleMinSaleValue(id string) error {
	return errors.New("[ERROR] - Not Elegible min salve value - ID: " + id)
}
