package commit

import (
	"errors"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/epicseven-cup/gt/internal/misc"
)

var message bool

func GtCommit(args []string) error {
	// Switch to cobra to make things easier to manage with flags, I want it to work natively with git, and some commands will have bubble tea effects
	return nil

}

// GtMessage corresponding to git commit -m
func GtMessage() error {

	gitRepo, err := misc.DetermineGitRepo()
	if err != nil {
		return err
	}
	if !gitRepo {
		return errors.New("not a git repository")
	}
	pName, err := misc.GetGitName()
	if err != nil {
		return err
	}

	model, err := NewModel(pName)
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
