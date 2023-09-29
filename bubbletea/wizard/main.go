package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type wizard struct{}

func (w wizard) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return w, tea.Quit
		}
	}
	return w, nil
}

func (w wizard) View() string {
	return "TODO"
}

func (w wizard) Init() tea.Cmd {
	return tea.EnterAltScreen
}

var _ tea.Model = wizard{}

func main() {
	p := tea.NewProgram(wizard{}, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
