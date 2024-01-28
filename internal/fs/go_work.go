package fs

import (
	"os/exec"

	"github.com/maxguuse/bruh/internal/types"
)

func GoWorkInit() error {
	goWorkInit := exec.Command("go", "work", "init")
	_, err := goWorkInit.Output()
	if err != nil {
		return err
	}

	return nil
}

func GoWorkUse(module *types.Module) error {
	rootDir := module.Type.String()

	addModToWork := exec.Command("go", "work", "use", ".")
	addModToWork.Dir = rootDir + "/" + module.Name
	_, err := addModToWork.Output()
	if err != nil {
		return err
	}

	return nil
}
