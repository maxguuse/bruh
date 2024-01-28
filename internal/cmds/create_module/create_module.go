package create_module

import (
	"log"

	"github.com/maxguuse/bruh/internal/forms"
	"github.com/maxguuse/bruh/internal/settings"
	"github.com/maxguuse/bruh/internal/types"
)

func Cmd() {
	stg, err := settings.TryParse()
	if err != nil {
		log.Fatal(err)
	}

	module := forms.NewModuleInfo().Run()

	switch module.Type {
	case types.App:
		createApp(module, stg.Project)
	case types.Lib:
		createLib(module, stg.Project)
	}
}
