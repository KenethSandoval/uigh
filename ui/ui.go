package ui

import (
	"github.com/KenethSandoval/uigh/ui/common"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type status int

const (
	statusInit status = iota
	statusLoading
)

type model struct {
	quit     bool
	spinner  spinner.Model
	status   status
	username string
}

func NewProgram(username string) *tea.Program {
	return tea.NewProgram(initialModel(username), tea.WithAltScreen())
}

func initialModel(username string) model {
	return model{
		username: username,
		status:   statusInit,
	}
}

func (m model) Init() tea.Cmd {
	return spinner.Tick
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			m.quit = true
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		// TODO: windows resizing
	case spinner.TickMsg:
		m.spinner, cmd = m.spinner.Update(msg)
	}
	var cmds []tea.Cmd
	cmds = common.AppendIfNotNil(cmds, cmd)
	m, cmd = updateChildre(m, msg)
	cmds = common.AppendIfNotNil(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func updateChildre(m model, msg tea.Msg) (model, tea.Cmd) {
	var cmd tea.Cmd
	switch m.status {
	case statusInit:
	case statusLoading:
		m.spinner, cmd = m.spinner.Update(msg)
		cmd = tea.Batch(cmd, m.loadUserCmd)
		m.status = statusLoading
	}
	return m, cmd
}

func (m model) View() string {
	s := "KENETH"
	switch m.status {
	case statusInit:
		s += " Loading user..."
	}

	return lipgloss.JoinVertical(lipgloss.Top, s)
}

func (m model) loadUserCmd() tea.Msg {
	if m.status == statusLoading {
		return spinner.Tick()
	}
	return tea.ExitAltScreen()
}
