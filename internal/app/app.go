package app

import (
	"github.com/maxguuse/bruh/internal/cmds/create_module"
	"github.com/maxguuse/bruh/internal/cmds/init_project"
	"github.com/maxguuse/bruh/internal/forms"
	"github.com/maxguuse/bruh/internal/types"
)

func Run() {
	subcommand := forms.NewWelcome().Run()

	switch subcommand {
	case types.InitProject:
		init_project.Cmd()
	case types.CreateModule:
		create_module.Cmd()
	}
}
