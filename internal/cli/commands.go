package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var configFile string
var rootCmd = &cobra.Command{
	Use:   "diffuser",
	Short: "Diffuser is a very fast data processing platform",
	Long:  `A fast and flexible data processing platform for modern application`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("starting server ...")
		fmt.Println(configFile)
	},
}

func init() {
	rootCmd.Flags().StringVarP(&configFile, "config", "c", "", "Path to the config file.")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// -topology(validate, export dot-graph)
// -
