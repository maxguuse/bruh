package forms

import (
	"github.com/charmbracelet/huh"
	"github.com/maxguuse/bruh/internal/types"
)

const (
	KeyModuleType = "module_type"
	KeyModuleName = "module_name"
)

func AskForModuleInfo() *huh.Form {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[types.ModuleType]().
				Title("Module Type").
				Options(
					huh.NewOption("Application", types.App),
					huh.NewOption("Library", types.Lib),
				).
				Key(KeyModuleType),
			huh.NewInput().
				Title("Module Name").
				CharLimit(20).
				Key(KeyModuleName),
		),
	)

	return form
}
