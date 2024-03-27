package cmd

import (
	"github.com/ngrink/tasker/core"
	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List tasks",
	Long:  "List tasks",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			core.Tasks.RenderByScope(args[0])
		} else {
			core.Tasks.RenderAll()
		}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
