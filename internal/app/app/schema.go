package app

import (
	"github.com/graphql-go/graphql"
	appSchema "github.com/wisaitas/graphql-golang/internal/app/schema"
)

func newSchema(resolver *resolver, graphqlType *graphqlType) graphql.Schema {
	userSchema := appSchema.NewUserSchema(resolver.userResolver, graphqlType.UserType)

	queryFields := graphql.Fields{}

	for name, field := range userSchema.Queries {
		queryFields[name] = field
	}

	mutationFields := graphql.Fields{}

	for name, field := range userSchema.Mutations {
		mutationFields[name] = field
	}

	rootQuery := graphql.ObjectConfig{
		Name:   "Query",
		Fields: queryFields,
	}

	rootMutation := graphql.ObjectConfig{
		Name:   "Mutation",
		Fields: mutationFields,
	}

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    graphql.NewObject(rootQuery),
		Mutation: graphql.NewObject(rootMutation),
	})

	if err != nil {
		panic(err)
	}

	return schema
}
