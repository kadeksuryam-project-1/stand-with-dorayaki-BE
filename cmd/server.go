package cmd

import (
	"backend/config"
	"backend/internal/handler"
	"backend/pkg/lib"
	svr "backend/pkg/server"

	"github.com/spf13/cobra"
)

func server() *cobra.Command {
	serverCmd := &cobra.Command{
		Use:   "server",
		Short: "Run server",
		Run: func(cmd *cobra.Command, args []string) {
			lib.DatabaseInit(config.C.DatabaseURL)
			lib.GStorageInit(config.C.ProjectID, config.C.BucketName)

			appServer := svr.NewServer()
			handler.RoutesInit(appServer.App(), lib.DB)
			appServer.Start(config.C.Port)
		},
	}
	return serverCmd
}
