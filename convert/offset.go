package convert

import "github.com/landru29/etopo2stl/earth"

// Offset ...
func Offset(xyz []earth.VectorA, offset float64) (xyzOut []earth.VectorA) {

	for _, xyzVector := range xyz {
		xyzOut = append(
			xyzOut,
			earth.VectorA{
				Lon:      xyzVector.Lon,
				Lat:      xyzVector.Lat,
				Altitude: xyzVector.Altitude + offset,
			},
		)
	}
	return
}
