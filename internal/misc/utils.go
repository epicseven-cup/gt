package misc

import (
	"os"
	"os/exec"
	"strings"
)

func DetermineGitRepo() (bool, error) {
	_, err := os.Stat("./.git")

	if err != nil {
		return false, err
	}

	return true, nil

}

func GetGitName() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	path := strings.TrimSpace(string(out))
	// strings.LastIndex return -1 if it does not exist
	pName := path[strings.LastIndex(path, "/"):]
	return pName, nil
}
