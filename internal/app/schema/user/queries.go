package user

import (
	"github.com/graphql-go/graphql"
	gqlTypes "github.com/wisaitas/graphql-golang/internal/app/graphql"
	"github.com/wisaitas/graphql-golang/internal/app/service"
)

// GetQueries ส่งกลับ user queries ทั้งหมด
func GetQueries(userService *service.UserService) graphql.Fields {
	return graphql.Fields{
		"users": &graphql.Field{
			Type:        graphql.NewList(gqlTypes.UserType),
			Description: "Get all users",
			Resolve:     NewUserQueryResolver(userService).Users,
		},
		"user": &graphql.Field{
			Type:        gqlTypes.UserType,
			Description: "Get user by ID",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: NewUserQueryResolver(userService).User,
		},
	}
}
