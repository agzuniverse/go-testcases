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
		if err := components.ExecFile(name, time); err != nil {
			utils.HandleErr(err)
		}
	},
}

var name string
var time int

func init() {
	exec.Flags().StringVar(&name, "process-fail", "", "Process that should NOT finish execution within the given time.")
	exec.Flags().IntVar(&time, "time", 0, "Maximum time the program can run.")
	rootCmd.AddCommand(exec)
}
