package domain

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type Visitor struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	NIN       string    `gorm:"type:VARCHAR(15);NOT NULL"`
	Name      string    `gorm:"type:VARCHAR(50);NOT NULL"`
	Instance  string    `gorm:"type:VARCHAR(50);NOT NULL"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP()"`
	UpdatedAt sql.NullTime
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
