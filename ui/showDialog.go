package ui

import (
	"fmt"
	"log"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func ShowMsgDialog(msg string, error bool) {
	p := tea.NewProgram(msgModel{
		timeCounter: 5,
		message:     msg,
		error:       error,
	})
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

// A model can be more or less any type of data. It holds all the data for a
// program, so often it's a struct. For this simple example, however, all
// we'll need is a simple integer.
type msgModel struct {
	timeCounter int
	message     string
	error       bool
	quit        bool
}

// Init optionally returns an initial command we should run. In this case we
// want to start the timer.
func (m msgModel) Init() tea.Cmd {
	return tick
}

// Update is called when messages are received. The idea is that you inspect the
// message and send back an updated model accordingly. You can also return
// a command, which is a function that performs I/O and returns a message.
func (m msgModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quit = true
			return m, tea.Quit
		case "ctrl+z":
			return m, tea.Suspend
		}

	case tickMsg:
		m.timeCounter--
		if m.timeCounter <= 0 {
			m.quit = true
			return m, tea.Quit
		}
		return m, tick
	}
	return m, nil
}

// View returns a string based on data in the model. That string which will be
// rendered to the terminal.
func (m msgModel) View() string {
	if m.quit {
		return "\n" + m.message + "\n\n"
	}
	msg := m.message
	if m.error {
		msg = errorStyle.Render(msg)
	}
	return fmt.Sprintf("\n%s\n%s\n\n    This message will quit in %d seconds\n", RenderTitle(), fourMarginLeft.Render(msg), m.timeCounter)
}

// Messages are events that we respond to in our Update function. This
// particular one indicates that the timer has ticked.
type tickMsg time.Time

func tick() tea.Msg {
	time.Sleep(time.Second)
	return tickMsg{}
}
