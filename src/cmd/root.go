package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "apollo",
	Short: "Apollo is a utility to generate testcases for online judges",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Root command ran")
	},
}

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
