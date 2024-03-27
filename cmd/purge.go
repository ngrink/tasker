package cmd

import (
	"fmt"

	"github.com/ngrink/tasker/core"
	"github.com/spf13/cobra"
)

var purgeCmd = &cobra.Command{
	Use:   "purge",
	Short: "Remove all completed tasks",
	Long:  "Remove all completed tasks",
	Run: func(cmd *cobra.Command, args []string) {
		core.Tasks.Purge()
		core.Tasks.Save()

		fmt.Println("All completed tasks removed")
	},
}

func init() {
	rootCmd.AddCommand(purgeCmd)
}
