package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/ngrink/tasker/core"
	"github.com/ngrink/tasker/lib"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit <idx> <task>",
	Short: "Edit a task",
	Long:  "Edit a task",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		idx, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}
		description := args[1]

		core.Tasks.Edit(idx, description)
		core.Tasks.Save()

		fmt.Println("Task changed")
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
	editCmd.Example = lib.GenerateExample([]string{
		`tasker edit 7 "Another task description"`,
	})
}
