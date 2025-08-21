package app

import gqlTypes "github.com/wisaitas/graphql-golang/internal/app/graphql"

type Type struct {
	UserType *gqlTypes.UserType
}

func newType() *Type {
	return &Type{
		UserType: gqlTypes.NewUserType(),
	}
}
