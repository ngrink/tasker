package cmd

import (
	"fmt"

	"github.com/ngrink/tasker/core"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add task to the list",
	Long:  "Add task to the list",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		category := ""
		if len(args) == 2 {
			category = args[1]
		}

		task := core.NewTask(core.CreateTaskDto{
			Description: args[0],
			Category:    category,
		})

		core.Tasks.Add(task)
		core.Tasks.Save()

		fmt.Println("Task added")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
