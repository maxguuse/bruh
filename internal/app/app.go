package app

import (
	"log"

	"github.com/maxguuse/bruh/internal/cmds/create_module"
	"github.com/maxguuse/bruh/internal/cmds/init_project"
	"github.com/maxguuse/bruh/internal/forms"
	"github.com/maxguuse/bruh/internal/types"
)

var subcommands = map[types.SubcommandType]func(){
	types.InitProject:  init_project.Cmd,
	types.CreateModule: create_module.Cmd,
}

func Run() {
	subcommand := forms.NewWelcome().Run()

	handler, ok := subcommands[subcommand]
	if !ok {
		log.Fatal("Invalid subcommand")
	}

	handler()
}
