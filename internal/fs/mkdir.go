package fs

import (
	"os/exec"
)

func Mkdir(root string, dirs ...string) error {
	createBaseFoldersCmd := exec.Command("mkdir", dirs...)
	createBaseFoldersCmd.Dir = root

	_, err := createBaseFoldersCmd.Output()
	if err != nil {
		return err
	}

	return nil
}
