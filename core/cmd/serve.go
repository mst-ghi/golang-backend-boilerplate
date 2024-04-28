package cmd

import (
	"app/core/bootstrap"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve app on dev server",
	Long:  `App will be served on host & port defined in config/config.yaml file`,
	Run: func(cmd *cobra.Command, args []string) {
		bootstrap.Serve()
	},
}
