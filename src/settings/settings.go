package settings

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var rowsToggle = []Toggle{
	Toggle{"Explorability", "List", []string{"Familiar", "Mixed", "Discovery"}, 0}, // Choose between Familiar, Mixed, Discovery
	Toggle{"Genre Preferences", "List", []string{"Rock", "Country", "Indie"}, 0}, // Select preferred music genres or exclude genres
	Toggle{"Artist Preferences", "List", []string{"New", "Same", "Mixed"}, 0}, // Follow specific artists or bands
	Toggle{"Language Settings", "List", []string{"English", "Nope"}, 0}, // Filter songs by language
	Toggle{"Content Explicitness", "Toggle", []string{"On", "Off"}, 0}, // Allow or filter out explicit content
	Toggle{"Notifications", "Toggle", []string{"On", "Off"}, 0}, // Receive notifications for recommendations, updates, offers
	Toggle{"Discovery Frequency", "Frequency", []string{"1", "2", "3"}, 0}, // Set how often new songs are introduced
	Toggle{"Feedback Options", "Toggle", []string{"On", "Off"}, 0}, // Provide feedback on suggestions
	Toggle{"Social Integration", "Toggle", []string{"On", "Off"}, 0}, // Connect with social media for sharing and seeing what friends listen to
	Toggle{"Dark Mode/Light Mode", "Mode", []string{"Light", "Dark"}, 0}, // Choose UI theme for comfort
	Toggle{"Sleep Timer", "Timer", []string{"Off", "10 Min", "20 Min", "30 Min"}, 0}, // Set a timer to stop playback
}

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type Toggle struct {
	name    string
	mode    string
	options []string
	value   int
}

func toggle(tog Toggle) string {
	// switch rows[index][1]{

	// }
	if(tog.value == len(tog.options)){
		tog.value = 0
	} else {
		tog.value++
	}
	tog.mode = tog.options[tog.value]
	return tog.mode
}

func RunSettings() {
	// explore := []string{"Familiar", "Mixed", "Discovery"}

	cols := []table.Column{
		{Title: "Setting", Width: 50},
		{Title: "Value", Width: 30},
	}
	

	rows := make([]table.Row, len(rowsToggle)+1)
	for i, toggle := range rowsToggle{
		// fmt.Println(rowsToggle)
		fmt.Println(rows)
		rows[i] = table.Row([]string{toggle.name, toggle.options[toggle.value]})
	}

	t := table.New(
		table.WithColumns(cols),
		table.WithRows(rows),
		table.WithFocused(true),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("240")).
			BorderBottom(true).
			Bold(false)
	s.Selected = s.Selected.
			Foreground(lipgloss.Color("229")).
			Background(lipgloss.Color("57")).
			Bold(false)
	t.SetStyles(s)


	if _, err := tea.NewProgram(settingsmodel{t, false}, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

type settingsmodel struct {
	table     table.Model
	quitting  bool
}

func (m settingsmodel) Init() tea.Cmd {
	return nil
}

func (m settingsmodel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) { // What type of message do we have
	case tea.KeyMsg: // If it is a keystroke
		switch msg.String() {
		case "q", "ctrl+c", "esc": // and you press quit, quit the program
			m.quitting = true
			return m, tea.Quit
		
		case " ", "enter": // select what you are currently highlighting
			m.table.Rows()[m.table.Cursor()][1] = toggle(rowsToggle[m.table.Cursor()])
			//TODO Fix the delay in change
		}
	}

	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m settingsmodel) View() string {
	if m.quitting {
		_ = tea.ExitAltScreen
		return "Bye!\n"
	}

	return baseStyle.Render(m.table.View()) + "\n"
}

