package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type wizard struct {
	sidebar sidebar

	width  int
	height int
}

func (w wizard) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return w, tea.Quit
		}

	case tea.WindowSizeMsg:
		w.width = msg.Width
		w.height = msg.Height
		w.sidebar.width = msg.Width / 4
		w.sidebar.height = msg.Height
	}
	return w, nil
}

func (w wizard) View() string {
	return lipgloss.JoinHorizontal(
		lipgloss.Center,
		w.sidebar.View(),
		"World",
	)
}

func (w wizard) Init() tea.Cmd {
	return tea.EnterAltScreen
}

var _ tea.Model = wizard{}

func main() {
	p := tea.NewProgram(wizard{sidebar: sidebar{
		steps: []string{
			"Intro",
			"Step 1",
			"Step 2",
			"Step 3",
			"Complete",
		},
	}}, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
