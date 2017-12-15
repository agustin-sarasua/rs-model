package m

import (
	"time"
)

type User struct {
	UID            string `gorm:"primary_key;AUTO_INCREMENT"`
	Name           string `gorm:"not null"`
	LastName       string
	Email          string `gorm:"not null"`
	Phone          string `gorm:"not null"`
	DocumentNumber string
	DocumentType   string
	CreatedAt      time.Time `gorm:"not null"`
}
