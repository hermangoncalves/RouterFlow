package textinput

import (
	"errors"
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle = lipgloss.NewStyle().Background(lipgloss.Color("#01FAC6")).Foreground(lipgloss.Color("#030303")).Bold(true).Padding(0, 1, 0)
	errorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF8700")).Bold(true).Padding(0, 0, 0)
)

type (
	errMsg error
)

type Output struct {
	Value string
}

func (o *Output) update(val string) {
	o.Value = val
}

type model struct {
	textInput textinput.Model
	err       error
	header    string
	output    *Output
	exit      *bool
}

func InitialTextInputModel(header string, output *Output) model {
	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 30
	exit := false
	return model{
		textInput: ti,
		err:       nil,
		header:    titleStyle.Render(header),
		output:    output,
		exit:      &exit,
	}
}

func CreateErrorInputModel(err error) model {
	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20
	exit := true

	return model{
		textInput: ti,
		err:       errors.New(errorStyle.Render(err.Error())),
		output:    nil,
		header:    "",
		exit:      &exit,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		{
			switch msg.Type {
			case tea.KeyEnter:
				if len(m.textInput.Value()) > 1 {
					m.output.update(m.textInput.Value())
					return m, tea.Quit
				}
			case tea.KeyCtrlC, tea.KeyEsc:
				*m.exit = true
				return m, tea.Quit
			}

		}
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)

	return m, cmd
}

func (m model) View() string {
	return fmt.Sprintf("%s\n\n%s\n\n",
		m.header,
		m.textInput.View(),
	)
}

func (m model) Err() string {
	return m.err.Error()
}

func (m model) GetOutput() string {
	if m.output != nil {
		return m.output.Value
	}
	return ""
}
