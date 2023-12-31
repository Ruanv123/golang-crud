package config

import (
	"first-crud/prisma/db"

	"github.com/rs/zerolog/log"
)

func ConnectDb() (*db.PrismaClient, error) {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		return nil, err
	}

	log.Info().Msg("Connect do database!")
	return client, nil
}
