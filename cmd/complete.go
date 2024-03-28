package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/ngrink/tasker/core"
	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete <idx>",
	Short: "Mark task as completed",
	Long:  "Mark task as completed",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		idx, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}
		core.Tasks.Complete(idx)
		core.Tasks.Save()

		fmt.Printf("Task with number %d marked as completed\n", idx)
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
