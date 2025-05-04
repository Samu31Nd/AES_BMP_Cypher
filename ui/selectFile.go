package ui

import (
	"errors"
	"log"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/filepicker"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	directoryOfFiles = "files"
	filePickerHeight = 9
)

type ModelFP struct {
	Filepicker   filepicker.Model
	SelectedFile string
	quitting     bool
	err          error
}

type clearErrorMsg struct{}

func clearErrorAfter(t time.Duration) tea.Cmd {
	return tea.Tick(t, func(_ time.Time) tea.Msg {
		return clearErrorMsg{}
	})
}

func (m ModelFP) Init() tea.Cmd {
	return m.Filepicker.Init()
}

func (m ModelFP) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		}
	case clearErrorMsg:
		m.err = nil
	}

	var cmd tea.Cmd
	m.Filepicker, cmd = m.Filepicker.Update(msg)

	// Did the user select a file?
	if didSelect, path := m.Filepicker.DidSelectFile(msg); didSelect {
		// Get the path of the selected file.
		m.SelectedFile = path
	}

	// Did the user select a disabled file?
	// This is only necessary to display an error to the user.
	if didSelect, path := m.Filepicker.DidSelectDisabledFile(msg); didSelect {
		// Let's clear the selectedFile and display an error.
		m.err = errors.New(path + " is not valid.")
		m.SelectedFile = ""
		return m, tea.Batch(cmd, clearErrorAfter(2*time.Second))
	}

	return m, cmd
}

func (m ModelFP) View() string {
	if m.quitting {
		return ""
	}
	var s strings.Builder
	s.WriteString("\n")
	s.WriteString(RenderTitle())
	if m.err != nil {
		s.WriteString(m.Filepicker.Styles.DisabledFile.Render(m.err.Error()))
	} else if m.SelectedFile == "" {
		s.WriteString(fourMarginLeft.Render("Selecciona el archivo:"))
	} else {
		s.WriteString(fourMarginLeft.Render("Archivo seleccionado: " + m.Filepicker.Styles.Selected.Render(m.SelectedFile)))
	}
	s.WriteString("\n\n" + fourMarginLeft.Render(m.Filepicker.View()) + "\n" +
		filePickerHelpStyle.Render("↑/k: up • ↓/j: down • esc: Go up directory • enter: select file • q: continue"))
	return s.String()
}

func GetFile() (bool, string) {
	fp := filepicker.New()
	info, errStat := os.Stat(directoryOfFiles)
	if os.IsNotExist(errStat) && !(info.IsDir()) {
		err := os.Mkdir(directoryOfFiles, 777)
		if err != nil {
			log.Fatal("Cant create directory: " + err.Error())
		}
	}
	fp.CurrentDirectory = "./" + directoryOfFiles
	fp.AllowedTypes = []string{".bmp"}
	fp.AutoHeight = false
	fp.ShowPermissions = false
	fp.SetHeight(filePickerHeight)

	m := ModelFP{
		Filepicker: fp,
	}
	tm, _ := tea.NewProgram(&m).Run()
	mm := tm.(ModelFP)
	if mm.SelectedFile == "" {
		return false, ""
	}
	return true, mm.SelectedFile
}
