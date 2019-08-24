package cmd

import (
	"github.com/agzuniverse/go-testcases/src/components"
	"github.com/agzuniverse/go-testcases/src/utils"
	"github.com/spf13/cobra"
)

var random = &cobra.Command{
	Use:   "random",
	Short: "Generate n random integers",
	Run: func(cmd *cobra.Command, args []string) {
		v, err := components.GenMultiRandom(np, minp, maxp)
		if err != nil {
			utils.HandleErr(err)
		}
		utils.PprintArray(v)
	},
}

var minp int
var maxp int
var np int

func init() {
	random.Flags().IntVar(&np, "n", 0, "Number of random numbers to generate")
	random.Flags().IntVar(&minp, "min", 0, "Minimum value of random number")
	random.Flags().IntVar(&maxp, "max", 0, "Maximum value of random number")
	rootCmd.AddCommand(random)

}
