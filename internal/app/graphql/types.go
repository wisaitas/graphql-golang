package graphql

import (
	"github.com/graphql-go/graphql"
)

// UserType GraphQL type สำหรับ User
var UserType = graphql.NewObject(graphql.ObjectConfig{
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
		},
		"updated_at": &graphql.Field{
			Type: graphql.String,
		},
	},
})

// CreateUserInputType สำหรับ input การสร้าง user
var CreateUserInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "CreateUserInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"name": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"email": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"age": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
	},
})

// UpdateUserInputType สำหรับ input การอัพเดท user
var UpdateUserInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "UpdateUserInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"name": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"email": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"age": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
	},
})

// BaseResponseType สำหรับ response ทั่วไป
var BaseResponseType = graphql.NewObject(graphql.ObjectConfig{
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

// UserResponseType สำหรับ response ที่มี user data
var UserResponseType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserResponse",
	Fields: graphql.Fields{
		"success": &graphql.Field{
			Type: graphql.Boolean,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"user": &graphql.Field{
			Type: UserType,
		},
	},
})
