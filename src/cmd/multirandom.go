package cmd

import (
	"fmt"

	"github.com/agzuniverse/go-testcases/src/components"
	"github.com/agzuniverse/go-testcases/src/utils"
	"github.com/spf13/cobra"
)

var multiRandom = &cobra.Command{
	Use:   "multirandom",
	Short: "Generate n positive random integers",
	Run: func(cmd *cobra.Command, args []string) {
		v, err := components.GenMultiRandom(n, min, max)
		if err != nil {
			utils.HandleErr(err)
		}
		fmt.Println(v)
	},
}

var min int
var max int
var n int

func init() {
	multiRandom.Flags().IntVar(&n, "n", 0, "Number of random numbers to generate")
	multiRandom.Flags().IntVar(&min, "min", 0, "Minimum value of random number")
	multiRandom.Flags().IntVar(&max, "max", 0, "Maximum value of random number")
	rootCmd.AddCommand(multiRandom)

}
