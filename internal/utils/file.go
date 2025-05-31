package utils

import (
	"os"
)

// CreateDirectory writes a directory at the specified path
func CreateDirectory(path string) error {
	return os.MkdirAll(path, 0755)
}
