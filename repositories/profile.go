package repositories

import (
	"_waysbook/models"

	"gorm.io/gorm"
)

type ProfileRepository interface {
	GetProfile(ID int) (models.Profile, error)
}

func RepositoryProfile(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetProfile(ID int) (models.Profile, error) {
	var profile models.Profile
	err := r.db.Preload("User").Preload("Transaction").Preload("Transaction.Book").First(&profile, "user_id = ?", ID).Error

	return profile, err
}
