package service

import (
	"alghanim/mediacmsAPI/model"
	"alghanim/mediacmsAPI/repository"
)

type UserService interface {
	Get(id int) (*model.User, error)
	Update(id int, user *model.User) (*model.User, error)
	// TODO: Add other operations like Create, Update, Delete
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepository: userRepo}
}

func (s *userService) Get(id int) (*model.User, error) {
	return s.userRepository.Get(id)
}

func (s *userService) Update(id int, user *model.User) (*model.User, error) {
	return s.userRepository.Update(id, user)
}

// TODO: Implement other operations like Create, Update, Delete
