package init_project

import (
	"log"

	"github.com/maxguuse/bruh/internal/forms"
	"github.com/maxguuse/bruh/internal/fs"
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

	isExists := fs.IsExists("bruh.yaml")
	if isExists {
		log.Fatal("bruh.yaml file already exists")
	}

	err := settings.Write(stg)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Created bruh.yaml")

	err = fs.Mkdir(".", types.AppsDir, types.LibsDir)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Created directories: ", types.AppsDir, types.LibsDir)

	err = fs.GoWorkInit()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Initialized go workspace")
}
