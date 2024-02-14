package exerciseService

import (
	"fmt"
	model "go-lenguage/models"
	"io/ioutil"
	"os"
	"os/exec"
)

func CheckExerciseGO(function string, examples []model.Example, UUID string, imageName string) (string, error) {

	err := createFolder(UUID)
	if err != nil {
		return "Error creating folder:", err
	}

	err = WriteFunctionContents(UUID, function)
	if err != nil {
		return "Error writing function file:", err
	}

	err = CreatePackageModFile(UUID)
	if err != nil {
		return "Error writing package.json file:", err
	}

	err = CreateDockerfile(UUID)
	if err != nil {
		return "Error writing Dockerfile:", err
	}

	err = BuildDockerImage(imageName, UUID)
	if err != nil {
		return "Error building Docker image:", err
	}

	err = PushDockerImage(imageName)
	if err != nil {
		os.RemoveAll(UUID)
		return "Error phsh Docker image:", err
	}

	createAndRunYmlFileRes, err := createAndRunJobs("go", examples, imageName, UUID)

	if err != nil {
		os.RemoveAll(UUID)
		return "Error create And Run Yml FileRes", err
	}

	os.RemoveAll(UUID)
	return createAndRunYmlFileRes, err
}

func createFolder(UUID string) error {
	folderName := UUID
	err := os.Mkdir(folderName, 0755)
	if err != nil {
		fmt.Println("Error creating folder:", err)
		return err
	}
	return nil
}

func WriteFunctionContents(folderName string, function string) error {
	functionFile := folderName + "/function.go"
	err := ioutil.WriteFile(functionFile, []byte(function), 0644)
	if err != nil {
		fmt.Println("Error writing function file:", err)
		return err
	}
	return nil
}

func CreatePackageModFile(folderName string) error {
	packageMOD := `module myproject

go 1.21.6
    `
	packageFile := folderName + "/go.mod"
	err := ioutil.WriteFile(packageFile, []byte(packageMOD), 0644)
	if err != nil {
		fmt.Println("Error writing package.json file:", err)
		return err
	}
	return nil
}

func CreateDockerfile(folderName string) error {
	dockerfile := `FROM golang:latest
    WORKDIR /app
    COPY . .
    RUN go build -o function
    CMD ["./function"]`
	dockerfileLocation := folderName + "/Dockerfile"
	err := ioutil.WriteFile(dockerfileLocation, []byte(dockerfile), 0644)
	if err != nil {
		fmt.Println("Error writing Dockerfile:", err)
		return err
	}
	return nil
}

func BuildDockerImage(imageName string, UUID string) error {
	cmd := exec.Command("docker", "build", "-t", imageName, UUID)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error building Docker image:", err)
		return err
	}
	return nil
}

func PushDockerImage(imageName string) error {
	cmd := exec.Command("docker", "push", imageName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("\nError phsh Docker image:", err)
		return err
	}
	return nil
}
