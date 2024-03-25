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
		idx, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}
		core.Tasks.Uncomplete(idx)
		core.Tasks.Save()

		fmt.Printf("Task with number %d marked as uncompleted\n", idx)
	},
}

func init() {
	rootCmd.AddCommand(uncompleteCmd)
}
