package core

import (
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

func (t *TaskContainer) RenderAll() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Id", "Description", "Category"})
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetBorders(tablewriter.Border{Left: false, Right: false, Top: false, Bottom: false})
	table.SetTablePadding("\t")

	fmt.Println()
	for id, task := range Tasks.Data {
		table.Append([]string{
			strconv.Itoa(id),
			task.Description,
			task.Category,
		})
	}
	table.Render()
	fmt.Println()
}

func (t *TaskContainer) Render() {

}

// func formatDeadline(deadline time.Time) string {
// 	var deadlineStr string

// 	if deadline.IsZero() {
// 		deadlineStr = ""
// 	} else {
// 		deadlineStr = deadline.Format(time.UnixDate)
// 	}

// 	return deadlineStr
// }
