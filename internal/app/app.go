package app

import (
	"app-news/config"

	"github.com/rs/zerolog/log"
)

func RunServer() {
	cfg := config.NewConfig()
	_, err := cfg.ConnectionPostgres()
	if err != nil {
		log.Fatal().Msgf("Error connecting to database: %v", err)
		return
	}
}
