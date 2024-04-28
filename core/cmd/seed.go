package cmd

import (
	"app/database/seeder"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(seedCmd)
}

var seedCmd = &cobra.Command{
	Use:   "db:seeder",
	Short: "Insert seed data",
	Long:  `Insert seed data to tables`,
	Run: func(cmd *cobra.Command, args []string) {
		seeder.Seeder()
	},
}
