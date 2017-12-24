package m

import "time"

var Guarantee = map[string]struct{}{
	"OWN":          {},
	"ANDA":         {},
	"PORTO_SEGURO": {},
	"CGN":          {},
	"PROPIEDAD":    {}}

type Transaction struct {
	ID            int64        `gorm:"primary_key;AUTO_INCREMENT"`
	CreatedAt     time.Time    `gorm:"not null"`
	Publication   *Publication `gorm:"ForeignKey:PublicationID;AssociationForeignKey:ID"`
	PublicationID int64
	Client        *User `gorm:"ForeignKey:ClientID;AssociationForeignKey:ID"`
	ClientID      int64
	Guarantee     string
}
