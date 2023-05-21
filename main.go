package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type status int

const (
	todo status = iota
	inProgress
	done
)

/* CUSTOM ITEM */

type Task struct {
	status      status
	title       string
	description string
}

// implement the list.Item interface
func (t Task) FilterValue() string {
	return t.title
}

func (t Task) Title() string {
	return t.title
}

func (t Task) Description() string {
	return t.description
}

/* MAIN MODEL */

type Model struct {
	focused status
	lists   []list.Model
	err     error
	loaded  bool
}

func New() *Model {
	return &Model{}
}

// TODO: call thi on tea.WindowSizeMsg
func (m *Model) initLists(width, height int) {
	defaultList := list.New([]list.Item{}, list.NewDefaultDelegate(), width, height)
	defaultList.SetShowHelp(false)

	m.lists = []list.Model{defaultList, defaultList, defaultList}
	// init To Do
	m.lists[todo].Title = "To Do"
	m.lists[todo].SetItems([]list.Item{
		Task{status: todo, title: "Write blog post", description: "Write a blog post about bubbles"},
		Task{status: todo, title: "Walk the dog", description: "Take the dog for a walk"},
		Task{status: todo, title: "Do the dishes", description: "Do the dishes and clean the kitchen"},
	})

	// init In Progress
	m.lists[inProgress].Title = "In Progress"
	m.lists[inProgress].SetItems([]list.Item{
		Task{status: inProgress, title: "Write more code", description: "Write more code for the blog post"},
	})

	// init Done
	m.lists[done].Title = "Done"
	m.lists[done].SetItems([]list.Item{
		Task{status: done, title: "Miss walk in the park", description: "Miss walk in the part because I was writing code"},
	})
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if !m.loaded {
			m.initLists(msg.Width, msg.Height)
			m.loaded = true
		}
	}
	var cmd tea.Cmd
	m.lists[m.focused], cmd = m.lists[m.focused].Update(msg)
	return m, cmd
}

func (m Model) View() string {
	if m.loaded {
		return lipgloss.JoinHorizontal(
			lipgloss.Left,
			m.lists[todo].View(),
			m.lists[inProgress].View(),
			m.lists[done].View())

	} else {
		return "Loading..."
	}

}

func main() {
	m := New()
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
