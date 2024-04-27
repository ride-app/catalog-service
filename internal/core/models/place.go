package models

type PlaceType int

const (
	PLACE_TYPE_UNSPECIFIED PlaceType = iota
	PLACE_TYPE_PICKUP
	PLACE_TYPE_DROPOFF
)

type Place struct {
	Id             string
	DisplayName    string
	Address        string
	Location       LatLng
	DistanceMeters int32
	Type           PlaceType
}
