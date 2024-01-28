package cmds

import (
	"log"
	"os"
	"os/exec"

	"github.com/maxguuse/bruh/internal/types"
	"gopkg.in/yaml.v2"
)

func InitProject() {
	project, err := runAskForProjectDetailsForm()
	if err != nil || project.Name == "" || project.Owner == "" {
		log.Fatal("Invalid project details")
	}
	log.Println("Project Details: ", project)

	cfg := types.Config{
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
