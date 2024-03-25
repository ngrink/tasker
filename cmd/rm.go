package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/ngrink/tasker/core"
	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove task",
	Long:  "Remove task",
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}

		core.Tasks.Remove(id)
		core.Tasks.Save()

		fmt.Printf("Removed task with id {%d}\n", id)
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
