package utils

import (
	"os/exec"
	"strings"
)

func GetDefaultBrowser() (string, error) {
	cmd := exec.Command("xdg-settings", "get", "default-web-browser")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	browser := strings.TrimSpace(string(output))

	if strings.HasSuffix(browser, ".desktop") {
		browser = strings.Split(browser, ".desktop")[0]
	}

	return browser, nil
}
