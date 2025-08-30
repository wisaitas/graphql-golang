package service

import (
	"errors"
	"strconv"

	"github.com/wisaitas/graphql-golang/internal/appv2/entity"
	"github.com/wisaitas/graphql-golang/internal/appv2/model"
	"github.com/wisaitas/graphql-golang/internal/appv2/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(input model.CreateUserInput) (*model.User, error)
	GetUserByID(id string) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	GetAllUsers() ([]model.User, error)
	UpdateUser(id string, input model.UpdateUserInput) (*model.User, error)
	DeleteUser(id string) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) CreateUser(input model.CreateUserInput) (*model.User, error) {
	// ตรวจสอบว่า email ไม่ซ้ำ
	existingUser, _ := s.userRepo.GetByEmail(input.Email)
	if existingUser != nil {
		return nil, errors.New("email already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		Username: input.Username,
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	if err = s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user.EntityToModel(), nil
}

func (s *userService) GetUserByID(id string) (*model.User, error) {
	userID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return nil, errors.New("invalid user ID")
	}

	user, err := s.userRepo.GetByID(uint(userID))
	if err != nil {
		return nil, err
	}

	return user.EntityToModel(), nil
}

func (s *userService) GetUserByEmail(email string) (*model.User, error) {
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	return user.EntityToModel(), nil
}

func (s *userService) GetAllUsers() ([]model.User, error) {
	userEntities, err := s.userRepo.GetAll()
	if err != nil {
		return nil, err
	}

	var users []model.User
	for _, user := range userEntities {
		userModel := user.EntityToModel()
		users = append(users, *userModel)
	}

	return users, nil
}

func (s *userService) UpdateUser(id string, input model.UpdateUserInput) (*model.User, error) {
	userID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return nil, errors.New("invalid user ID")
	}

	user, err := s.userRepo.GetByID(uint(userID))
	if err != nil {
		return nil, err
	}

	updateData := &entity.User{}
	if input.Username != nil {
		updateData.Username = *input.Username
	}
	if input.Email != nil {
		updateData.Email = *input.Email
	}
	if input.Password != nil {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*input.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		updateData.Password = string(hashedPassword)
	}

	err = s.userRepo.Update(uint(userID), updateData)
	if err != nil {
		return nil, err
	}

	return user.EntityToModel(), nil
}

func (s *userService) DeleteUser(id string) error {
	userID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return errors.New("invalid user ID")
	}

	return s.userRepo.Delete(uint(userID))
}
