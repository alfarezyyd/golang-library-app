package domain

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type Kind struct {
	ID          uint           `gorm:"primaryKey;autoIncrement"`
	Name        string         `gorm:"type:VARCHAR(35);NOT NULL"`
	Description sql.NullString `gorm:"type:TEXT"`
	CreatedAt   time.Time      `gorm:"default:CURRENT_TIMESTAMP()"`
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Books       []*Book        `gorm:"many2many:books_kinds"`
}
