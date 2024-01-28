package cmds

import (
	"log"
	"os/exec"

	"github.com/maxguuse/bruh/internal/types"
)

func createApp(appName string, project types.ProjectDetails) {
	err := createGoModule(appName, types.App, project)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Created Go module: ", appName)

	createBaseFoldersCmd := exec.Command("mkdir", "cmd", "internal")
	createBaseFoldersCmd.Dir = types.AppsDir + "/" + appName
	_, stderr := createBaseFoldersCmd.Output()
	if stderr != nil {
		log.Fatal(stderr)
	}

	log.Println("Created base folders: ", "cmd", "internal")

	createMainFileCmd := exec.Command("touch", "main.go")
	createMainFileCmd.Dir = types.AppsDir + "/" + appName + "/cmd"
	_, stderr = createMainFileCmd.Output()
	if stderr != nil {
		log.Fatal(stderr)
	}

	log.Println("Created main file: ", "main.go")
}
