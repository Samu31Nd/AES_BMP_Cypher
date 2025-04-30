package ui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func GetKey(limit int, placeholder, description string) (quit bool, key string) {
	ti := textinput.New()
	ti.CharLimit = limit
	ti.Placeholder = placeholder
	ti.Width = len(placeholder) | limit
	ti.Focus()

	m := inputModel{
		description: description,
		textInput:   ti,
		err:         nil,
	}
	tm, _ := tea.NewProgram(&m).Run()
	mm := tm.(inputModel)

	return mm.fillValue, mm.textInput.Value()
}

type (
	errMsg error
)

type inputModel struct {
	description string
	textInput   textinput.Model
	err         error
	fillValue   bool
	quit        bool
}

func (m inputModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m inputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyTab:
			if len(m.textInput.Value()) < 16 {
				return m, nil
			}
			m.fillValue = true
			m.quit = true
			return m, tea.Quit
		case tea.KeyCtrlC, tea.KeyEsc:
			m.quit = true
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m inputModel) View() string {
	if m.quit {
		return ""
	}
	return "\n" +
		RenderTitle() +
		fourMarginLeft.Render(m.description+"\n\n"+
			m.textInput.View()+"\n\n"+
			"(esc to quit)\n")
}
