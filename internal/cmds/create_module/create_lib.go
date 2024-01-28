package create_module

import (
	"log"

	"github.com/maxguuse/bruh/internal/fs"
	"github.com/maxguuse/bruh/internal/settings"
	"github.com/maxguuse/bruh/internal/types"
)

func createLib(lib *types.Module, project *settings.ProjectDetails) {
	err := createGoModule(lib, project)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Created Go module: ", lib.Name)

	err = fs.Touch(types.LibsDir+"/"+lib.Name, "main.go")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Created main file: ", "main.go")
}
