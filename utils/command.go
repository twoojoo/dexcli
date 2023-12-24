package utils

import (
	"os/exec"
)

func RunCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	_, err := cmd.Output()
	return err
}
