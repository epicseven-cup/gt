package commit

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
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
		}
	default:
		return m, nil
	}
	return m, nil
}
func (m Model) View() string {
	return fmt.Sprintf("")
}

func NewModel() Model {
	return Model{}
}
