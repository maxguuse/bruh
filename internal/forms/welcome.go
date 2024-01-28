package forms

import (
	"log"

	"github.com/charmbracelet/huh"
	"github.com/maxguuse/bruh/internal/types"
)

const (
	KeySubcommand = "subcommand"
)

type Welcome struct {
	*huh.Form
}

func NewWelcome() *Welcome {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[types.SubcommandType]().
				Title("Subcommand").
				Options(
					huh.NewOption("Init Project", types.InitProject),
					huh.NewOption("Create Module", types.CreateModule),
				).
				Key(KeySubcommand),
		),
	)

	return &Welcome{
		form,
	}
}

func (w *Welcome) Run() types.SubcommandType {
	err := w.Form.Run()
	if err != nil {
		log.Fatal(err)
	}

	subcommand, ok := w.Form.Get(KeySubcommand).(types.SubcommandType)
	if !ok {
		log.Fatal("Invalid subcommand type")
	}

	return subcommand
}
