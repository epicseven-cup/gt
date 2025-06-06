package main

import (
	"errors"
	"flag"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/epicseven-cup/gt/internal/ui/commit"
	"os"
	"os/exec"
	"strings"
)

const VERSION = "0.0.1"

func determineGitRepo() (bool, error) {
	_, err := os.Stat("./.git")

	if err != nil {
		return false, err
	}

	return true, nil

}

func getGitName() (string, error) {
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

func gtCommit(mode string, args []string) error {
	gitRepo, err := determineGitRepo()
	if err != nil {
		return err
	}
	if !gitRepo {
		return errors.New("not a git repository")
	}
	pName, err := getGitName()
	if err != nil {
		return err
	}

	model, err := commit.NewModel(pName)
	if err != nil {
		return err
	}

	p := tea.NewProgram(model)
	_, err = p.Run()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("usage: gt <command>, gt -h")
		os.Exit(1)
	}
	fmt.Println(args)
	mode := args[0]
	switch mode {
	case "help":
		flag.Usage()
	case "version":
		fmt.Printf("gt version: %s\n", VERSION)
	case "commit":
		err := gtCommit(mode, args[1:])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	default:
		out, err := exec.Command("git", args...).Output()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(out))
	}
}
