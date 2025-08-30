package resolver

import (
	"github.com/wisaitas/graphql-golang/internal/appv2/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserService service.UserService
}

func NewResolver(userService service.UserService) *Resolver {
	return &Resolver{
		UserService: userService,
	}
}
