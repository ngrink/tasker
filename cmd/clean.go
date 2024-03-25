package cmd

import (
	"github.com/ngrink/tasker/core"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Removes all tasks",
	Long:  "Removes all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		core.Tasks.Clean()
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}
