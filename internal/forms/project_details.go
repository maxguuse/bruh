package forms

import (
	"github.com/charmbracelet/huh"
)

func AskForProjectDetails() *huh.Form {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Project Name:").
				CharLimit(50).
				Inline(true).
				Key("project_name"),
			huh.NewInput().
				Title("Owner:").
				CharLimit(50).
				Inline(true).
				Key("project_owner"),
		),
	)

	return form
}
