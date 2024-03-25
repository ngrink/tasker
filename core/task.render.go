package core

import (
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

func (t *TaskContainer) RenderAll() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Id", "Done", "Description", "Category", "Age"})
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetColumnAlignment([]int{3, 1, 3, 3, 3})
	table.SetBorders(tablewriter.Border{Left: false, Right: false, Top: false, Bottom: false})

	fmt.Println()
	for id, task := range Tasks.Data {
		if !task.IsDone {
			continue
		}

		table.Append([]string{
			strconv.Itoa(id),
			task.getStatusMark(),
			task.Description,
			task.Category,
			task.getAge(),
		})
	}

	for id, task := range Tasks.Data {
		if task.IsDone {
			continue
		}

		table.Append([]string{
			strconv.Itoa(id),
			task.getStatusMark(),
			task.Description,
			task.Category,
			task.getAge(),
		})
	}

	table.Render()
	fmt.Println()
}
