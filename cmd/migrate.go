package cmd

import (
	"backend/config"
	"backend/internal/schema"
	"backend/pkg/lib"

	"backend/cmd/seed"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

func migrate() *cobra.Command {
	migrateCmd := &cobra.Command{
		Use:   "migrate",
		Short: "Migrate database",
		Run: func(cmd *cobra.Command, args []string) {
			lib.DatabaseInit(config.C.DatabaseURL)

			log.Info().Msg("starting database migration")

			dropAllTables(lib.DB)
			err := lib.DB.AutoMigrate(
				&schema.Dorayaki{},
				&schema.Store{},
				&schema.DorayakiStoreStock{},
			)
			seed.SeedDorayaki(lib.DB)
			seed.SeedStore(lib.DB)
			seed.SeedDorayakiStoreStock(lib.DB)

			if err != nil {
				log.Error().Err(err).Msg("database migration failed")
				return
			}

			log.Info().Msg("database migrated succesfully")
		},
	}

	return migrateCmd
}

func dropAllTables(db *gorm.DB) error {
	migrator := db.Migrator()

	tables, _ := migrator.GetTables()
	for _, table := range tables {
		if err := migrator.DropTable(table); err != nil {
			return err
		}
	}
	return nil
}
