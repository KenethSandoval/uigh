package ui

import (
	"github.com/KenethSandoval/uigh/ui/common"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type status int

type model struct {
	quit     bool
	done     bool
	username string
	spinner  spinner.Model
	errorMsg string
	status   status
}

func NewProgram(username string) *tea.Program {
	return tea.NewProgram(initialModel(username), tea.WithAltScreen())
}

func initialModel(username string) model {
	return model{
		username: username,
		spinner:  common.NewSpinnerModel(),
	}
}

func (m model) Init() tea.Cmd {
	return spinner.Tick
}

func (m model) Update() {}

func (m model) View() {}
