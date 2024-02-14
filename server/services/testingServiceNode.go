package exerciseService

import (
	"fmt"
	model "go-lenguage/models"
	"io/ioutil"
	"os"
	"os/exec"
)

func CheckExerciseNode(function string, examples []model.Example, UUID string, imageName string) (string, error) {

	err := createFolderGo(UUID)
	if err != nil {
		return "error create folder: ", err
	}

	err = WriteFunctionGOContents(UUID, function)
	if err != nil {
		return "error create folder: ", err
	}

	err = CreatePackageJsonFile(UUID)
	if err != nil {
		return "error create folder: ", err
	}

	err = CreateGODockerfile(UUID)
	if err != nil {
		return "error create folder: ", err
	}

	err = BuildDockerImageGO(imageName, UUID)
	if err != nil {
		return "error create folder: ", err
	}

	err = phshDockerImage(imageName)
	if err != nil {
		return "error create folder: ", err
	}

	createAndRunYmlFileRes, err := createAndRunJobs("nodejs", examples, imageName, UUID)

	if err != nil {
		fmt.Print(err)
	}
	os.RemoveAll(UUID)
	return createAndRunYmlFileRes, nil

}
func createFolderGo(UUID string) error {
	folderName := UUID
	err := os.Mkdir(folderName, 0755)
	if err != nil {
		fmt.Println("Error creating folder:", err)
		return err
	}
	return nil

}
func WriteFunctionGOContents(UUID string, function string) error {
	functionFile := UUID + "/function.js"
	err := ioutil.WriteFile(functionFile, []byte(function), 0644)
	if err != nil {
		fmt.Println("Error writing function file:", err)
		return err
	}
	return nil
}
func CreatePackageJsonFile(UUID string) error {
	packageJSON := `{
		"name": "function_container",
		"version": "1.0.0",
		"scripts": {
			"start": "node function.js"
		}
	}`
	packageFile := UUID + "/package.json"
	err := ioutil.WriteFile(packageFile, []byte(packageJSON), 0644)
	if err != nil {
		fmt.Println("Error writing package.json file:", err)
		return err
	}
	return nil
}
func CreateGODockerfile(UUID string) error {
	dockerfile := `FROM node:latest
	WORKDIR /app
	COPY . .
	CMD ["npm", "start"]`
	dockerfileLocation := UUID + "/Dockerfile"
	err := ioutil.WriteFile(dockerfileLocation, []byte(dockerfile), 0644)
	if err != nil {
		fmt.Println("Error writing Dockerfile:", err)
		return err
	}
	return nil

}
func BuildDockerImageGO(imageName string, UUID string) error {
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
func phshDockerImage(imageName string) error {
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
