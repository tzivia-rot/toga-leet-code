package main

import (
	"command-line/modules"
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type model struct {
	table table.Model
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			return m, tea.Batch(
				tea.Printf("Let's go to %s!", m.table.SelectedRow()[1]),
			)
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}

func tableExercise(exercises []modules.Exercise) string {
	columns := []table.Column{
		{Title: "ID", Width: 4},
		{Title: "Name", Width: 15},
		{Title: "Discription", Width: 20},
		{Title: "BasisOperation", Width: 20},
		{Title: "Examples", Width: 25},
	}

	var exerciseRows []table.Row
	for _, exercise := range exercises {
		exampleString := ""
		for _, example := range exercise.Examples {
			exampleString += fmt.Sprintf("Input: %s, Output: %s\n", strings.Join(example.Inputs, ", "), example.Output)
		}
		exerciseRow := table.Row{
			exercise.ID,
			exercise.Name,
			exercise.Description,
			exercise.BasisOperation,
			exampleString,
		}
		exerciseRows = append(exerciseRows, exerciseRow)
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(exerciseRows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	m := model{t}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
	return "sucsses"
}
