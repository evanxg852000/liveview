package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(validateCmd)
}

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validates the configartion file.",
	Long:  "Validates the configartion file.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Validating the config file")
	},
}
