package models

import (
    "time"

    "gorm.io/gorm"
)

type User struct {
    ID        uint
    Username  string `gorm:"unique;not null"`
    Hash      string `gorm:"not null"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`
}