package forms

import (
	"log"

	"github.com/charmbracelet/huh"
	"github.com/maxguuse/bruh/internal/types"
)

const (
	KeyProjectName  = "project_name"
	KeyProjectOwner = "project_owner"
)

type ProjectDetails struct {
	*huh.Form
}

func NewProjectDetails() *ProjectDetails {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Project Name:").
				CharLimit(50).
				Inline(true).
				Key(KeyProjectName),
			huh.NewInput().
				Title("Owner:").
				CharLimit(50).
				Inline(true).
				Key(KeyProjectOwner),
		),
	)

	return &ProjectDetails{
		form,
	}
}

func (p *ProjectDetails) Run() *types.ProjectDetails {
	err := p.Form.Run()
	if err != nil {
		log.Fatal(err)
	}

	return &types.ProjectDetails{
		Name:  p.Form.GetString(KeyProjectName),
		Owner: p.Form.GetString(KeyProjectOwner),
	}
}
