package multiselect

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/hermangoncalves/routerflow/pkg/ui"
)

type model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
	header   string
	footer   string
}

func InitialModelMultiSelect(header string, choices []string, footer string) model {
	return model{
		header:   header,
		choices:  choices,
		selected: make(map[int]struct{}),
		footer:   footer,
	}
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	s := fmt.Sprintf("%s\n\n", m.header)

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ui.FocusedStyle.Render(">")
			choice = ui.FocusedStyle.Render(choice)
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = ui.FocusedStyle.Render("x")
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += m.footer

	s += "\nPress q to quit.\n"
	return s
}
