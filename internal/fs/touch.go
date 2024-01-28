package fs

import (
	"os/exec"
)

func Touch(root string, files ...string) error {
	createMainFileCmd := exec.Command("touch", files...)
	createMainFileCmd.Dir = root

	_, err := createMainFileCmd.Output()
	if err != nil {
		return err
	}

	return nil
}
