package cmds

import (
	"log"
	"os"

	"github.com/maxguuse/bruh/internal/forms"
	"github.com/maxguuse/bruh/internal/types"
	"gopkg.in/yaml.v2"
)

func CreateModule() {
	module := forms.NewModuleInfo().Run()

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
