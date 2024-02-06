package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/charmbracelet/huh"
)

var (
	action      string
	description string
	examples    []string
	id          string
	name        string
	exercises   []Exercise
	exercise    Exercise
	discount    bool
)

type Exercise struct {
	ID          string
	Name        string
	Description string
	Examples    []Example
}
type Example struct {
	input  string
	output string
}

func main() {

	url := "http://localhost:8080/exercises"
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
	AddForm := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("What is the name of the exercise?").
				Value(&name),
			huh.NewInput().
				Title("What is the description of the exercise?").
				Value(&description),
			huh.NewInput().
				Title("What is the example of the exercise?,Inserting input and output with a space between them").
				Value(&name),
		),
	)
	DeleteForm := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("What is the ID of the exercise you want to delete?").
				Value(&id),
		),
	)
	// getByIdForm := huh.NewForm(
	// 	huh.NewGroup(
	// 		huh.NewInput().
	// 			Title("What is the ID of the exercise you want to delete?").
	// 			Value(&id),
	// 	),
	// )
	UpdateForm := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("What is the ID of the exercise you want to update?,Then, fill in only the fields you want to update").
				Value(&id),
		),
		huh.NewGroup(
			huh.NewInput().
				Title("What is the name of the exercise?").
				Value(&name),
			huh.NewInput().
				Title("What is the description of the exercise?").
				Value(&description),
			// huh.NewInput().
			// 	Title("What is the example of the exercise?,Inserting input and output with a space between them").
			// 	Value(&examples[0]),
		),
	)
	err := form.Run()

	if err != nil {
		log.Fatal(err)
	}
	switch action {

	case "add":
		errAdd := AddForm.Run()
		if errAdd != nil {
			log.Fatal(err)
		}

		postBody, _ := json.Marshal(map[string]string{
			"name":        name,
			"description": description,
		})
		responseBody := bytes.NewBuffer(postBody)
		response, err := http.Post(url+"/create", "application/json", responseBody)
		if err != nil {
			fmt.Println("Error:", err)
		}
		defer response.Body.Close()
		// קריאת תוצאת הקריאה
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error:", err)
		}
		// הדפסת תוצאת הקריאה
		fmt.Println("Response:", string(body))

	case "delete":
		errDelete := DeleteForm.Run()
		if errDelete != nil {
			log.Fatal(err)
		}
		request, err := http.NewRequest("delete", url+"/delete/"+id, nil)
		if err != nil {
			fmt.Println("Error creating request:", err)
			return
		}
		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			fmt.Println("Error sending request:", err)
			return
		}
		defer response.Body.Close()
		// הדפסת קוד התגובה
		fmt.Println("Response Status:", response.Status)

	case "update":
		errUpdate := UpdateForm.Run()
		if errUpdate != nil {
			log.Fatal(err)
		}
		data := make(map[string]interface{})
		if name != "" {
			data["name"] = name
		}
		if description != "" {
			data["description"] = description
		}
		if examples == nil {
			data["examples"] = examples
		}
		payload, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Error encoding JSON:", err)
			return
		}
		request, err := http.NewRequest("PUT", url+"/update/"+id, bytes.NewBuffer(payload))
		if err != nil {
			fmt.Println("Error creating request:", err)
			return
		}
		request.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			fmt.Println("Error sending request:", err)
			return
		}
		defer response.Body.Close()
		// הדפסת קוד התגובה
		fmt.Println("Response Status:", response.Status)

	case "getSpecific":
		errDelete := DeleteForm.Run()
		if errDelete != nil {
			log.Fatal(err)
		}
		response, err := http.Get(url + "/getById/" + id)
		if err != nil {
			fmt.Println("Error sending request:", err)
		}
		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			fmt.Errorf("Failed to get item by ID. Status code: %d", response.StatusCode)
		}

		err = json.NewDecoder(response.Body).Decode(&exercise)
		if err != nil {
			fmt.Errorf("Failed to get item by ID.")
		}
		fmt.Printf("Exercise ID: %d, Name: %s, Description: %s\n", exercise.ID, exercise.Name, exercise.Description)

	case "getAll":
		response, err := http.Get(url + "/getAll")
		if err != nil {
			fmt.Println("Failed to get items")
			return
		}
		defer response.Body.Close()
		if response.StatusCode != http.StatusOK {
			fmt.Println("Failed to get items. Status code: %d", response.StatusCode)
			return
		}
		err = json.NewDecoder(response.Body).Decode(&exercises)
		if err != nil {
			fmt.Println("Failed to get items")
		}
		for _, exercise := range exercises {
			fmt.Printf("Exercise ID: %d, Name: %s, Description: %s\n", exercise.ID, exercise.Name, exercise.Description)
		}

	case "check":
		name := `function addone(num){
			return num++;
		 }
		 const myVariable = process.env.MY_VARIABLE;
		 console.log("The signel num is:",addone(myVariable));`
		postBody, _ := json.Marshal(map[string]string{
			"function": name,
			// "description": description,
		})
		responseBody := bytes.NewBuffer(postBody)
		response, err := http.Post(url+"/check/"+"65c0ecc862253d1368d6e48d", "application/json", responseBody)
		if err != nil {
			fmt.Println("Error:", err)
		}
		defer response.Body.Close()
		// קריאת תוצאת הקריאה
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error:", err)
		}
		// הדפסת תוצאת הקריאה
		fmt.Println("Response:", string(body))
	}

	if !discount {
		fmt.Println("What? You didn’t take the discount?!")
	}
}
