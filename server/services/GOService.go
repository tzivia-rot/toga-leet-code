package exerciseService

import (
	"fmt"
	model "go-lenguage/models"
	"io/ioutil"
	"os"
	"os/exec"
)

func CheckExerciseGO(function string, examples []model.Example) string {

	folderName := "function_folder"
	err := os.Mkdir(folderName, 0755)
	if err != nil {
		fmt.Println("Error creating folder:", err)
	}
	// Write function contents to a go file
	functionFile := folderName + "/main.go"
	err = ioutil.WriteFile(functionFile, []byte(function), 0644)
	if err != nil {
		fmt.Println("Error writing function file:", err)
	}
	fmt.Print("\nWrite function contents to a go file\n")

	// Create go.mod file
	packageMOD := `module myproject

go 1.21.6
		`
	packageFile := folderName + "/go.mod"
	err = ioutil.WriteFile(packageFile, []byte(packageMOD), 0644)
	if err != nil {
		fmt.Println("Error writing package.json file:", err)
	}
	fmt.Print("\nWrite mod\n")

	// Create Dockerfile
	dockerfile := `FROM golang:latest
	WORKDIR /app
	COPY . .
	RUN go build -o function
	CMD ["./main"]`
	dockerfileLocation := folderName + "/Dockerfile"
	err = ioutil.WriteFile(dockerfileLocation, []byte(dockerfile), 0644)
	if err != nil {
		fmt.Println("Error writing Dockerfile:", err)
	}
	fmt.Print("\nWrite docker")

	// Build Docker image
	cmd := exec.Command("docker", "build", "-t", "tziviarot/function_image_go:latest", folderName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error building Docker image:", err)
	}

	//push Docker image
	cmd = exec.Command("docker", "push", "tziviarot/function_image_go:latest")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println("\nError phsh Docker image:", err)
	}

	createAndRunYmlFile, err := createAndRunYmlFile(examples, "tziviarot/function_image_go:latest")

	if err != nil {
		fmt.Print("err:", err)
	}

	os.RemoveAll(folderName)
	return createAndRunYmlFile
}
