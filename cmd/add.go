package cmd

import (
	"fmt"

	"github.com/ngrink/tasker/core"
	"github.com/ngrink/tasker/lib"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add <task> [<scope>]",
	Short: "Add task to the list",
	Long:  "Add task to the list",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		Scope := ""
		if len(args) == 2 {
			Scope = args[1]
		}

		task := core.NewTask(core.CreateTaskDto{
			Description: args[0],
			Scope:       Scope,
		})

		core.Tasks.Add(task)
		core.Tasks.Save()

		fmt.Println("Task added")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Example = lib.GenerateExample([]string{
		`tasker add "Some task"`,
		`tasker add "Task with scope" dev`,
	})
}
