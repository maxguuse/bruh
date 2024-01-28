package forms

import (
	"log"

	"github.com/charmbracelet/huh"
	"github.com/maxguuse/bruh/internal/types"
)

const (
	KeyModuleType = "module_type"
	KeyModuleName = "module_name"
)

type ModuleInfo struct {
	*huh.Form
}

func NewModuleInfo() *ModuleInfo {
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

	return &ModuleInfo{
		form,
	}
}

func (m *ModuleInfo) Run() *types.Module {
	err := m.Form.Run()
	if err != nil {
		log.Fatal(err)
	}

	return &types.Module{
		Name: m.Form.GetString(KeyModuleName),
		Type: m.Form.Get(KeyModuleType).(types.ModuleType),
	}
}
