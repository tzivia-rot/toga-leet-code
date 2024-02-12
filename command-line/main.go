package main

import (
	httpCall "command-line/http"
	"command-line/modules"
	"fmt"
	"log"

	"github.com/charmbracelet/huh"
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
	// קבלת פרטי התרגיל מהמשתמש
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
	formInput := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("הכנס קלט").
				Value(&input),
		),
	)
	// for
	// ex := Example{}

	// ex.Input.append(input1)

	// ex.Output.append(output)

	// var numberExample int
	// fmt.Print("How many examples do you have to enter?")

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	// err := formOutput.Run()
	// if err != nil {

	// 	log.Fatal(err)
	// }
	err = formOutput.Run()
	if err != nil {
		log.Fatal(err)
	}
	example.Output = output
	// countInput, err := strconv.Atoi(countInput)
	if err != nil {
		log.Fatal(err)
	}
	// j := 0

	// for {
	err = formInput.Run()
	if err != nil {
		log.Fatal(err)
	}
	inputs = append(inputs, input)
	// 	j++
	// 	if j > countInput {
	// 		break
	// 	}
	// }
	example.Inputs = inputs
	exercise.Examples = append(exercise.Examples, example)

	response, err := httpCall.AddExercise(exercise)
	fmt.Println("there is! The exercise was successfully added! This is his code: ", response)

}

func deleteExercise() {
	// קבלת מזהה התרגיל למחיקה מהמשתמש
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
	// קבלת מזהה התרגיל לעדכון מהמשתמש
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

	// בניית הבקשה לשרת
	data := map[string]string{
		"name":        exercise.Name,
		"description": exercise.Description,
	}
	response, err := httpCall.UpdateExercise(exercise.ID, data)

	// הדפסת קוד התגובה
	fmt.Println("Response Status:", response)
}

func getExerciseByID() {
	// קבלת מזהה התרגיל להצגה מהמשתמש
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

	// קריאת רשימת התרגילים מהתגובה
	var exercises []modules.Exercise
	exercises, err := httpCall.GetAllExercises()
	if err != nil {
		fmt.Print("erroeGetAll")
	}
	// הדפסת רשימת התרגילים
	for _, exercise := range exercises {
		fmt.Printf("Exercise ID: %s, Name: %s, Description: %s\n", exercise.ID, exercise.Name, exercise.Description)
	}
}

func checkExercise() {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("What you went to do?").
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
	formNodeJS := huh.NewForm(
		huh.NewGroup(
			huh.NewText().
				Title("Tell me a story.").
				Value(&functionCode).CharLimit(100000000000000000),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}
	// err = formNodeJS.Run()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// switch lenguage {
	// case "node.js":

	// case "go":

	// }
	err = formNodeJS.Run()
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
