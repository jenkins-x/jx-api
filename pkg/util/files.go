package util

import (
	"fmt"
	"os"
)

const (
	DefaultWritePermissions = 0760
)

// FileExists checks if path exists and is a file
func FileExists(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err == nil {
		return !fileInfo.IsDir(), nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, fmt.Errorf("failed to check if file exists %s: %w", path, err)
}

// DeleteFile deletes a file from the operating system. This should NOT be used to delete any sensitive information
// because it can easily be recovered. Use DestroyFile to delete sensitive information
func DeleteFile(fileName string) (err error) {
	if fileName != "" {
		exists, err := FileExists(fileName)
		if err != nil {
			return fmt.Errorf("could not check if file exists %s due to %s", fileName, err)
		}

		if exists {
			err = os.Remove(fileName)
			if err != nil {
				return fmt.Errorf("Could not remove file due to %s: %w", fileName, err)
			}
		}
	} else {
		return fmt.Errorf("filename is not valid")
	}
	return nil
}
