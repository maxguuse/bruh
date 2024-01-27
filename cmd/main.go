package main

import (
	"log"

	"github.com/charmbracelet/huh"
)

type ModuleType int

const (
	App ModuleType = iota
	Lib
)

func main() {
	var moduleType ModuleType
	var moduleName string

	err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[ModuleType]().
				Title("Module Type").
				Options(
					huh.NewOption("Application", App),
					huh.NewOption("Library", Lib),
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

	switch moduleType {
	case App:
		log.Println("App: ", moduleName)
	case Lib:
		log.Println("Lib: ", moduleName)
	}
}
