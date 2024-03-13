package main

import (
	"fmt"
	"os"

	"github.com/CamoGuy3000/mouse"
	"github.com/CamoGuy3000/screen"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choices []string
	programs []func()
	selector int
}

var exit = false

func initModel() model {
	return model{
		choices: []string{
			"Mouse coords",
		  "Toggle fullscreen",
			"Button"},
		programs: []func(){
			mouse.RunCords,
			screen.RunAlt,
			screen.RunButton,
		},
		selector: 0,
	}
}

func main() {

	p := tea.NewProgram(initModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("ERRORRRRR")
		os.Exit(1)
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if exit {
		os.Exit(1)
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		
		case "up":
			if m.selector > 0 {
				m.selector--
			}
		
		case "down":
			if m.selector < len(m.programs)-1 {
				m.selector++
			}
		case "enter":
			run := m.programs[m.selector]
			exit = true
			run()
		}
	}

	return m, nil
}

func (m model) View() string {
	s := ""
	if exit {
		s += "\n\n\n"
	}
	s += "Choose what test program you want to run\n\n"

	// fmt.Printf("\n\n Toggle: %d, Click: %d \n\n", screen.Toggle, screen.Click)
	

	for i, choice := range m.choices {
		cursor := " "
		if m.selector == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += "q to quit.\n"

	return s
}

