package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/ngrink/tasker/core"
	"github.com/spf13/cobra"
)

var uncompleteCmd = &cobra.Command{
	Use:   "uncomplete",
	Short: "Mark task as uncompleted",
	Long:  "Mark task as uncompleted",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}
		core.Tasks.Uncomplete(id)
		core.Tasks.Save()

		fmt.Printf("Task with id %d marked as uncompleted\n", id)
	},
}

func init() {
	rootCmd.AddCommand(uncompleteCmd)
}
