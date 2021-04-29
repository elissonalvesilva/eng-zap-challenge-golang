package errors

import "errors"

func InvalidLonAndLat(id string) error {
	return errors.New("[ID - " + id + "] - Invalid Longitude and Latitude")
}
