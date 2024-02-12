package exerciseService

import (
	"fmt"
	model "go-lenguage/models"
	"io/ioutil"
	"os"
	"os/exec"
)

func CheckExerciseGO(function string, examples []model.Example) string,err{

	folderName := "function_folder"
	err := os.Mkdir(folderName, 0755)
	if err != nil {
		return "Error creating folder:",err
	}

	// Write function contents to a go file
	functionFile := folderName + "/function.go"
	err = ioutil.WriteFile(functionFile, []byte(function), 0644)
	if err != nil {
		return "Error writing function file:", err
	}
	// Create go.mod file
	packageMOD := `module myproject

go 1.21.6
		`
	packageFile := folderName + "/go.mod"
	err = ioutil.WriteFile(packageFile, []byte(packageMOD), 0644)
	if err != nil {
		return "Error writing package.json file:", err
	}

	// Create Dockerfile
	dockerfile := `FROM golang:latest
	WORKDIR /app
	COPY . .
	RUN go build -o function
	CMD ["./function"]`
	dockerfileLocation := folderName + "/Dockerfile"
	err = ioutil.WriteFile(dockerfileLocation, []byte(dockerfile), 0644)
	if err != nil {
		return "Error writing Dockerfile:", err
	}

	// Build Docker image
	cmd := exec.Command("docker", "build", "-t", "tziviarot/function_image_go:latest", folderName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return "Error building Docker image:", err
	}

	//push Docker image
	cmd = exec.Command("docker", "push", "tziviarot/function_image_go:latest")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return "Error phsh Docker image:",err
	}

	createAndRunYmlFileRes, err := createAndRunYmlFile("go", examples, "tziviarot/function_image_go:latest")

	if err != nil {
		return "Error create And Run Yml FileRes",err
	}

	os.RemoveAll(folderName)
	return createAndRunYmlFileRes
}
