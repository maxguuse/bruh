package create_module

import (
	"log"

	"github.com/maxguuse/bruh/internal/fs"
	"github.com/maxguuse/bruh/internal/settings"
	"github.com/maxguuse/bruh/internal/types"
)

func createApp(app *types.Module, project *settings.ProjectDetails) {
	root := types.AppsDir + "/" + app.Name

	err := createGoModule(app, project)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Created Go module: ", app.Name)

	err = fs.Mkdir(root, "cmd", "internal")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Created base folders: ", "cmd", "internal")

	err = fs.Touch(root+"/cmd", "main.go")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Created main file: ", "main.go")
}
