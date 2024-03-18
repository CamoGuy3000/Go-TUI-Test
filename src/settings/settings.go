package settings

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func RunSettings() {
	if _, err := tea.NewProgram(settingsmodel{}, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

type settingsmodel struct {
	quitting  bool
}

func (m settingsmodel) Init() tea.Cmd {
	return nil
}

func (m settingsmodel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			m.quitting = true
			return m, tea.Quit
		case " ":
			return m, nil
		}
	}
	return m, nil
}

func (m settingsmodel) View() string {
	if m.quitting {
		_ = tea.ExitAltScreen
		return "Bye!\n"
	}

	return "\n\n Doing the thing :)"
}

