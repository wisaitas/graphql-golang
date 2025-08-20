package user

import (
	"github.com/graphql-go/graphql"
	gqlTypes "github.com/wisaitas/graphql-golang/internal/app/graphql"
	"github.com/wisaitas/graphql-golang/internal/app/service"
)

// GetMutations ส่งกลับ user mutations ทั้งหมด
func GetMutations(userService *service.UserService) graphql.Fields {
	return graphql.Fields{
		"createUser": &graphql.Field{
			Type:        gqlTypes.UserResponseType,
			Description: "Create a new user",
			Args: graphql.FieldConfigArgument{
				"input": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(gqlTypes.CreateUserInputType),
				},
			},
			Resolve: NewUserMutationResolver(userService).CreateUser,
		},
		"updateUser": &graphql.Field{
			Type:        gqlTypes.UserResponseType,
			Description: "Update an existing user",
			Args: graphql.FieldConfigArgument{
				"input": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(gqlTypes.UpdateUserInputType),
				},
			},
			Resolve: NewUserMutationResolver(userService).UpdateUser,
		},
		"deleteUser": &graphql.Field{
			Type:        gqlTypes.BaseResponseType,
			Description: "Delete a user",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: NewUserMutationResolver(userService).DeleteUser,
		},
	}
}
