package create_module

import (
	"log"
	"os"

	"github.com/maxguuse/bruh/internal/forms"
	"github.com/maxguuse/bruh/internal/settings"
	"github.com/maxguuse/bruh/internal/types"
	"gopkg.in/yaml.v2"
)

func Cmd() {
	module := forms.NewModuleInfo().Run()

	cfg := &settings.Settings{}

	blob, err := os.ReadFile("bruh.yaml")
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(blob, cfg)
	if err != nil {
		log.Fatal(err)
	}

	switch module.Type {
	case types.App:
		createApp(module, cfg.Project)
	case types.Lib:
		createLib(module, cfg.Project)
	}
}
