package cmd

import (
	"github.com/agzuniverse/go-testcases/src/components"
	"github.com/agzuniverse/go-testcases/src/utils"
	"github.com/spf13/cobra"
)

var exec = &cobra.Command{
	Use:   "exec",
	Short: "Execute a program and feed input to stdin",
	Run: func(cmd *cobra.Command, args []string) {
		if err := components.ExecFile(name, time, min, max, n); err != nil {
			utils.HandleErr(err)
		}
	},
}

var name string
var time int
var min int
var max int
var n int

func init() {
	exec.Flags().StringVar(&name, "process-fail", "", "Process that should NOT finish execution within the given time.")
	exec.Flags().IntVar(&time, "time", 0, "Maximum time the program can run.")
	exec.Flags().IntVar(&n, "n", 0, "Number of random numbers to generate")
	exec.Flags().IntVar(&min, "min", 0, "Minimum value of random number")
	exec.Flags().IntVar(&max, "max", 0, "Maximum value of random number")
	rootCmd.AddCommand(exec)
}
