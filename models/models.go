
package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
    gorm.Model
    ID       uint   `gorm:"primaryKey;autoIncrement"`
    Username string `gorm:"not null"`
    Email    string `gorm:"unique;not null"`
    Password string `gorm:"not null"`
    CreatedAt time.Time
    UpdatedAt time.Time
    Photos    []Photo    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
}

type Photo struct {
    gorm.Model
    ID       uint   `gorm:"primaryKey;autoIncrement"`
    Title    string `gorm:"not null"`
    Caption  string
    PhotoUrl  string `json:"photo_url" form:"photo_url" valid:"required~Photo URL of your photo is required"`
    UserID   uint   `gorm:"not null"`
    CreatedAt time.Time
    UpdatedAt time.Time
}
