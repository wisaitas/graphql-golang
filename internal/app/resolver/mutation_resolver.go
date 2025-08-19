package resolver

import (
	"github.com/graphql-go/graphql"
	"github.com/wisaitas/graphql-golang/internal/app/model"
	"github.com/wisaitas/graphql-golang/internal/app/service"
)

// MutationResolver สำหรับ GraphQL mutations
type MutationResolver struct {
	userService *service.UserService
}

// NewMutationResolver สร้าง MutationResolver ใหม่
func NewMutationResolver(userService *service.UserService) *MutationResolver {
	return &MutationResolver{
		userService: userService,
	}
}

// CreateUser resolver สำหรับ createUser mutation
func (r *MutationResolver) CreateUser(p graphql.ResolveParams) (interface{}, error) {
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
func (r *MutationResolver) UpdateUser(p graphql.ResolveParams) (interface{}, error) {
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
func (r *MutationResolver) DeleteUser(p graphql.ResolveParams) (interface{}, error) {
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
