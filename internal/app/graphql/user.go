package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/wisaitas/graphql-golang/internal/app/model"
)

type UserType struct {
	Base *graphql.Object
	User *graphql.Object
}

func NewUserType() *UserType {
	return &UserType{
		Base: Base,
		User: graphql.NewObject(graphql.ObjectConfig{
			Name: "User",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.String,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"email": &graphql.Field{
					Type: graphql.String,
				},
				"age": &graphql.Field{
					Type: graphql.Int,
				},
				"created_at": &graphql.Field{
					Type: graphql.String,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						if user, ok := p.Source.(*model.User); ok {
							return user.CreatedAt.Format("2006-01-02T15:04:05Z07:00"), nil
						}
						return nil, nil
					},
				},
				"updated_at": &graphql.Field{
					Type: graphql.String,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						if user, ok := p.Source.(*model.User); ok {
							return user.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"), nil
						}
						return nil, nil
					},
				},
			},
		}),
	}
}

// // CreateUserInputType สำหรับ input การสร้าง user
// var CreateUserInputType = graphql.NewInputObject(graphql.InputObjectConfig{
// 	Name: "CreateUserInput",
// 	Fields: graphql.InputObjectConfigFieldMap{
// 		"name": &graphql.InputObjectFieldConfig{
// 			Type: graphql.NewNonNull(graphql.String),
// 		},
// 		"email": &graphql.InputObjectFieldConfig{
// 			Type: graphql.NewNonNull(graphql.String),
// 		},
// 		"age": &graphql.InputObjectFieldConfig{
// 			Type: graphql.Int,
// 		},
// 	},
// })

// // UpdateUserInputType สำหรับ input การอัพเดท user
// var UpdateUserInputType = graphql.NewInputObject(graphql.InputObjectConfig{
// 	Name: "UpdateUserInput",
// 	Fields: graphql.InputObjectConfigFieldMap{
// 		"id": &graphql.InputObjectFieldConfig{
// 			Type: graphql.NewNonNull(graphql.String),
// 		},
// 		"name": &graphql.InputObjectFieldConfig{
// 			Type: graphql.String,
// 		},
// 		"email": &graphql.InputObjectFieldConfig{
// 			Type: graphql.String,
// 		},
// 		"age": &graphql.InputObjectFieldConfig{
// 			Type: graphql.Int,
// 		},
// 	},
// })
