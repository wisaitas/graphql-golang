package app

import appResolver "github.com/wisaitas/graphql-golang/internal/app/resolver"

type resolver struct {
	userResolver appResolver.UserResolver
}

func newResolver(service *service) *resolver {
	return &resolver{
		userResolver: appResolver.NewUserResolver(service.userService),
	}
}
