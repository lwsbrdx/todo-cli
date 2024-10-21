package helpers

import (
	"fmt"
	"todo/models"

	"github.com/jedib0t/go-pretty/v6/table"
)

type TasksRenderer struct {
	Tasks []models.Task
}

func (t *TasksRenderer) Render() {
	tb := table.NewWriter()

	tb.SetStyle(table.StyleRounded)

	tb.AppendHeader(table.Row{"ID", "Name", "Status"})

	for _, task := range t.Tasks {
		tb.AppendRow(table.Row{
			task.ID,
			task.Name,
			task.Status.String(),
		})
	}

	fmt.Println(tb.Render())
}
