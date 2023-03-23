package domain

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type Book struct {
	ID              uint `gorm:"primaryKey;autoIncrement"`
	PublisherID     uint
	ISBN            string    `gorm:"type:VARCHAR(30);NOT NULL"`
	Title           string    `gorm:"type:VARCHAR(50);NOT NULL"`
	Author          string    `gorm:"type:VARCHAR(50);NOT NULL"`
	PublicationYear string    `gorm:"type:VARCHAR(4);NOT NULL"`
	Amount          uint      `gorm:"type:INT(11);NOT NULL"`
	Bookshelf       string    `gorm:"type:VARCHAR(10);NOT NULL"`
	CreatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP()"`
	UpdatedAt       sql.NullTime
	DeletedAt       gorm.DeletedAt
	Kinds           []Kind `gorm:"many2many:books_kinds"`
}
