package initial

import (
	"github.com/wisaitas/graphql-golang/internal/appv2/resolver"
)

func NewInitial() *resolver.Resolver {
	config := NewConfig()
	repository := NewRepository(config)
	service := NewService(repository)
	return resolver.NewResolver(
		service.userService,
	)
}
