package ui

import (
	"fmt"
	"io"
	"labAES28_04/utils"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	listHeight   = 15
	programTitle = "🔐 Practica de cifrado AES con modos de operación"
)

type item string

func (i item) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%s", i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	_, _ = fmt.Fprint(w, fn(str))
}

type model struct {
	list     list.Model
	choice   string
	quitting bool
	selected bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c", "esc":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.selected = true
				m.choice = string(i)
			}
			m.quitting = true
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func RenderTitle() string {
	return borderUnder.Render(prevTitle.Render("")+titleStyle.Render(programTitle)) + "\n\n"
}

func (m model) View() string {
	if m.quitting {
		return ""
	}
	return "\n" +
		RenderTitle() +
		m.list.View()
}

func GetOption(title string, options []string) (quit bool, option int) {
	items := make([]list.Item, 0, len(options))
	for i := 0; i < len(options); i++ {
		items = append(items, item(options[i]))
	}
	const defaultWidth = 20

	l := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	l.Title = title
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleListStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	m := model{list: l}

	p := tea.NewProgram(m)
	finalModel, err := p.Run()
	if err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	// Convertir finalModel a tu tipo concreto para acceder a su lista
	if m, ok := finalModel.(model); ok {
		selectedItem := m.list.SelectedItem()
		if selectedItem != nil {
			i := utils.IndexOf(m.choice, options)
			return m.selected, i
		}
	}
	return m.selected, 0
}
