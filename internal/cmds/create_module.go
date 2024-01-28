package cmds

import (
	"log"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/maxguuse/bruh/internal/types"
	"gopkg.in/yaml.v2"
)

func CreateModule() {
	var module *types.Module

	err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[types.ModuleType]().
				Title("Module Type").
				Options(
					huh.NewOption("Application", types.App),
					huh.NewOption("Library", types.Lib),
				).
				Value(&module.Type),
			huh.NewInput().
				Title("Module Name").
				CharLimit(20).
				Value(&module.Name),
		),
	).Run()
	if err != nil {
		log.Fatal(err)
	}

	cfg := &types.Config{}

	blob, err := os.ReadFile("bruh.yaml")
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(blob, cfg)
	if err != nil {
		log.Fatal(err)
	}

	project := cfg.Project

	switch module.Type {
	case types.App:
		createApp(module, project)
	case types.Lib:
		createLib(module, project)
	}
}
