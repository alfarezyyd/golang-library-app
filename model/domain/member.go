package domain

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type Member struct {
	ID              int       `gorm:"primaryKey;autoIncrement"`
	NIN             string    `gorm:"type:VARCHAR(15);NOT NULL"`
	Name            string    `gorm:"type:VARCHAR(50);NOT NULL"`
	BirthPlace      string    `gorm:"type:VARCHAR(50);NOT NULL"`
	BirthDate       time.Time `gorm:"type:DATE;NOT NULL"`
	Address         string    `gorm:"type:VARCHAR(100);NOT NULL"`
	Gender          string    `gorm:"type:ENUM('Men', 'Woman');NOT NULL"`
	TelephoneNumber string    `gorm:"type:VARCHAR(15);NOT NULL"`
	CreatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP()"`
	UpdatedAt       sql.NullTime
	DeletedAt       gorm.DeletedAt `gorm:"index"`
	User            User
}
