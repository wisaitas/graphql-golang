package schema

import (
	"github.com/graphql-go/graphql"
	gqlTypes "github.com/wisaitas/graphql-golang/internal/app/graphql"
	"github.com/wisaitas/graphql-golang/internal/app/resolver"
)

type UserSchema struct {
	Queries   graphql.Fields
	Mutations graphql.Fields
}

func NewUserSchema(
	userResolver resolver.UserResolver,
	userType *gqlTypes.UserType,
) *UserSchema {
	return &UserSchema{
		Queries: graphql.Fields{
			"users": &graphql.Field{
				Type:        graphql.NewList(userType.User),
				Description: "Get all users",
				Resolve:     userResolver.Users,
			},
			"user": &graphql.Field{
				Type:        userType.User,
				Description: "Get a user by ID",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: userResolver.User,
			},
		},
		Mutations: graphql.Fields{
			"createUser": &graphql.Field{
				Type:        userType.User,
				Description: "Create a new user",
				Args: graphql.FieldConfigArgument{
					"input": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(userType.User),
					},
				},
				Resolve: userResolver.CreateUser,
			},
			"updateUser": &graphql.Field{
				Type:        userType.User,
				Description: "Update an existing user",
				Args: graphql.FieldConfigArgument{
					"input": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(userType.User),
					},
				},
				Resolve: userResolver.UpdateUser,
			},
			"deleteUser": &graphql.Field{
				Type:        userType.Base,
				Description: "Delete a user",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: userResolver.DeleteUser,
			},
		},
	}
}
