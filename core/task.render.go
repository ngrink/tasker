package core

import (
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

func (t *TaskContainer) RenderAll() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "Done", "Description", "Category", "Age"})
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetColumnAlignment([]int{3, 1, 3, 3, 3})
	table.SetBorders(tablewriter.Border{Left: false, Right: false, Top: false, Bottom: false})

	fmt.Println()
	for idx, task := range t.Data {
		table.Append([]string{
			strconv.Itoa(idx),
			task.getStatusMark(),
			task.Description,
			task.Category,
			task.getAge(),
		})
	}

	table.Render()
	fmt.Println()
}
