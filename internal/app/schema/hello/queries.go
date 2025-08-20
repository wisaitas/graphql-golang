package hello

import (
	"github.com/graphql-go/graphql"
)

// GetQueries ส่งกลับ hello queries ทั้งหมด
func GetQueries() graphql.Fields {
	return graphql.Fields{
		"hello": &graphql.Field{
			Type:        graphql.String,
			Description: "Simple hello world query",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: HelloResolver,
		},
	}
}
