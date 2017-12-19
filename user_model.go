package m

import (
	"time"
)

var AccountTypes = map[string]struct{}{
	"CAJA_AHORRO":      {},
	"CUENTA_CORRIENTE": {}}

type User struct {
	ID             int64  `gorm:"primary_key;AUTO_INCREMENT"`
	FirebaseID     string `gorm:"index"`
	Name           string `gorm:"not null"`
	LastName       string
	Email          string `gorm:"not null"`
	Phone          string `gorm:"not null"`
	DocumentNumber string
	DocumentType   string
	CreatedAt      time.Time `gorm:"not null"`
}

type Account struct {
	ID       int64 `gorm:"primary_key;AUTO_INCREMENT"`
	Owner    *User `gorm:"ForeignKey:OwnerID;AssociationForeignKey:ID"`
	OwnerID  int64
	Name     string
	Type     string // CC, CA
	Number   string
	BankName string
	Sucursal string
}
