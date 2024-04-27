package models

import (
	"google.golang.org/genproto/googleapis/type/latlng"
)

type LatLng struct {
	Latitude  float64
	Longitude float64
}

func LatLngFromProtobuf(pb *latlng.LatLng) LatLng {
	return LatLng{
		Latitude:  pb.Latitude,
		Longitude: pb.Longitude,
	}
}

func (value LatLng) ToProtobuf() *latlng.LatLng {
	return &latlng.LatLng{
		Latitude:  value.Latitude,
		Longitude: value.Longitude,
	}
}
