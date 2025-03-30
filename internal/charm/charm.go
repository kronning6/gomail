package charm

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kronning6/gomail/gmail"
)

type model struct {
	spinner  spinner.Model
	messages []string
	cursor   int
	selected map[int]struct{}
	viewing  int
}

func ProgramModel() model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#f4b8e4"))

	return model{
		spinner:  s,
		messages: []string{},
		selected: make(map[int]struct{}),
	}
}

type emailMsg struct{ messages []string }

func retrieveEmails() tea.Msg {
	messages := gmail.Screener()
	return emailMsg{messages}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(m.spinner.Tick, retrieveEmails)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, CommonKeyMap.Quit):
			return m, tea.Quit

		case key.Matches(msg, CommonKeyMap.Up):
			if m.cursor > 0 {
				m.cursor--
			}

		case key.Matches(msg, CommonKeyMap.Down):
			if m.cursor < len(m.messages)-1 {
				m.cursor++
			}

		case key.Matches(msg, CommonKeyMap.Select):
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
			// case key.Matches(msg, CommonKeyMap.View):
			// 	// m.viewing = m.cursor
			// 	return m, tea.ExitAltScreen
		}
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case emailMsg:
		m.messages = append(m.messages, msg.messages...)
	}

	return m, nil
}

func (m model) View() string {
	var b strings.Builder

	b.WriteString("gmail")
	fmt.Fprintf(&b, " - loading %s", m.spinner.View())

	b.WriteString("\n\n")
	if len(m.messages) > 0 {
		fmt.Fprintf(&b, "%d emails\n\n", len(m.messages))
	}
	for i, choice := range m.messages {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		fmt.Fprintf(&b, "%s [%s] %s\n", cursor, checked, choice)
	}

	b.WriteString("\nPress q to quit.\n")

	return b.String()
}
