package app

import appService "github.com/wisaitas/graphql-golang/internal/app/service"

type service struct {
	userService *appService.UserService
}

func newService() *service {
	return &service{
		userService: appService.NewUserService(),
	}
}
