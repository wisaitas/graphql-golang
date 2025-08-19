package app

import (
	"github.com/graphql-go/graphql"
	gqlTypes "github.com/wisaitas/graphql-golang/internal/app/graphql"
	"github.com/wisaitas/graphql-golang/internal/app/resolver"
	"github.com/wisaitas/graphql-golang/internal/app/service"
)

// CreateSchema สร้าง GraphQL schema
func CreateSchema() (graphql.Schema, error) {
	// สร้าง services
	userService := service.NewUserService()

	// สร้าง resolvers
	queryResolver := resolver.NewQueryResolver(userService)
	mutationResolver := resolver.NewMutationResolver(userService)

	// กำหนด Query type
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"hello": &graphql.Field{
				Type:        graphql.String,
				Description: "Simple hello world query",
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: queryResolver.Hello,
			},
			"users": &graphql.Field{
				Type:        graphql.NewList(gqlTypes.UserType),
				Description: "Get all users",
				Resolve:     queryResolver.Users,
			},
			"user": &graphql.Field{
				Type:        gqlTypes.UserType,
				Description: "Get user by ID",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: queryResolver.User,
			},
		},
	})

	// กำหนด Mutation type
	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createUser": &graphql.Field{
				Type:        gqlTypes.UserResponseType,
				Description: "Create a new user",
				Args: graphql.FieldConfigArgument{
					"input": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(gqlTypes.CreateUserInputType),
					},
				},
				Resolve: mutationResolver.CreateUser,
			},
			"updateUser": &graphql.Field{
				Type:        gqlTypes.UserResponseType,
				Description: "Update an existing user",
				Args: graphql.FieldConfigArgument{
					"input": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(gqlTypes.UpdateUserInputType),
					},
				},
				Resolve: mutationResolver.UpdateUser,
			},
			"deleteUser": &graphql.Field{
				Type:        gqlTypes.BaseResponseType,
				Description: "Delete a user",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: mutationResolver.DeleteUser,
			},
		},
	})

	// สร้าง schema
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	})

	return schema, err
}
