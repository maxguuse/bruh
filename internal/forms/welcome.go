package forms

import (
	"log"

	"github.com/charmbracelet/huh"
	"github.com/maxguuse/bruh/internal/settings"
	"github.com/maxguuse/bruh/internal/types"
	"github.com/samber/lo"
)

const (
	KeySubcommand = "subcommand"
)

type Welcome struct {
	*huh.Form
}

func NewWelcome() *Welcome {
	_, err := settings.TryParse()

	options := []huh.Option[types.SubcommandType]{
		huh.NewOption("Init Project", types.InitProject),
		huh.NewOption("Create Module", types.CreateModule),
	}

	filteredOptions := lo.Filter(options, func(option huh.Option[types.SubcommandType], i int) bool {
		if err == nil && option.Value == types.InitProject {
			return false
		}

		return true
	})

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[types.SubcommandType]().
				Title("Subcommand").
				Options(filteredOptions...).
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
