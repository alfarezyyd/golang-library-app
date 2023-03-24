package domain

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID              int    `gorm:"primaryKey;autoIncrement"`
	MemberID        int    `gorm:"UNIQUE"`
	Username        string `gorm:"type:VARCHAR(30);NOT NULL"`
	Email           string `gorm:"type:VARCHAR(50);NOT NULL;UNIQUE"`
	Password        string `gorm:"type:VARCHAR(100);NOT NULL"`
	EmailVerifiedAt sql.NullTime
	CreatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP()"`
	UpdatedAt       sql.NullTime
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}
