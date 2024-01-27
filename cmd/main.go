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

const (
	AppsDir = "apps"
	LibsDir = "libs"
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
	mkdirCmd := exec.Command("mkdir", AppsDir, LibsDir)
	_, stderr := mkdirCmd.Output()
	if stderr != nil {
		log.Fatal(stderr)
	}

	log.Println("Created directories: ", AppsDir, LibsDir)

	goWorkInit := exec.Command("go", "work", "init")
	_, stderr = goWorkInit.Output()
	if stderr != nil {
		log.Fatal(stderr)
	}

	log.Println("Initialized go workspace")
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
		createAppCmd(moduleName)
	case Lib:
		log.Println("Lib: ", moduleName)
	}
}

func createAppCmd(appName string) {
	createAppFolderCmd := exec.Command("mkdir", appName)
	createAppFolderCmd.Dir = AppsDir
	_, stderr := createAppFolderCmd.Output()
	if stderr != nil {
		log.Fatal(stderr)
	}

	log.Println("Created app folder: ", appName)

	goModInitCmd := exec.Command("go", "mod", "init", appName)
	goModInitCmd.Dir = AppsDir + "/" + appName
	_, stderr = goModInitCmd.Output()
	if stderr != nil {
		log.Fatal(stderr)
	}

	log.Println("Initialized go module: ", appName)

	addModToWork := exec.Command("go", "work", "use", ".")
	addModToWork.Dir = AppsDir + "/" + appName
	_, stderr = addModToWork.Output()
	if stderr != nil {
		log.Fatal(stderr)
	}

	log.Println("Added module to go workspace: ", appName)

	createBaseFoldersCmd := exec.Command("mkdir", "cmd", "internal")
	createBaseFoldersCmd.Dir = AppsDir + "/" + appName
	_, stderr = createBaseFoldersCmd.Output()
	if stderr != nil {
		log.Fatal(stderr)
	}

	log.Println("Created base folders: ", "cmd", "internal")

	createMainFileCmd := exec.Command("touch", "main.go")
	createMainFileCmd.Dir = AppsDir + "/" + appName + "/cmd"
	_, stderr = createMainFileCmd.Output()
	if stderr != nil {
		log.Fatal(stderr)
	}

	log.Println("Created main file: ", "main.go")
}
