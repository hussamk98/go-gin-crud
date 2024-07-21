package service

import (
	"crud-app/internal/app/crud/dto"
	"crud-app/internal/app/crud/models"
	"crud-app/internal/app/crud/repository"
)

type UserService struct {
	UserRepo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		UserRepo: repo,
	}
}

func (s *UserService) CreateUser(user models.User) (dto.UserResponseDTO, error) {
	createdUser, err := s.UserRepo.CreateUser(user)
	if err != nil {
		return dto.UserResponseDTO{}, err
	}
	return *createdUser.ToUserResponseDTO(), nil
}

func (s *UserService) GetUserByID(id uint) (dto.UserResponseDTO, error) {
	user, err := s.UserRepo.GetUserByID(id)
	if err != nil {
		return dto.UserResponseDTO{}, err
	}
	return *user.ToUserResponseDTO(), nil
}

func (s *UserService) GetUserByUsername(username string) (dto.UserResponseDTO, error) {
	user, err := s.UserRepo.GetUserByUsername(username)
	if err != nil {
		return dto.UserResponseDTO{}, err
	}
	return *user.ToUserResponseDTO(), nil
}

func (s *UserService) GetAllUsers() ([]dto.UserResponseDTO, error) {
	users, err := s.UserRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}
	var userDTOs []dto.UserResponseDTO
	for _, user := range users {
		userDTOs = append(userDTOs, *user.ToUserResponseDTO())
	}
	return userDTOs, nil
}

func (s *UserService) UpdateUser(user models.User) (dto.UserResponseDTO, error) {
	updatedUser, err := s.UserRepo.UpdateUser(user)
	if err != nil {
		return dto.UserResponseDTO{}, err
	}
	return *updatedUser.ToUserResponseDTO(), nil
}

func (s *UserService) DeleteUser(id uint) error {
	return s.UserRepo.DeleteUser(id)
}
