package model

import "gorm.io/gorm"

// User struct
type Dokter struct {
	gorm.Model
	Names string `json:"size:191;names"`
}
