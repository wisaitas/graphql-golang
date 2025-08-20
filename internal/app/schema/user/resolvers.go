package user

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/wisaitas/graphql-golang/internal/app/model"
	"github.com/wisaitas/graphql-golang/internal/app/service"
)

// UserQueryResolver สำหรับ user queries
type UserQueryResolver struct {
	userService *service.UserService
}

// UserMutationResolver สำหรับ user mutations
type UserMutationResolver struct {
	userService *service.UserService
}

// NewUserQueryResolver สร้าง UserQueryResolver ใหม่
func NewUserQueryResolver(userService *service.UserService) *UserQueryResolver {
	return &UserQueryResolver{userService: userService}
}

// NewUserMutationResolver สร้าง UserMutationResolver ใหม่
func NewUserMutationResolver(userService *service.UserService) *UserMutationResolver {
	return &UserMutationResolver{userService: userService}
}

// === Query Resolvers ===

// Users resolver สำหรับ users query
func (r *UserQueryResolver) Users(p graphql.ResolveParams) (interface{}, error) {
	return r.userService.GetAllUsers()
}

// User resolver สำหรับ user query
func (r *UserQueryResolver) User(p graphql.ResolveParams) (interface{}, error) {
	id, ok := p.Args["id"].(string)
	if !ok {
		return nil, fmt.Errorf("id is required")
	}
	return r.userService.GetUserByID(id)
}

// === Mutation Resolvers ===

// CreateUser resolver สำหรับ createUser mutation
func (r *UserMutationResolver) CreateUser(p graphql.ResolveParams) (interface{}, error) {
	input, ok := p.Args["input"].(map[string]interface{})
	if !ok {
		return &model.UserResponse{
			Success: false,
			Message: "Invalid input",
		}, nil
	}

	createInput := &model.CreateUserInput{}

	if name, ok := input["name"].(string); ok {
		createInput.Name = name
	}
	if email, ok := input["email"].(string); ok {
		createInput.Email = email
	}
	if age, ok := input["age"].(int); ok {
		createInput.Age = age
	}

	user, err := r.userService.CreateUser(createInput)
	if err != nil {
		return &model.UserResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &model.UserResponse{
		Success: true,
		Message: "User created successfully",
		User:    user,
	}, nil
}

// UpdateUser resolver สำหรับ updateUser mutation
func (r *UserMutationResolver) UpdateUser(p graphql.ResolveParams) (interface{}, error) {
	input, ok := p.Args["input"].(map[string]interface{})
	if !ok {
		return &model.UserResponse{
			Success: false,
			Message: "Invalid input",
		}, nil
	}

	updateInput := &model.UpdateUserInput{}

	if id, ok := input["id"].(string); ok {
		updateInput.ID = id
	}
	if name, ok := input["name"].(string); ok {
		updateInput.Name = &name
	}
	if email, ok := input["email"].(string); ok {
		updateInput.Email = &email
	}
	if age, ok := input["age"].(int); ok {
		updateInput.Age = &age
	}

	user, err := r.userService.UpdateUser(updateInput)
	if err != nil {
		return &model.UserResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &model.UserResponse{
		Success: true,
		Message: "User updated successfully",
		User:    user,
	}, nil
}

// DeleteUser resolver สำหรับ deleteUser mutation
func (r *UserMutationResolver) DeleteUser(p graphql.ResolveParams) (interface{}, error) {
	id, ok := p.Args["id"].(string)
	if !ok {
		return &model.BaseResponse{
			Success: false,
			Message: "ID is required",
		}, nil
	}

	err := r.userService.DeleteUser(id)
	if err != nil {
		return &model.BaseResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &model.BaseResponse{
		Success: true,
		Message: "User deleted successfully",
	}, nil
}
