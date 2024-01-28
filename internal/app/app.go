package app

import (
	"log"

	"github.com/maxguuse/bruh/internal/cmds"
	"github.com/maxguuse/bruh/internal/forms"
	"github.com/maxguuse/bruh/internal/types"
)

func Run() {
	subcommand := runWelcomeForm()

	switch subcommand {
	case types.InitProject:
		cmds.InitProject()
	case types.CreateModule:
		cmds.CreateModule()
	}
}

func runWelcomeForm() (subcommand types.SubcommandType) {
	form := forms.Welcome()

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	subcommand, ok := form.Get(forms.KeySubcommand).(types.SubcommandType)
	if !ok {
		log.Fatal("Invalid subcommand type")
	}

	return
}
