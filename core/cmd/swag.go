package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(swagCmd)
}

var swagCmd = &cobra.Command{
	Use:   "swag:init",
	Short: "Generate swagger document",
	Long:  `This will parse your comments and generate the required files`,
	Run: func(cmd *cobra.Command, args []string) {
		swagInit()
	},
}

func swagInit() {
	out, err := exec.Command("swag", "init").Output()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))
}
