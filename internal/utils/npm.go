package utils

import (
	"os/exec"
)

// RunNpmInit init npm project
func RunNpmInit() error {
	cmd := exec.Command("npm", "init", "-y")
	cmd.Stdout = nil
	cmd.Stderr = nil
	return cmd.Run()
}
