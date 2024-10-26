package renderers

import (
	"fmt"
	"todo/models"

	"github.com/jedib0t/go-pretty/v6/table"
)

/*
 * implements `Renderer`
 */
type TasksRenderer struct {
	Tasks []models.Task
}

func (t *TasksRenderer) Render() {
	tb := table.NewWriter()
	tb.SetStyle(table.StyleRounded)
	tb.AppendHeader(table.Row{"ID", "Project", "Name", "Status"})

	for _, task := range t.Tasks {
		tb.AppendRow(table.Row{
			task.ID,
			task.Project.Name,
			task.Name,
			task.Status.String(),
		})
	}

	fmt.Println(tb.Render())
}
