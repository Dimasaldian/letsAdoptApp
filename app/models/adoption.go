package models

import (
	"gorm.io/gorm"
)

type Adoption struct {
	gorm.Model
	IDAdopt          uint   `gorm:"primaryKey;autoIncrement:true" json:"id_adopt"`
	Name             string `gorm:"size:100;not null" json:"name"`    // Nama pengguna
	Email            string `gorm:"size:100;not null" json:"email"`   // Email pengguna
	PetID            uint   `json:"pet_id"`                           // ID Pet
	Pet              Pet    `gorm:"foreignKey:PetID;references:ID" json:"pet"` // Relasi ke tabel Pet
	Reason           string `gorm:"type:text;not null" json:"reason"` // Alasan adopsi
	Status           string `gorm:"size:10;not null;default:'pending'" json:"status"` // Status adopsi
	NotificationSent bool   `gorm:"default:false" json:"notification_sent"` // Notifikasi terkirim
}
