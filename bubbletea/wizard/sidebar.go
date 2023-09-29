package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type sidebar struct {
	width, height int

	steps      []string
	activeStep int
}

func (s sidebar) Init() tea.Cmd {
	//TODO implement me
	panic("implement me")
}

func (s sidebar) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	//TODO implement me
	panic("implement me")
}

func (s sidebar) View() string {
	var content string

	for i, step := range s.steps {
		step = fmt.Sprintf("%d. %s", i+1, step)
		if i == s.activeStep {
			content += lipgloss.NewStyle().
				Foreground(lipgloss.Color("#E95678")).
				Render(step)
		} else {
			content += step
		}
		content += "\n"
	}

	return lipgloss.NewStyle().
		Width(s.width - 2).
		Height(s.height - 2).
		Border(lipgloss.NormalBorder()).
		Render(content)
}

var _ tea.Model = sidebar{}
