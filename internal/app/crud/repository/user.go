package repository

import (
	"crud-app/internal/app/crud/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user models.User) (models.User, error)
	GetUserByID(id uint) (models.User, error)
	GetUserByUsername(username string) (models.User, error)
	GetAllUsers() ([]models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(id uint) error
}

type GormUserRepository struct {
	DB *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{
		DB: db,
	}
}

func (r *GormUserRepository) CreateUser(user models.User) (models.User, error) {
	err := r.DB.Create(&user).Error
	return user, err
}

func (r *GormUserRepository) GetUserByID(id uint) (models.User, error) {
	var user models.User
	err := r.DB.First(&user, id).Error
	return user, err
}

func (r *GormUserRepository) GetUserByUsername(username string) (models.User, error) {
	var user models.User
	err := r.DB.Where(&models.User{Username: username}).First(&user).Error
	return user, err
}

func (r *GormUserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := r.DB.Find(&users).Error
	return users, err
}

func (r *GormUserRepository) UpdateUser(user models.User) (models.User, error) {
	err := r.DB.Save(&user).Error
	return user, err
}

func (r *GormUserRepository) DeleteUser(id uint) error {
	err := r.DB.Delete(&models.User{}, id).Error
	return err
}
