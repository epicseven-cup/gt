package commit

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/epicseven-cup/gt/internal/cache"
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
		case "ctrl+c", "q":
			return m, tea.Quit

		case "enter":
			m.ViewController.NextStage()
			return m, nil
		}

	default:
		return m, nil
	}
	return m, nil
}
func (m Model) View() string {
	return m.ViewController.Render()
}

func NewModel(projectName string) (Model, error) {

	c, err := cache.NewCache(projectName)
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
