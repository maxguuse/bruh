package forms

import (
	"log"

	"github.com/charmbracelet/huh"
	"github.com/maxguuse/bruh/internal/settings"
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

func (p *ProjectDetails) Run() *settings.ProjectDetails {
	err := p.Form.Run()
	if err != nil {
		log.Fatal(err)
	}

	return &settings.ProjectDetails{
		Name:  p.Form.GetString(KeyProjectName),
		Owner: p.Form.GetString(KeyProjectOwner),
	}
}
