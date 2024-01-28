package app

import (
	"github.com/maxguuse/bruh/internal/cmds"
	"github.com/maxguuse/bruh/internal/forms"
	"github.com/maxguuse/bruh/internal/types"
)

func Run() {
	subcommand := forms.NewWelcome().Run()

	switch subcommand {
	case types.InitProject:
		cmds.InitProject()
	case types.CreateModule:
		cmds.CreateModule()
	}
}
