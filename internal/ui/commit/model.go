package commit

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/epicseven-cup/gt/internal/cache"
	"log"
	"os/exec"
)

type Model struct {
	ViewController *ViewController
	Cache          *cache.Cache
	Commit         string
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			m.ViewController.NextStage()
			return m, nil
		case "ctrl+y":
			go func() {
				err := m.commit()
				if err != nil {
					log.Fatal(err)
				}
			}()
			return m, tea.Quit

		}

	default:
		return m, nil
	}

	tm, cmd := m.ViewController.Update(msg)
	m.ViewController.SetCurrentStageContent(tm)
	return m, cmd
}
func (m Model) View() string {
	v := m.ViewController.Render()
	h := "[ctrl+q] for quit, [ctrl+y] for accept / finish"
	return fmt.Sprintf("%s\n\n%s", v, h)
}

func NewModel(projectName string) (Model, error) {

	c, err := cache.GetCache(projectName)
	if err != nil {
		return Model{}, err
	}
	v := NewView()
	return Model{
		ViewController: v,
		Cache:          c,
		Commit:         "",
	}, nil
}

func (m Model) commit() error {
	m.Commit = m.ViewController.OutputContent()

	out, err := exec.Command("git", "commit", "-m", m.Commit).Output()
	if err != nil {
		return err
	}
	m.ViewController.content = string(out)
	return nil
}
