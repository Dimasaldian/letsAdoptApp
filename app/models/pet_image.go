package models

import (
	"gorm.io/gorm"
)

type PetImage struct {
    gorm.Model
    PetID uint   `json:"pet_id"`
    Pet   Pet    `gorm:"foreignKey:PetID;references:ID" json:"pet"`
    URL   string `gorm:"size:255;not null" json:"url"`
}