package service

import (
	"fmt"
	"time"

	"github.com/wisaitas/graphql-golang/internal/app/model"
)

// UserService สำหรับ business logic ของ User
type UserService struct {
	// ในอนาคตอาจจะเชื่อมต่อกับ database
	users map[string]*model.User
}

// NewUserService สร้าง UserService instance ใหม่
func NewUserService() *UserService {
	// สร้างข้อมูลตัวอย่าง
	users := map[string]*model.User{
		"1": {
			ID:        "1",
			Name:      "John Doe",
			Email:     "john@example.com",
			Age:       30,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		"2": {
			ID:        "2",
			Name:      "Jane Smith",
			Email:     "jane@example.com",
			Age:       25,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	return &UserService{users: users}
}

// GetAllUsers ดึงข้อมูลผู้ใช้ทั้งหมด
func (s *UserService) GetAllUsers() ([]*model.User, error) {
	users := make([]*model.User, 0, len(s.users))
	for _, user := range s.users {
		users = append(users, user)
	}
	return users, nil
}

// GetUserByID ดึงข้อมูลผู้ใช้ตาม ID
func (s *UserService) GetUserByID(id string) (*model.User, error) {
	user, exists := s.users[id]
	if !exists {
		return nil, fmt.Errorf("user with id %s not found", id)
	}
	return user, nil
}

// CreateUser สร้างผู้ใช้ใหม่
func (s *UserService) CreateUser(input *model.User) (*model.User, error) {
	// ตรวจสอบข้อมูล
	if input.Name == "" {
		return nil, fmt.Errorf("name is required")
	}
	if input.Email == "" {
		return nil, fmt.Errorf("email is required")
	}

	// สร้าง ID ใหม่ (แบบง่ายๆ)
	newID := fmt.Sprintf("%d", len(s.users)+1)

	user := &model.User{
		ID:        newID,
		Name:      input.Name,
		Email:     input.Email,
		Age:       input.Age,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	s.users[newID] = user
	return user, nil
}

// UpdateUser อัพเดทข้อมูลผู้ใช้
func (s *UserService) UpdateUser(input *model.User) (*model.User, error) {
	user, exists := s.users[input.ID]
	if !exists {
		return nil, fmt.Errorf("user with id %s not found", input.ID)
	}

	// อัพเดทข้อมูลเฉพาะที่ส่งมา
	if input.Name != "" {
		user.Name = input.Name
	}
	if input.Email != "" {
		user.Email = input.Email
	}
	if input.Age != 0 {
		user.Age = input.Age
	}
	user.UpdatedAt = time.Now()

	s.users[input.ID] = user
	return user, nil
}

// DeleteUser ลบผู้ใช้
func (s *UserService) DeleteUser(id string) error {
	if _, exists := s.users[id]; !exists {
		return fmt.Errorf("user with id %s not found", id)
	}
	delete(s.users, id)
	return nil
}
