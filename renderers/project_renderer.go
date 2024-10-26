package renderers

import (
	"fmt"
	"todo/models"

	"github.com/jedib0t/go-pretty/v6/table"
)

/*
 * implements `Renderer`
 */
type ProjectsRenderer struct {
	Projects []models.Project
}

func (p *ProjectsRenderer) Render() {
	tb := table.NewWriter()
	tb.SetStyle(table.StyleRounded)
	tb.AppendHeader(table.Row{"ID", "Name"})

	for _, project := range p.Projects {
		tb.AppendRow(table.Row{
			project.ID,
			project.Name,
		})
	}

	fmt.Println(tb.Render())
}
