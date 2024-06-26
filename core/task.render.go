package core

import (
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

func NewTable() *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "Description", "Scope", "Age"})
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetColumnAlignment([]int{3, 3, 3, 3})
	table.SetColMinWidth(1, 40)
	table.SetAutoWrapText(false)
	table.SetBorders(tablewriter.Border{Left: false, Right: false, Top: false, Bottom: false})

	return table
}

func (t *TaskContainer) RenderAll() {
	table := NewTable()

	fmt.Println()
	for idx, task := range t.Data {
		table.Append([]string{
			strconv.Itoa(idx),
			task.GetDesriptionWithStatus(),
			task.Scope,
			task.GetAge(),
		})
	}

	table.Render()
	fmt.Println()
}

func (t *TaskContainer) RenderByScope(scope string) {
	table := NewTable()

	fmt.Println()
	for idx, task := range t.Data {
		if task.Scope != scope {
			continue
		}
		table.Append([]string{
			strconv.Itoa(idx),
			task.GetStatusMark(),
			task.Description,
			task.Scope,
			task.GetAge(),
		})
	}

	table.Render()
	fmt.Println()
}
