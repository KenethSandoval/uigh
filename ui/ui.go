package ui

import (
	"context"

	"github.com/KenethSandoval/uigh/ui/common"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/google/go-github/v43/github"

	"github.com/KenethSandoval/uigh/ui/activity"
)

type status int

const (
	statusInit status = iota
	statusLoading
	statusReady
)

type model struct {
	quit     bool
	done     bool
	username string
	gh       *github.Client
	spinner  spinner.Model
	status   status
	activity activity.Model
	errorMsg string
	user     *github.User
}

type (
	userLoadedMsg *github.User
	errorMsg      error
)

func NewProgram(username string, gh *github.Client) *tea.Program {
	return tea.NewProgram(initialModel(username, gh), tea.WithAltScreen())
}

func initialModel(username string, gh *github.Client) model {
	return model{
		username: username,
		status:   statusInit,
		gh:       gh,
		spinner:  common.NewSpinnerModel(),
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
	m, cmd = updateChildren(m, msg)
	cmds = common.AppendIfNotNil(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func updateChildren(m model, msg tea.Msg) (model, tea.Cmd) {
	var cmd tea.Cmd
	switch m.status {
	case statusInit:
		m.spinner, cmd = m.spinner.Update(msg)
		cmd = tea.Batch(cmd, m.loadUserCmd)
		m.status = statusLoading
	case statusLoading:
		m.spinner, cmd = m.spinner.Update(msg)
	}
	return m, cmd
}

func (m model) View() string {
	s := ""
	switch m.status {
	case statusInit:
	case statusLoading:
		s += common.AppStyle().Render(m.spinner.View() + " Loading user...")
	}

	return lipgloss.JoinVertical(lipgloss.Top, s)
}

func (m model) loadUserCmd() tea.Msg {
	if m.status == statusLoading {
		return spinner.Tick()
	}
	user, _, err := m.gh.Users.Get(context.Background(), m.username)
	if err != nil {
		return errorMsg(err)
	}
	return userLoadedMsg(user)
}
