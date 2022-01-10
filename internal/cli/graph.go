package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(graphCmd)
}

var graphCmd = &cobra.Command{
	Use:   "graph",
	Short: "Generates the topology graph.",
	Long:  "Generates the topology graph.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Generating the topology graph")
	},
}
