package repositories

import (
	"_waysbook/models"

	"gorm.io/gorm"
)

type UserRepository interface {
  FindUsers() ([]models.User, error)
  GetUser(ID int) (models.User, error)
  CreateUser(user models.User) (models.User, error)
  UpdateUser(user models.User, ID int) (models.User, error)
  DeleteUser(user models.User, ID int) (models.User, error)
}

func RepositoryUser(db *gorm.DB) *repository {
  return &repository{db}
}

func (r *repository) FindUsers() ([]models.User, error) {
  var users []models.User
  err := r.db.Preload("Transaction").Preload("Transaction.Book").Preload("Transaction.User").Find(&users).Error

  return users, err
}

func (r *repository) GetUser(ID int) (models.User, error) {
  var user models.User
  err := r.db.Preload("Transaction").Preload("Transaction.Book").Preload("Transaction.User").First(&user).Error

  return user, err
}

 func (r *repository) CreateUser(user models.User) (models.User, error) {
	err := r.db.Preload("Transaction").Create(&user).Error
  
	return user, err
  }

func (r *repository) UpdateUser(user models.User, ID int) (models.User, error) {
	err := r.db.Preload("Transaction").Raw("UPDATE users SET full_name=?, email=?, password=? WHERE id=?", user.FullName, user.Email, user.Password,ID).Scan(&user).Error
  
	return user, err
  }

 func (r *repository) DeleteUser(user models.User,ID int) (models.User, error) {
	err := r.db.Raw("DELETE FROM users WHERE id=?",ID).Scan(&user).Error
  
	return user, err
  }