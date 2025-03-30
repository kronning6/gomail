package charm

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type KeyMap struct {
	Quit   key.Binding
	Up     key.Binding
	Down   key.Binding
	Select key.Binding
	View   key.Binding
	Back   key.Binding
	Delete key.Binding
}

var CommonKeyMap = KeyMap{
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q/ctrl+c", "quit"),
	),
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "down"),
	),
	Select: key.NewBinding(
		key.WithKeys("x"),
		key.WithHelp("x", "select"),
	),
	View: key.NewBinding(
		key.WithKeys("enter", "o"),
		key.WithHelp("enter", "view"),
	),
	Back: key.NewBinding(
		key.WithKeys("u"),
		key.WithHelp("u", "back to threadlist"),
	),
	Delete: key.NewBinding(
		key.WithKeys("#"),
		key.WithHelp("#", "delete"),
	),
}

func StartTeaProgram(model tea.Model) {
	p := tea.NewProgram(model, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
