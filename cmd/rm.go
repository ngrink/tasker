package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/ngrink/tasker/core"
	"github.com/ngrink/tasker/lib"
	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove task",
	Long:  "Remove task",
	Run: func(cmd *cobra.Command, args []string) {
		idx, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}

		description := core.Tasks.Data[idx].Description
		action := fmt.Sprintf("Remove task \"%s\"", description)

		if ok := lib.ConfirmAction(action); !ok {
			fmt.Println("Action rejected")
			return
		}
		core.Tasks.Remove(idx)
		core.Tasks.Save()

		fmt.Printf("Removed task with number {%d}\n", idx)
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
