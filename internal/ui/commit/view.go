package commit

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ContentDisplay interface {
	View() string
	Reset()
	Update(msg tea.Msg) (textinput.Model, tea.Cmd)
}

type ViewController struct {
	content string
	stages  []ContentDisplay
	current int
}

func NewView() *ViewController {
	header := textinput.New()
	header.Prompt = "Header >>> "
	header.Width = 80
	header.CharLimit = 128
	header.PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("32"))
	header.Cursor.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	header.Focus()

	title := textinput.New()
	title.Prompt = "Title >>> "
	title.Width = 80
	title.CharLimit = 128
	title.PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	title.Cursor.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	title.Focus()

	body := textinput.New()
	body.Prompt = "Body >>> "
	body.CharLimit = 128
	body.ShowSuggestions = true
	body.Focus()

	footer := textinput.New()
	footer.Prompt = "Footer >>> "
	footer.Width = 80
	footer.CharLimit = 128
	footer.PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	footer.Focus()

	return &ViewController{
		content: "PLACEHOLDER",
		stages: []ContentDisplay{
			&header,
			&title,
			&body,
			&footer,
		},
		current: 0,
	}
}

func (v *ViewController) Render() string {
	return v.stages[v.current].View()
}

func (v *ViewController) Update(msg tea.Msg) (textinput.Model, tea.Cmd) {
	return v.stages[v.current].Update(msg)
}

func (v *ViewController) SetCurrentStageContent(model textinput.Model) {
	v.stages[v.current] = &model
}

func (v *ViewController) NextStage() {
	v.current = (v.current + 1) % len(v.stages)
	v.content = v.stages[v.current].View()
}

func (v *ViewController) PreviousStage() {
	v.current = (v.current - 1) % len(v.stages)
	v.content = v.stages[v.current].View()
}

func (v *ViewController) Reset() {
	v.current = 0
	v.content = v.stages[v.current].View()
}

func (v *ViewController) OutputContent() string {
	c := ""
	for _, stage := range v.stages {
		c += stage.View() + "\n"
		stage.Reset()
	}
	return c
}
