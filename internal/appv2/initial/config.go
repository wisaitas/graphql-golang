package initial

import (
	configApp "github.com/wisaitas/graphql-golang/internal/appv2/config"
	"gorm.io/gorm"
)

type config struct {
	database *gorm.DB
}

func NewConfig() *config {
	return &config{
		database: configApp.ConnectDatabase(),
	}
}
