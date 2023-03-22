package domain

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type Publisher struct {
	ID          uint           `gorm:"type:INT;primaryKey;autoIncrement"`
	Name        string         `gorm:"type:VARCHAR(35);NOT NULL"`
	Description sql.NullString `gorm:"type:TEXT"`
	CreatedAt   time.Time      `gorm:"default:CURRENT_TIMESTAMP()"`
	UpdatedAt   sql.NullTime
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Books       []Book
}
