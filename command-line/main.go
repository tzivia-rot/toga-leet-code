package main

import (
	httpCall "command-line/http"
	"os"
	"strconv"
	"strings"

	"command-line/modules"
	"fmt"
	"log"

	"github.com/charmbracelet/huh"
)

var (
	action              string
	exercises           []modules.Exercise
	exercise            modules.Exercise
	functionCode        string
	countInput          string
	inputs              []string
	output              string
	input               string
	example             modules.Example
	lenguage            string
	actionExercise      string
	actionExerciseIndex int
)

var uri string

func main() {
	uri = os.Getenv("URI")
	var confirm = true
	for confirm {
		showBegin()
		formFinish := huh.NewForm(
			huh.NewGroup(
				huh.NewConfirm().
					Title("Would you like to continue?").
					Affirmative("Yes!").
					Negative("No.").
					Value(&confirm),
			),
		)
		err := formFinish.Run()

		if err != nil {
			log.Fatal(err)
		}
	}
}

func showBegin() {
	var exercises []modules.Exercise
	exercises, err := httpCall.GetAllExercises()
	if err != nil {
		fmt.Print("ErrorGetAll")
	}
	stringsExercise := exercisesToStrings(exercises)
	formViewExercise := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("This is all exercise").
				Options(
					huh.NewOptions(stringsExercise...)...,
				).
				Value(&actionExercise),
		),
	)
	formActions := huh.NewForm(
		huh.NewGroup(

			huh.NewSelect[string]().
				Title("What you went to do?").
				Options(
					huh.NewOption("Delete Exercise", "delete"),
					huh.NewOption("Update Exercise", "update"),
					huh.NewOption("Check Exercise", "check"),
					huh.NewOption("Get Specific Exercise", "getSpecific"),
				).
				Value(&action),
		))
	err = formViewExercise.Run()
	if err != nil {
		log.Fatal(err)
	}
	index, err := strconv.Atoi(strings.Split(actionExercise, ".")[0])
	if index > len(exercises) {
		otherAction()
	} else {
		exercise.ID = exercises[index-1].ID
		fmt.Print(stringExercise(exercises[index-1]))

		err = formActions.Run()
		if err != nil {
			log.Fatal(err)
		}

		switch action {
		case "delete":
			deleteExercise()
		case "update":
			updateExercise()
		case "getSpecific":
			getExerciseByID()
		case "check":
			checkExercise(exercise)
		}
	}
}

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
			huh.NewText().
				Title("Insert a basus opereation of the exercise for those who want to solve it in go").
				CharLimit(1000000000000000000).
				Value(&exercise.BasisOperationGO),
			huh.NewText().
				CharLimit(1000000000000000000).
				Title("Insert a signature of the exercise for those who want to solve it in node.js").
				Value(&exercise.BasisOperationNodeJS),
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
				Title("Enter output").
				Value(&output),
		),
		huh.NewGroup(
			huh.NewInput().
				Title("How many input the function getting?").
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
		fmt.Printf("enter input %d: ", i+1)
		fmt.Scanln(&input)
		inputs = append(inputs, input)
	}

	example.Inputs = inputs
	inputs = nil
	exercise.Examples = append(exercise.Examples, example)
}
func deleteExercise() {
	response, err := httpCall.DeleteExercise(exercise.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Response Status:", response)
}

func updateExercise() {
	form := huh.NewForm(
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
	exercise, err := httpCall.GetExerciseByID(exercise.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Name: %s, Description: %s\n", exercise.Name, exercise.Description)
}
func exercisesToStrings(exercises []modules.Exercise) []string {
	var result []string
	for i, exercise := range exercises {
		str := fmt.Sprintf("%d. Name: %s, Description: %s", i+1, exercise.Name, exercise.Description)
		result = append(result, str)
	}
	result = append(result, fmt.Sprintf("%d. other action?", len(result)+1))
	return result
}
func otherAction() {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("What you went to do?").
				Options(
					huh.NewOption("Add Exercise", "add"),
				).
				Value(&action),
		),
	)
	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}
	addExercise()

}
func checkExercise(exercise modules.Exercise) {

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("What language do you write in??").
				Options(
					huh.NewOption("node.js", "node.js"),
					huh.NewOption("go", "GO"),
				).
				Value(&lenguage),
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
	response, err := httpCall.CheckExercise(exercise.ID, functionCode, lenguage)
	fmt.Println("Response:", response)
}

func stringExercise(e modules.Exercise) string {
	basisGO := extractFunctionGO(e.BasisOperationGO)
	basisNodeJS := extractFunctionNode(e.BasisOperationNodeJS)
	return fmt.Sprintf("\nName: %s\nDescription: %s\nThe function title in go: %s\nThe function title in node.js: %s\n", e.Name, e.Description, basisGO, basisNodeJS)
}

func extractFunctionNode(code string) string {
	closingBracketIndex1 := strings.Index(code, "function action")
	closingBracketIndex2 := strings.Index(code, "{}")

	if closingBracketIndex1 != -1 && closingBracketIndex2 != -1 {
		return code[closingBracketIndex1 : closingBracketIndex2+1]
	}
	return code
}
func extractFunctionGO(code string) string {
	closingBracketIndex1 := strings.Index(code, "func action")
	closingBracketIndex2 := strings.Index(code, "{}")

	if closingBracketIndex1 != -1 && closingBracketIndex2 != -1 {
		return code[closingBracketIndex1 : closingBracketIndex2+1]
	}
	return code
}
