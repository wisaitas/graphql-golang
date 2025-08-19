package resolver

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/wisaitas/graphql-golang/internal/app/service"
)

// QueryResolver สำหรับ GraphQL queries
type QueryResolver struct {
	userService *service.UserService
}

// NewQueryResolver สร้าง QueryResolver ใหม่
func NewQueryResolver(userService *service.UserService) *QueryResolver {
	return &QueryResolver{
		userService: userService,
	}
}

// Hello resolver สำหรับ hello query
func (r *QueryResolver) Hello(p graphql.ResolveParams) (interface{}, error) {
	name, ok := p.Args["name"].(string)
	if !ok || name == "" {
		return "Hello, World! This is GraphQL in Go! 🚀", nil
	}
	return fmt.Sprintf("Hello, %s! 👋", name), nil
}

// Users resolver สำหรับ users query
func (r *QueryResolver) Users(p graphql.ResolveParams) (interface{}, error) {
	return r.userService.GetAllUsers()
}

// User resolver สำหรับ user query
func (r *QueryResolver) User(p graphql.ResolveParams) (interface{}, error) {
	id, ok := p.Args["id"].(string)
	if !ok {
		return nil, fmt.Errorf("id is required")
	}
	return r.userService.GetUserByID(id)
}
