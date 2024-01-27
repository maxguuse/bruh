package main

import (
	"log"
	"os/exec"

	"github.com/charmbracelet/huh"
	"github.com/samber/lo"
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
		createLibCmd(moduleName)
	}
}

func createAppCmd(appName string) {
	err := createGoModule(appName, App)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Created Go module: ", appName)

	createBaseFoldersCmd := exec.Command("mkdir", "cmd", "internal")
	createBaseFoldersCmd.Dir = AppsDir + "/" + appName
	_, stderr := createBaseFoldersCmd.Output()
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

func createLibCmd(libName string) {
	err := createGoModule(libName, Lib)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Created Go module: ", libName)

	createMainFileCmd := exec.Command("touch", "main.go")
	createMainFileCmd.Dir = LibsDir + "/" + libName
	_, stderr := createMainFileCmd.Output()
	if stderr != nil {
		log.Fatal(stderr)
	}

	log.Println("Created main file: ", "main.go")
}

func createGoModule(moduleName string, moduleType ModuleType) error {
	err := createModuleFolder(moduleName, moduleType)
	if err != nil {
		return err
	}

	err = initGoModule(moduleName, moduleType)
	if err != nil {
		return err
	}

	err = addGoModuleToWorkspace(moduleName, moduleType)
	if err != nil {
		return err
	}

	return nil
}

func createModuleFolder(moduleName string, moduleType ModuleType) error {
	rootDir := lo.If(moduleType == App, AppsDir).Else(LibsDir)

	createAppFolderCmd := exec.Command("mkdir", moduleName)
	createAppFolderCmd.Dir = rootDir
	_, err := createAppFolderCmd.Output()
	if err != nil {
		return err
	}

	return nil
}

func initGoModule(moduleName string, moduleType ModuleType) error {
	rootDir := lo.If(moduleType == App, AppsDir).Else(LibsDir)

	goModInitCmd := exec.Command("go", "mod", "init", moduleName)
	goModInitCmd.Dir = rootDir + "/" + moduleName
	_, err := goModInitCmd.Output()
	if err != nil {
		return err
	}

	return nil
}

func addGoModuleToWorkspace(moduleName string, moduleType ModuleType) error {
	rootDir := lo.If(moduleType == App, AppsDir).Else(LibsDir)

	addModToWork := exec.Command("go", "work", "use", ".")
	addModToWork.Dir = rootDir + "/" + moduleName
	_, err := addModToWork.Output()
	if err != nil {
		return err
	}

	return nil
}
