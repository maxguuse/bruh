package forms

import (
	"github.com/charmbracelet/huh"
	"github.com/maxguuse/bruh/internal/types"
)

const (
	KeySubcommand = "subcommand"
)

func Welcome() *huh.Form {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[types.SubcommandType]().
				Title("Subcommand").
				Options(
					huh.NewOption("Init Project", types.InitProject),
					huh.NewOption("Create Module", types.CreateModule),
				).
				Key(KeySubcommand),
		),
	)

	return form
}
