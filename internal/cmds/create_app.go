package cmds

import (
	"log"
	"os/exec"

	"github.com/maxguuse/bruh/internal/types"
)

func createApp(app *types.Module, project types.ProjectDetails) {
	err := createGoModule(app, project)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Created Go module: ", app.Name)

	createBaseFoldersCmd := exec.Command("mkdir", "cmd", "internal")
	createBaseFoldersCmd.Dir = types.AppsDir + "/" + app.Name
	_, stderr := createBaseFoldersCmd.Output()
	if stderr != nil {
		log.Fatal(stderr)
	}

	log.Println("Created base folders: ", "cmd", "internal")

	createMainFileCmd := exec.Command("touch", "main.go")
	createMainFileCmd.Dir = types.AppsDir + "/" + app.Name + "/cmd"
	_, stderr = createMainFileCmd.Output()
	if stderr != nil {
		log.Fatal(stderr)
	}

	log.Println("Created main file: ", "main.go")
}
