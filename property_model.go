package m

import (
	"time"
)

type Property struct {
	ID                 uint64 `gorm:"primary_key;AUTO_INCREMENT"`
	Description        string
	Type               string `gorm:"not null"` //Apartamento, Casa, LocalComercial
	Orientation        string
	State              int // Del 0 al 10, 10 seria a estrenar
	Bedrooms           int `gorm:"not null"`
	BedroomsSizes      string
	KitchenSize        string
	LivingRoomSize     string
	CourtyardSize      int
	Bathrooms          int
	Size               int `gorm:"not null"`
	ConstructionYear   int
	Padron             string
	BuildingName       string
	Address            *Address `gorm:"ForeignKey:AddressID;AssociationForeignKey:ID"`
	AddressID          uint
	ApartmentsPerFloor int
	Floors             int
	TerraceSize        string
	BalconySize        string
	Showers            int
	Expenses           int32
	Amenities          string `gorm:"type:varchar(64)"`
	CreatedAt          time.Time
	PropertyState      *PropertyState `gorm:"-"`
	Elevators          int
	GarageSize         int
}

type PropertyState struct {
	ID               uint64 `gorm:"primary_key;AUTO_INCREMENT"`
	PropertyID       uint64 `gorm:"index"`
	Overall          int
	Safety           int
	Windows          int
	Floors           int
	Kitchen          int
	KitchenFurniture int
	NaturalLight     int
	Appearence       int
	WaterPreasure    int
	CreatedAt        time.Time
}

type Address struct {
	ID              uint64 `gorm:"primary_key;AUTO_INCREMENT"`
	Street          string `gorm:"not null"`
	Number          string `gorm:"not null"`
	ApartmentNumber string
	Neighborhood    string `gorm:"not null"`
	City            string `gorm:"not null"`
	Country         string `gorm:"not null"`
	PostalCode      string
	Location        Location `gorm:"embedded;embedded_prefix:loc_"`
}

type Location struct {
	Latitude  float32 `gorm:"not null"`
	Longitude float32 `gorm:"not null"`
}
