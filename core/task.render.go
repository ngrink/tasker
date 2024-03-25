package core

import (
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

func (t *TaskContainer) RenderAll() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Id", "Description", "Category", "Age"})
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
			task.getAge(),
		})
	}
	table.Render()
	fmt.Println()
}
