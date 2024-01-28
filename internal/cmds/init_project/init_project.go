package init_project

import (
	"log"
	"os"
	"os/exec"

	"github.com/maxguuse/bruh/internal/forms"
	"github.com/maxguuse/bruh/internal/settings"
	"github.com/maxguuse/bruh/internal/types"
)

func Cmd() {
	project := forms.NewProjectDetails().Run()
	if project.Name == "" || project.Owner == "" {
		log.Fatal("Invalid project details")
	}
	log.Println("Project Details: ", project)

	stg := &settings.Settings{
		Project: project,
	}

	if _, err := os.Stat(settings.SettingsFile); !os.IsNotExist(err) {
		log.Fatal("bruh.yaml already exists.")
	}

	err := settings.Write(stg)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Created bruh.yaml")

	mkdirCmd := exec.Command("mkdir", types.AppsDir, types.LibsDir)
	_, stderr := mkdirCmd.Output()
	if stderr != nil {
		log.Fatal(stderr)
	}

	log.Println("Created directories: ", types.AppsDir, types.LibsDir)

	goWorkInit := exec.Command("go", "work", "init")
	_, stderr = goWorkInit.Output()
	if stderr != nil {
		log.Fatal(stderr)
	}

	log.Println("Initialized go workspace")
}
