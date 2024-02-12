package formtable

import (
	"command-line/modules"
	"fmt"
	"log"
	"strconv"

	"github.com/charmbracelet/huh"
)

func addExerciseBasic(exercise *modules.Exercise, example *modules.Example, output *string, countInput *string, inputs *[]string) {
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
	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

}

func addExerciseEample(exercise *modules.Exercise, example *modules.Example, output *string, countInput *string, inputs *[]string) {
	formOutput := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("הכנס פלט").
				Value(output),
		),
		huh.NewGroup(
			huh.NewInput().
				Title("כמה קלטים?").
				Value(countInput),
		),
	)

	example.Output = *output
	countInputInt, err := strconv.Atoi(*countInput)
	for i := 0; i < countInputInt; i++ {
		var input string
		fmt.Printf("הכנס קלט %d: ", i+1)
		fmt.Scanln(&input)
		*inputs = append(*inputs, input)
	}
	err = formOutput.Run()
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	example.Inputs = *inputs
	exercise.Examples = append(exercise.Examples, *example)
}

func deleteExercise(exercise *modules.Exercise) {
	// קוד מחיקת התרגיל
}

func updateExercise(exercise *modules.Exercise) {
	// קוד עדכון התרגיל
}

func getExerciseByID(exercise *modules.Exercise) {
	// קוד לקבלת התרגיל לפי המזהה
}

func getAllExercises() {
	// קוד לקבלת כל התרגילים
}

func checkExercise(exercise *modules.Exercise, functionCode *string, language *string) {
	// קוד לבדיקת התרגיל
}
