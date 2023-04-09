package cmd

import "github.com/spf13/cobra"

func Init() {
	mainCommand := &cobra.Command{
		Use:   "swd-backend",
		Short: "Stand for Dorayaki backend service",
	}

	mainCommand.AddCommand(
		migrate(),
		server(),
	)

	if err := mainCommand.Execute(); err != nil {
		panic(err)
	}
}
