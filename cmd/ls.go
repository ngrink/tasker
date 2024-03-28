package cmd

import (
	"github.com/ngrink/tasker/core"
	"github.com/ngrink/tasker/lib"
	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:   "ls [<scope>]",
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
	lsCmd.Example = lib.GenerateExample([]string{
		`tasker ls`,
		`tasker ls dev`,
	})
}
