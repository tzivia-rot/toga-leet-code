package main

import (
	httpCall "command-line/http"
	"os"
	"strings"

	"strconv"

	"command-line/modules"
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

var (
	action             string
	description        string
	examples           []string
	id                 string
	name               string
	exercises          []modules.Exercise
	exercise           modules.Exercise
	discount           bool
	functionCode       string
	checkForPlagiarism bool
	countInput         string
	inputs             []string
	output             string
	input              string
	example            modules.Example
	lenguage           string
)

var url = "http://localhost:8080/exercises"

func addExercise() {
	addExerciseBasic()
	addExerciseEample()
	addExerciseEample()
	addExerciseEample()
	response, err := httpCall.AddExercise(exercise)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("there is! The exercise was successfully added! This is his code: ", response)

}
func addExerciseBasic() {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("What is the name of the exercise?").
				Value(&exercise.Name),
			huh.NewInput().
				Title("What is the description of the exercise?").
				Value(&exercise.Description),
			huh.NewInput().
				Title("Insert a signature of the exercise for those who want to solve it?").
				Value(&exercise.BasisOperation),
		),
	)
	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

}

func addExerciseEample() {
	formOutput := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("הכנס פלט").
				Value(&output),
		),
		huh.NewGroup(
			huh.NewInput().
				Title("כמה קלטים?").
				Value(&countInput),
		),
	)
	err := formOutput.Run()
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}
	example.Output = output
	countInputInt, err := strconv.Atoi(countInput)
	for i := 0; i < countInputInt; i++ {
		var input string
		fmt.Printf("הכנס קלט %d: ", i+1)
		fmt.Scanln(&input)
		inputs = append(inputs, input)
	}

	example.Inputs = inputs
	inputs = nil
	exercise.Examples = append(exercise.Examples, example)
}
func deleteExercise() {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("What is the ID of the exercise you want to delete?").
				Value(&exercise.ID),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}
	response, err := httpCall.DeleteExercise(exercise.ID)
	fmt.Println("Response Status:", response)
}

func updateExercise() {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("What is the ID of the exercise you want to update?").
				Value(&exercise.ID),
		),
		huh.NewGroup(
			huh.NewInput().
				Title("What is the new name of the exercise?").
				Value(&exercise.Name),
			huh.NewInput().
				Title("What is the new description of the exercise?").
				Value(&exercise.Description),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	data := map[string]string{
		"name":        exercise.Name,
		"description": exercise.Description,
	}
	response, err := httpCall.UpdateExercise(exercise.ID, data)

	fmt.Println("Response Status:", response)
}

func getExerciseByID() {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("What is the ID of the exercise you want to view?").
				Value(&exercise.ID),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}
	exercise, err = httpCall.GetExerciseByID(exercise.ID)
	// הדפסת פרטי התרגיל
	fmt.Printf("Exercise ID: %s, Name: %s, Description: %s\n", exercise.ID, exercise.Name, exercise.Description)
}

func getAllExercises() {

	var exercises []modules.Exercise
	exercises, err := httpCall.GetAllExercises()
	if err != nil {
		fmt.Print("erroeGetAll")
	}
	for _, exercise := range exercises {
		fmt.Printf("Name: %s, Description: %s\n", exercise.Name, exercise.Description)
	}

	result := tableExercise(exercises)
	fmt.Print(result)
}

func checkExercise() {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("What language do you write in??").
				Options(
					huh.NewOption("node.js", "node.js"),
					huh.NewOption("go", "GO"),
				).
				Value(&lenguage),
			huh.NewText().
				Title("what is the id?").
				Value(&id),
		),
	)
	formFunctionCode := huh.NewForm(
		huh.NewGroup(
			huh.NewText().
				Title("Enter the code.").
				Value(&functionCode).CharLimit(1000000000000000000),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}
	err = formFunctionCode.Run()
	if err != nil {
		log.Fatal(err)
	}
	response, err := httpCall.CheckExercise(id, functionCode, lenguage)
	fmt.Println("Response:", response)
}

func main() {

	var confirm = true
	for confirm {
		form := huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[string]().
					Title("What you went to do?").
					Options(
						huh.NewOption("Add Exercise", "add"),
						huh.NewOption("Delete Exercise", "delete"),
						huh.NewOption("Update Exercise", "update"),
						huh.NewOption("Check Exercise", "check"),
						huh.NewOption("Get All Exercises", "getAll"),
						huh.NewOption("Get Specific Exercise", "getSpecific"),
					).
					Value(&action),
			),
		)
		err := form.Run()

		if err != nil {
			log.Fatal(err)
		}
		switch action {
		case "add":
			addExercise()
		case "delete":
			deleteExercise()
		case "update":
			updateExercise()
		case "getSpecific":
			getExerciseByID()
		case "getAll":
			getAllExercises()
		case "check":
			checkExercise()
		}
		formFinish := huh.NewForm(
			huh.NewGroup(
				huh.NewConfirm().
					Title("Are you sure?").
					Affirmative("Yes!").
					Negative("No.").
					Value(&confirm),
			),
		)
		err = formFinish.Run()

		if err != nil {
			log.Fatal(err)
		}
	}
}

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
