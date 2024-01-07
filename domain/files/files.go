package files

import (
	"os"
)

func TempDirPath() string {
	return os.TempDir()
}

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
