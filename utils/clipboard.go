package utils

import (
	"os/exec"
)

func SetClipboard(text string) error {
	cmd := exec.Command("pbcopy")
	in, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}
	if _, err := in.Write([]byte(text)); err != nil {
		return err
	}
	if err := in.Close(); err != nil {
		return err
	}
	return cmd.Wait()
}


