package exerciseService

import (
	"fmt"
	model "go-lenguage/models"
	"io/ioutil"
	"os"
	"os/exec"
)

func CheckExerciseNode(function string, examples []model.Example) string {

	// Create a folder to contain the function files
	folderName := "function_folder"
	err := os.Mkdir(folderName, 0755)
	if err != nil {
		fmt.Println("Error creating folder:", err)
	}
	// Write function contents to a .js file
	functionFile := folderName + "/function.js"
	err = ioutil.WriteFile(functionFile, []byte(function), 0644)
	if err != nil {
		fmt.Println("Error writing function file:", err)
	}

	// Create package.json file
	packageJSON := `{
		"name": "function_container",
		"version": "1.0.0",
		"scripts": {
			"start": "node function.js"
		}
	}`
	packageFile := folderName + "/package.json"
	err = ioutil.WriteFile(packageFile, []byte(packageJSON), 0644)
	if err != nil {
		fmt.Println("Error writing package.json file:", err)
	}

	// Create Dockerfile
	dockerfile := `FROM node:latest
	WORKDIR /app
	COPY . .
	CMD ["npm", "start"]`
	dockerfileLocation := folderName + "/Dockerfile"
	err = ioutil.WriteFile(dockerfileLocation, []byte(dockerfile), 0644)
	if err != nil {
		fmt.Println("Error writing Dockerfile:", err)
	}
	// Build Docker image
	cmd := exec.Command("docker", "build", "-t", "tziviarot/function_image_node:latest", folderName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error building Docker image:", err)
	}

	//push Docker image
	cmd = exec.Command("docker", "push", "tziviarot/function_image_node:latest")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println("\nError phsh Docker image:", err)
	}

	createAndRunYmlFile, err := createAndRunYmlFile(examples, "tziviarot/function_image_node:latest")

	// fmt.Print("2\n")
	// fmt.Print(st)
	if err != nil {
		fmt.Print(err)
	}
	os.RemoveAll(folderName)
	return createAndRunYmlFile
}
