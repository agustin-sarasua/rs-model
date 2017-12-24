package m

import (
	"time"
)

var AccountTypes = map[string]struct{}{
	"CAJA_AHORRO":      {},
	"CUENTA_CORRIENTE": {}}

var Roles = map[string]string{
	"ADMIN":   "ABM_PROPIEDADES,ABM_PUBLICACIONES,ABM_USUARIOS",
	"MANAGER": "ABM_PROPIEDADES,ABM_PUBLICACIONES",
	"PUBLIC":  ""}

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
	Roles          string
}

type Account struct {
	ID        int64 `gorm:"primary_key;AUTO_INCREMENT"`
	Owner     *User `gorm:"ForeignKey:OwnerID;AssociationForeignKey:ID"`
	OwnerID   int64
	Name      string
	Type      string // CC, CA
	Number    string
	BankName  string
	Sucursal  string
	CreatedAt time.Time `gorm:"not null"`
}
