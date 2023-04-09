package lib

import (
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit(url string) {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	log.Info().Msg("database connected")
	DB = db
}
