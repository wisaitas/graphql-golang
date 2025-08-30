package resolver

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/wisaitas/graphql-golang/internal/app/model"
	"github.com/wisaitas/graphql-golang/internal/app/response"
	"github.com/wisaitas/graphql-golang/internal/app/service"
)

type UserResolver interface {
	Users(p graphql.ResolveParams) (interface{}, error)
	User(p graphql.ResolveParams) (interface{}, error)
	CreateUser(p graphql.ResolveParams) (interface{}, error)
	UpdateUser(p graphql.ResolveParams) (interface{}, error)
	DeleteUser(p graphql.ResolveParams) (interface{}, error)
}

type userResolver struct {
	userService *service.UserService
}

func NewUserResolver(
	userService *service.UserService,
) UserResolver {
	return &userResolver{
		userService: userService,
	}
}

func (r *userResolver) Users(p graphql.ResolveParams) (interface{}, error) {
	return r.userService.GetAllUsers()
}

func (r *userResolver) User(p graphql.ResolveParams) (interface{}, error) {
	id, ok := p.Args["id"].(string)
	if !ok {
		return nil, fmt.Errorf("id is required")
	}
	return r.userService.GetUserByID(id)
}

func (r *userResolver) CreateUser(p graphql.ResolveParams) (interface{}, error) {
	input, ok := p.Args["input"].(map[string]interface{})
	if !ok {
		return &response.UserResponse{
			BaseResponse: response.BaseResponse{
				Success: false,
				Message: "Invalid input",
			},
		}, nil
	}

	createInput := &model.User{}

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
		return &response.UserResponse{
			BaseResponse: response.BaseResponse{
				Success: false,
				Message: err.Error(),
			},
		}, nil
	}

	return &response.UserResponse{
		BaseResponse: response.BaseResponse{
			Success: true,
			Message: "User created successfully",
		},
		User: user,
	}, nil
}

// UpdateUser resolver สำหรับ updateUser mutation
func (r *userResolver) UpdateUser(p graphql.ResolveParams) (interface{}, error) {
	input, ok := p.Args["input"].(map[string]interface{})
	if !ok {
		return &response.UserResponse{
			BaseResponse: response.BaseResponse{
				Success: false,
				Message: "Invalid input",
			},
		}, nil
	}

	updateInput := &model.User{}

	if id, ok := input["id"].(string); ok {
		updateInput.ID = id
	}
	if name, ok := input["name"].(string); ok {
		updateInput.Name = name
	}
	if email, ok := input["email"].(string); ok {
		updateInput.Email = email
	}
	if age, ok := input["age"].(int); ok {
		updateInput.Age = age
	}

	user, err := r.userService.UpdateUser(updateInput)
	if err != nil {
		return &response.UserResponse{
			BaseResponse: response.BaseResponse{
				Success: false,
				Message: err.Error(),
			},
		}, nil
	}

	return &response.UserResponse{
		BaseResponse: response.BaseResponse{
			Success: true,
			Message: "User updated successfully",
		},
		User: user,
	}, nil
}

// DeleteUser resolver สำหรับ deleteUser mutation
func (r *userResolver) DeleteUser(p graphql.ResolveParams) (interface{}, error) {
	id, ok := p.Args["id"].(string)
	if !ok {
		return &response.BaseResponse{
			Success: false,
			Message: "ID is required",
		}, nil
	}

	err := r.userService.DeleteUser(id)
	if err != nil {
		return &response.BaseResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &response.BaseResponse{
		Success: true,
		Message: "User deleted successfully",
	}, nil
}
