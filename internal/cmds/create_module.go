package cmds

import (
	"log"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/maxguuse/bruh/internal/types"
	"gopkg.in/yaml.v2"
)

func CreateModule() {
	var moduleType types.ModuleType
	var moduleName string

	err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[types.ModuleType]().
				Title("Module Type").
				Options(
					huh.NewOption("Application", types.App),
					huh.NewOption("Library", types.Lib),
				).
				Value(&moduleType),
			huh.NewInput().
				Title("Module Name").
				CharLimit(20).
				Value(&moduleName),
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

	switch moduleType {
	case types.App:
		createApp(moduleName, project)
	case types.Lib:
		createLib(moduleName, project)
	}
}
