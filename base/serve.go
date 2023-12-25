package base

import (
	"os/exec"

	"github.com/urfave/cli"
)

var ServeFlags []cli.Flag = []cli.Flag{}

func Serve(c *cli.Context) error {
	// config := c.Args().Get(0)

	return nil
}

func isInstalled(program string) (bool, error) {
	out, err := exec.Command("which", program).Output()
	if err != nil {
		return false, err
	}

	return len(out) != 0, nil
}
