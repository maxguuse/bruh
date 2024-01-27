package main

import (
	"log"
	"os/exec"

	"github.com/charmbracelet/huh"
)

type ModuleType int

const (
	App ModuleType = iota
	Lib
)

type SubcommandType int

const (
	InitProject SubcommandType = iota
	CreateModule
)

func main() {
	subcommand := welcomeFormCmd()

	switch subcommand {
	case InitProject:
		initProjectCmd()
	case CreateModule:
		createModuleCmd()
	}
}

func welcomeFormCmd() (subcommand SubcommandType) {
	err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[SubcommandType]().
				Title("Subcommand").
				Options(
					huh.NewOption("Init Project", InitProject),
					huh.NewOption("Create Module", CreateModule),
				).
				Value(&subcommand),
		),
	).Run()
	if err != nil {
		log.Fatal(err)
	}

	return
}

func initProjectCmd() {
	appsDir := "apps"
	libsDir := "libs"

	mkdirCmd := exec.Command("mkdir", appsDir, libsDir)

	_, stderr := mkdirCmd.Output()
	if stderr != nil {
		log.Fatal(stderr)
	}

	log.Println("Created directories: ", appsDir, libsDir)

	welcomeFormCmd()
}

func createModuleCmd() {
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
