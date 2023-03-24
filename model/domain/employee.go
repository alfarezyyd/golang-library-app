package domain

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type Employee struct {
	ID              int       `gorm:"primaryKey;autoIncrement"`
	Name            string    `gorm:"type:VARCHAR(50);NOT NULL"`
	Gender          string    `gorm:"type:ENUM('Men', 'Woman');NOT NULL"`
	Position        string    `gorm:"type:VARCHAR(50);NOT NULL"`
	TelephoneNumber string    `gorm:"type:VARCHAR(15);NOT NULL"`
	Address         string    `gorm:"type:VARCHAR(100);NOT NULL"`
	CreatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP()"`
	UpdatedAt       sql.NullTime
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}
