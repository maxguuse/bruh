package forms

import (
	"github.com/charmbracelet/huh"
)

const (
	KeyProjectName  = "project_name"
	KeyProjectOwner = "project_owner"
)

func AskForProjectDetails() *huh.Form {
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

	return form
}
