package m

import (
	"time"
)

var Orientation = map[string]struct{}{
	"FRENTE":       {},
	"CONTRAFRENTE": {}}

type Property struct {
	ID               uint64 `gorm:"primary_key;AUTO_INCREMENT"`
	Title            string `gorm:"not null"`
	Description      string
	Type             string `gorm:"not null"` //Apartamento, Casa, LocalComercial
	Disposition      string
	Orientation      string
	State            string
	Bedrooms         int `gorm:"not null"`
	Bathrooms        int `gorm:"not null"`
	Garages          int `gorm:"not null"`
	Size             int `gorm:"not null"`
	ConstructionSize int `gorm:"not null"`
	ConstructionYear int
	Address          *Address `gorm:"ForeignKey:AddressID;AssociationForeignKey:ID"`
	AddressID        uint
	Floors           int
	TerraceSize      string
	Expenses         int32
	CreatedAt        time.Time
	Amenities        string `gorm:"type:varchar(1024)"`
	// HasKitchenette bool
	// Kitchens       int
	// KitchenSizes   string
	// LivingroomSize int
	// CourtyardSize  int
	// Padron             string
	// BuildingName       string
	// ApartmentsPerFloor int
	// BalconySize        string
	// Showers            int
	// PropertyState      *PropertyState `gorm:"-"`
	// Elevators          int
	// GarageSize         int
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
