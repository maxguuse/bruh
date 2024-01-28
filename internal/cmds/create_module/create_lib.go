package create_module

import (
	"log"
	"os/exec"

	"github.com/maxguuse/bruh/internal/types"
)

func createLib(lib *types.Module, project types.ProjectDetails) {
	err := createGoModule(lib, project)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Created Go module: ", lib.Name)

	createMainFileCmd := exec.Command("touch", "main.go")
	createMainFileCmd.Dir = types.LibsDir + "/" + lib.Name
	_, stderr := createMainFileCmd.Output()
	if stderr != nil {
		log.Fatal(stderr)
	}

	log.Println("Created main file: ", "main.go")
}
