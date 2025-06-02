package utils

import (
	"os"
	"path/filepath"
)

/*
CreateDirectory writes a directory at the specified path
*/
func CreateDirectory(path string) error {
	return os.MkdirAll(path, 0755)
}

/*
WriteFile writes a file with the specified content at the given filename.
*/
func WriteFile(filename, content string) error {
	// Create the folder if it does not exist
	dir := filepath.Dir(filename)
	if err := CreateDirectory(dir); err != nil {
		return err
	}
	// write the file with the content
	return os.WriteFile(filename, []byte(content), 0644)
}
