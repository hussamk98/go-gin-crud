package models

import (
	"crud-app/internal/app/crud/dto"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string    `gorm:"size:191;uniqueIndex;not null"`
	Password  string    `gorm:"not null"`
	FirstName string    `gorm:"not null"`
	LastName  string    `gorm:"not null"`
	Email     string    `gorm:"size:191;uniqueIndex;not null"`
	IsAdmin   bool      `gorm:"default:false"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func NewUserFromDTO(userDTO dto.CreateUserDTO) User {
	return User{
		Username:  userDTO.Username,
		Password:  userDTO.Password,
		FirstName: userDTO.FirstName,
		LastName:  userDTO.LastName,
		Email:     userDTO.Email,
		IsAdmin:   userDTO.IsAdmin,
	}
}

func (u *User) ToUserResponseDTO() *dto.UserResponseDTO {
	return &dto.UserResponseDTO{
		ID:        u.ID,
		Username:  u.Username,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		IsAdmin:   u.IsAdmin,
		CreatedAt: u.CreatedAt.Format(time.RFC3339),
		UpdatedAt: u.UpdatedAt.Format(time.RFC3339),
	}
}
