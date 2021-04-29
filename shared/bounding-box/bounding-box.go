package boundingbox

import (
	"github.com/elissonalvesilva/eng-zap-challenge-golang/domain/protocols"
)

func IsInBoundingBoxZap(data protocols.Imovel) bool {

	minlon := -46.693419 // left
	minlat := -23.568704 // top
	maxlon := -46.641146 // right
	maxlat := -23.546686 // botom
	point := data.Address.Geolocation.Location
	/* Check latitude bounds first. */
	if minlat >= point.Lat && point.Lat >= maxlat {
		if minlon <= maxlon && minlon <= point.Lon && point.Lon <= maxlon {
			return true
		} else if minlon > maxlon && (minlon <= point.Lon || point.Lon <= maxlon) {
			return true
		}
	}

	return false
}
