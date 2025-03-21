package appFs

import (
	"os"
)

func EnsureDirectoryExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0777)

		return err
	}

	return nil
}
