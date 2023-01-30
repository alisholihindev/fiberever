package model

import "gorm.io/gorm"

// User struct
type User struct {
	gorm.Model
	Username string `gorm:"size:191;uniqueIndex;not null" json:"username"`
	Email    string `gorm:"size:191;uniqueIndex;not null;" json:"email"`
	Password string `gorm:"size:191;not null" json:"password"`
	Names    string `json:"size:191;names"`
}
