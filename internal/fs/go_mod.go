package fs

import (
	"os/exec"
)

func GoModInit(root, importPath string) error {
	goModInitCmd := exec.Command("go", "mod", "init", importPath)
	goModInitCmd.Dir = root
	_, err := goModInitCmd.Output()
	if err != nil {
		return err
	}

	return nil
}
