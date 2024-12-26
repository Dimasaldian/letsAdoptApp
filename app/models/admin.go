package models

import (
	"gorm.io/gorm"
)

type Admin struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"size:100;not null"`
	Email      string `gorm:"size:100;unique;not null"`
	Password   string `gorm:"not null"`
	Privileges string `gorm:"type:text"`
}

// GetAdmin retrieves all admins from the database
func (a *Admin) GetAdmin(db *gorm.DB) ([]Admin, error) {
	var admins []Admin
	if err := db.Find(&admins).Error; err != nil {
		return nil, err
	}
	return admins, nil
}
