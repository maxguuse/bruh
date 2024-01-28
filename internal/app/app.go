package app

import (
	"log"
	"os"
	"os/exec"

	"github.com/charmbracelet/huh"
	"github.com/samber/lo"
	"gopkg.in/yaml.v2"
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

func Run() {
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

type Config struct {
	Project ProjectDetails `yaml:"project"`
}

func initProjectCmd() {
	project, err := askProjectDetailsForm()
	if err != nil || project.Name == "" || project.Owner == "" {
		log.Fatal("Invalid project details")
	}
	log.Println("Project Details: ", project)

	cfg := Config{
		Project: *project,
	}

	file := "bruh.yaml"

	if _, err := os.Stat(file); !os.IsNotExist(err) {
		log.Fatal("bruh.yaml already exists.")
	}

	blob, err := yaml.Marshal(cfg)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(file, blob, 0644)
	if err != nil {
		log.Fatal(err)
	}

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

type ProjectDetails struct {
	Name  string `yaml:"name"`
	Owner string `yaml:"owner"`
}

func askProjectDetailsForm() (*ProjectDetails, error) {
	var project ProjectDetails

	err := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Project Name:").
				CharLimit(50).
				Inline(true).
				Value(&project.Name),
			huh.NewInput().
				Title("Owner:").
				CharLimit(50).
				Inline(true).
				Value(&project.Owner),
		),
	).Run()

	if err != nil {
		return nil, err
	}

	return &project, nil
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

	cfg := &Config{}

	blob, err := os.ReadFile("bruh.yaml")
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(blob, cfg)
	if err != nil {
		log.Fatal(err)
	}

	project := cfg.Project

	switch moduleType {
	case App:
		createAppCmd(moduleName, project)
	case Lib:
		createLibCmd(moduleName, project)
	}
}

func createAppCmd(appName string, project ProjectDetails) {
	err := createGoModule(appName, App, project)
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

func createLibCmd(libName string, project ProjectDetails) {
	err := createGoModule(libName, Lib, project)
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

func createGoModule(moduleName string, moduleType ModuleType, project ProjectDetails) error {
	err := createModuleFolder(moduleName, moduleType)
	if err != nil {
		return err
	}

	err = initGoModule(moduleName, moduleType, project)
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

func initGoModule(moduleName string, moduleType ModuleType, project ProjectDetails) error {
	rootDir := lo.If(moduleType == App, AppsDir).Else(LibsDir)

	module := "github.com/" + project.Owner + "/" + project.Name + "/" + rootDir + "/" + moduleName

	goModInitCmd := exec.Command("go", "mod", "init", module)
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
