package app

import appGraphqlType "github.com/wisaitas/graphql-golang/internal/app/graphqltype"

type graphqlType struct {
	UserType *appGraphqlType.UserType
}

func newGraphqlType() *graphqlType {
	return &graphqlType{
		UserType: appGraphqlType.NewUserType(),
	}
}
