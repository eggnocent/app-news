package seeds

import (
	"app-news/internal/core/domain/model"
	"app-news/lib/conv"

	"github.com/rs/zerolog/log"

	"gorm.io/gorm"
)

func SeedRoles(db *gorm.DB) {
	bytes, err := conv.HashPassword("admin123")

	if err != nil {
		log.Fatal().Err(err).Msg("Error creating password hash")
	}

	admin := model.User{
		Name:     "Admin",
		Email:    "admin@gmail.com",
		Password: string(bytes),
	}

	if err := db.FirstOrCreate(&admin, model.User{Email: "admin@gmail.com"}).Error; err != nil {
		log.Fatal().Err(err).Msg("Error creating admin user")
	} else {
		log.Info().Msgf("admin role successfully created")
	}
}
