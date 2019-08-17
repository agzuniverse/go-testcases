package cmd

import (
	"fmt"

	"github.com/agzuniverse/go-testcases/src/utils"

	"github.com/agzuniverse/go-testcases/src/components"

	"github.com/spf13/cobra"
)

var singleRandom = &cobra.Command{
	Use:   "singlerandom",
	Short: "Generate a single random positive integer between MIN and MAX",
	Run: func(cmd *cobra.Command, args []string) {
		v, err := components.GenSingleRandom(min, max)
		if err != nil {
			utils.HandleErr(err)
		}
		fmt.Println(v)
	},
}

var min int
var max int

func init() {
	singleRandom.Flags().IntVar(&min, "min", 0, "Minimum value of random number")
	singleRandom.Flags().IntVar(&max, "max", 0, "Maximum value of random number")
	rootCmd.AddCommand(singleRandom)
}
