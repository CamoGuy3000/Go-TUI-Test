package screen

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func RunButton() {
	p := tea.NewProgram(Button(50, 20, 20, 20, Click), tea.WithMouseAllMotion())
	if _, err := p.Run(); err != nil {
		fmt.Print("ERROR")
		os.Exit(1)
	}
}

type buttonType int

const (
	Toggle buttonType = iota
	Click
)

type button struct {
	// where the center is
	xCord  int
	yCord  int

	height int // defines the distance from the center to the top and bottom
	width  int // defines the distance from the center to the right and left

	left   int
	right  int
	top    int
	bottom int

	bType  buttonType
}

func Button(x, y, h, w int, ty buttonType) button {
	return button{
		xCord: x,
		yCord: y,
		height: h,
		width: w,

		left: x-w,
		right: x+w,
		top: y+h,
		bottom: y-h,

		bType: ty,
	}
}

func (b button) Init() tea.Cmd {
	return nil
}

func (b button) Update(msg tea.Msg) (tea.Model, tea.Cmd){
	// s := ""

	switch msg := msg.(type) {
	
	case tea.KeyMsg:
		if s := msg.String(); s == "ctrl+c" || s == "q" || s == "esc" {
			return b, tea.Quit
		}

	case tea.MouseMsg:
		switch tea.MouseEvent(msg).String(){
		case "left press":
			if msg.X <= b.left || msg.X >= b.right || msg.Y <= b.bottom || msg.Y >= b.height {
				return b, tea.Printf("(%d, %d) not in (left=%d, right=%d . bottom=%d, top=%d)", msg.X, msg.Y, b.left, b.right, b.bottom, b.top)
			}
			return b, tea.Printf("Button Pressed at (%d, %d) with (x=%d, y=%d . left=%d, right=%d . bottom=%d, top=%d)", msg.X, msg.Y, b.xCord, b.yCord, b.left, b.right, b.bottom, b.top)

		case "left release":
			return b, nil
		default:
			return b, nil
		}

		

		// if tea.MouseEvent(msg).String() == "left press"{
		// 	return b, tea.Printf("left press")
		// }
		// return b, tea.Printf("(X: %d, Y: %d) %s", msg.X, msg.Y, tea.MouseEvent(msg))
	

	}

	return b, nil
}

func (b button) View() string {
	return ""

}
