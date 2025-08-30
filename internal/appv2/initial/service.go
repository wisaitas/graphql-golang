package initial

import serviceApp "github.com/wisaitas/graphql-golang/internal/appv2/service"

type service struct {
	userService serviceApp.UserService
}

func NewService(
	repository *repository,
) *service {
	return &service{userService: serviceApp.NewUserService(repository.userRepository)}
}
