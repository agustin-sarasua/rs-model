package m

import (
	c "github.com/agustin-sarasua/rs-common"
)

type Publication struct {
	ID              int64
	Operation       string //Sale, Rent
	Property        *Property
	Price           int64
	Timestamp       string
	Owner           *User
	MinContractTime int64
}

type Property struct {
	ID                 int64
	Description        string
	Type               string //Apartamento, Casa, LocalComercial
	Orientation        string
	State              int // Del 0 al 10, 10 seria a estrenar
	Bedrooms           int
	BedroomsSizes      []string
	KitchenSize        string
	LivingRoomSize     string
	CourtyardSize      int
	Bathrooms          int
	Size               int
	ConstructionYear   int
	Padron             string
	BuildingName       string
	Address            *Address
	ApartmentsPerFloor int
	Floors             int
	TerraceSize        string
	BalconySize        string
	Showers            int
	Expenses           int64
	Amenities          []string
	CreatedDate        c.DateTime
	PropertyState      *PropertyState
	Elevators          int
	GarageSize         int
}

type PropertyState struct {
	PropertyId       int
	Overall          int
	Safety           int
	Windows          int
	Floors           int
	Kitchen          int
	KitchenFurniture int
	NaturalLight     int
	Appearence       int
	WaterPreasure    int
}

type Address struct {
	Street          string
	Number          string
	ApartmentNumber string
	Neighborhood    string
	City            string
	Country         string
	PostalCode      string
	Location        *Location
}

type Location struct {
	Latitude  string
	Longitude string
}
