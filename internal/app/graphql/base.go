package graphql

import "github.com/graphql-go/graphql"

var Base = graphql.NewObject(graphql.ObjectConfig{
	Name: "BaseResponse",
	Fields: graphql.Fields{
		"success": &graphql.Field{
			Type: graphql.Boolean,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
	},
})
