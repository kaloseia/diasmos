package files

import (
	"os"
)

func TempDirPath() string {
	return os.TempDir()
}

// BNUDO: This, and then file_system_registry_test / file_system_registry
func DirAccessible(dirPath string) bool {
	tmpHandle, readErr := os.Open(dirPath)
	tmpHandle.Close()
	if readErr != nil {
		return false
	}
	return true
}

func NewTempDir() (string, error) {
	return "", nil
}
