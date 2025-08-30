package config

import (
	"fmt"
	"log"

	"github.com/wisaitas/graphql-golang/internal/appv2"
	"github.com/wisaitas/graphql-golang/internal/appv2/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDatabase() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable timezone=Asia/Bangkok",
		appv2.ENV.Database.Host,
		appv2.ENV.Database.Port,
		appv2.ENV.Database.User,
		appv2.ENV.Database.Password,
		appv2.ENV.Database.DBName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	if err := db.AutoMigrate(&entity.User{}); err != nil {
		log.Fatal("failed to migrate database:", err)
	}

	log.Println("database connected successfully!")
	return db
}
