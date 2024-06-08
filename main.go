package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	return m, nil
}

func (m model) View() string {

	return "sup"
}

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatalf("err: %v", err)
	}

	defer f.Close()

	p := tea.NewProgram(model{}, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatalf("err: %v", err)
	}
}
