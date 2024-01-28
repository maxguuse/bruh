package cmds

import (
	"log"
	"os/exec"

	"github.com/maxguuse/bruh/internal/types"
)

func createLib(libName string, project types.ProjectDetails) {
	err := createGoModule(libName, types.Lib, project)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Created Go module: ", libName)

	createMainFileCmd := exec.Command("touch", "main.go")
	createMainFileCmd.Dir = types.LibsDir + "/" + libName
	_, stderr := createMainFileCmd.Output()
	if stderr != nil {
		log.Fatal(stderr)
	}

	log.Println("Created main file: ", "main.go")
}
