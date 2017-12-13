package m

import (
	"time"
)

type Property struct {
	ID                 uint `gorm:"AUTO_INCREMENT"`
	Description        string
	Type               string `gorm:"not null"` //Apartamento, Casa, LocalComercial
	Orientation        string
	State              int      // Del 0 al 10, 10 seria a estrenar
	Bedrooms           int      `gorm:"not null"`
	BedroomsSizes      []string `gorm:"type:varchar(64)"`
	KitchenSize        string
	LivingRoomSize     string
	CourtyardSize      int
	Bathrooms          int
	Size               int `gorm:"not null"`
	ConstructionYear   int
	Padron             string
	BuildingName       string
	Address            *Address
	AddressID          uint `gorm:"ForeignKey:id"`
	ApartmentsPerFloor int
	Floors             int
	TerraceSize        string
	BalconySize        string
	Showers            int
	Expenses           int64
	Amenities          []string `gorm:"type:varchar(64)"`
	CreatedAt          time.Time
	PropertyState      *PropertyState `gorm:"-"`
	Elevators          int
	GarageSize         int
}

type PropertyState struct {
	PropertyID       uint
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
	ID              uint   `gorm:"AUTO_INCREMENT"`
	Street          string `gorm:"not null"`
	Number          string `gorm:"not null"`
	ApartmentNumber string
	Neighborhood    string `gorm:"not null"`
	City            string `gorm:"not null"`
	Country         string `gorm:"not null"`
	PostalCode      string
	Location        *Location `gorm:"embedded;embedded_prefix:loc_"`
}

type Location struct {
	Latitude  string `gorm:"not null"`
	Longitude string `gorm:"not null"`
}
