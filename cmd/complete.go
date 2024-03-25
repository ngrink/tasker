package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/ngrink/tasker/core"
	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Mark task as completed",
	Long:  "Mark task as completed",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}
		core.Tasks.Complete(id)
		core.Tasks.Save()

		fmt.Printf("Task with id %d marked as completed\n", id)
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
