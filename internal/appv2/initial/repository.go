package initial

import (
	repositoryApp "github.com/wisaitas/graphql-golang/internal/appv2/repository"
)

type repository struct {
	userRepository repositoryApp.UserRepository
}

func NewRepository(
	config *config,
) *repository {
	return &repository{
		userRepository: repositoryApp.NewUserRepository(config.database),
	}
}
