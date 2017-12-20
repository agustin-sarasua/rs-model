package m

import "time"

var Operation = map[string]struct{}{
	"SALE": {},
	"RENT": {}}

type Publication struct {
	ID                int64
	Operation         string    //Sale, Rent
	Property          *Property `gorm:"ForeignKey:PropertyID;AssociationForeignKey:ID"`
	PropertyID        uint
	Price             int64
	Owner             *User `gorm:"ForeignKey:OwnerID;AssociationForeignKey:ID"`
	OwnerID           int64
	MinContractTime   int64
	MaxContractTime   int64
	StartDate         time.Time
	EndDate           time.Time
	CreatedAt         time.Time
	ReservationAmount int64
}
