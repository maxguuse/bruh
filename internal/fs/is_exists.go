package fs

import (
	"os"
)

func IsExists(file string) bool {
	if _, err := os.Stat(file); !os.IsNotExist(err) {
		return true
	}

	return false
}
