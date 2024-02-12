package exerciseService

import (
	model "go-lenguage/models"

	"fmt"
	"os"
	"os/exec"
	"strings"
)

// type Example struct {
// 	Input  []string `json:"input"`
// 	Output string   `json:"output"`
// }

// type Exercise struct {
// 	ID             string `json:"id" bson:"_id,omitempty"`
// 	Name           string `json:"name" bson:"name"`
// 	Description    string `json:"description" bson:"description"`
// 	Examples       []Example
// 	BasisOperation string
// }
// type Answer struct {
// 	Function   string `json:"function" bson:"function"`
// 	Lenguage   string `json:"lenguage" bson:"lenguage"`
// 	ID         string `json:"id" bson:"_id,omitempty"`
// 	ExerciseId string `json:"exerciseId" bson:"exerciseId"`
// }

func createYAMLString(arrays []model.Example, imageName string) (string, error) {
	yamlContent := ""
	for i, arr := range arrays {
		yamlContent += fmt.Sprintf(`---
apiVersion: batch/v1
kind: Job
metadata:
  name: functionex-%d
spec:
  template:
    spec:
      containers:
      - name: function
        image: %s
        env:`, i+1, imageName)

		// Set environment variables for each input in the array
		for j, input := range arr.Input {
			yamlContent += fmt.Sprintf(`
        - name: MY_VARIABLE_%d
          value: "%s"`, j+1, input)
		}

		yamlContent += `
      restartPolicy: Never`
	}
	return yamlContent, nil
}

func runCommand(command string) (string, error) {
	fmt.Print("cooo", command)
	cmd := exec.Command("cmd", "/C", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Print("errr", err)
		return "", err
	}
	return string(output), nil
}

func compareOutputs(examples []model.Example) (bool, error) {
	fmt.Print("\nfdfdf1", examples)

	for i, arr := range examples {
		podName, err := runCommand(fmt.Sprintf("kubectl get pods -l app=function-%d -o jsonpath='{.items[0].metadata.name}'", i+1))
		if err != nil {
			return false, err
		}
		// fmt.Print("kubectl logssssssssssssss ", podName)
		fmt.Print("\nfdfdf2\n", podName)
		podNameWithoutHyphen := strings.Replace(podName, "'", "", -1)
		logs, err := runCommand(fmt.Sprintf("kubectl logs %s", podNameWithoutHyphen))
		fmt.Print("\nfdfdf2\n", logs)

		if err != nil {
			fmt.Print("\nerrrrrrrrr\n")
			// return false, err
		}

		if logs != arr.Output {
			fmt.Print()
			return false, nil
		}
	}
	return true, nil
}

func createAndRunYmlFile(examples []model.Example, imageName string) (string, error) {
	yamlContent, err := createYAMLString(examples, imageName)
	fmt.Print("\nimageNameymaelll-----------------\n", yamlContent)
	fmt.Print("\nanddddddddddd")
	if err != nil {
		fmt.Println("Error creating YAML string:", err)
		return "err", err
	}

	// Write YAML content to a temporary file
	tmpfile, err := os.CreateTemp("", "temp-*.yaml")
	if err != nil {
		fmt.Println("Error creating temporary file:", err)
		return "err", err
	}

	defer os.Remove(tmpfile.Name()) // Clean up the temporary file
	_, err = tmpfile.WriteString(yamlContent)
	if err != nil {
		fmt.Println("Error writing to temporary file:", err)
		return "err", err
	}
	tmpfile.Close()

	// Apply YAML from the temporary file
	fmt.Print("tmpfile.Name())))))))))", tmpfile.Name())
	_, err = runCommand(fmt.Sprintf("kubectl apply -f %s", tmpfile.Name()))
	if err != nil {
		fmt.Println("Error applying YAML from temporary file:", err)
		return "err", err
	}

	outputEqual, err := compareOutputs(examples)
	if err != nil {
		fmt.Println("Error comparing outputs:", err)
		return "err", err
	}

	if !outputEqual {
		fmt.Println("Output does not match expected output")
		return "Output does not match expected output:(", err
	}

	fmt.Println("Output matches expected output")
	return "Output matches expected output", err
}
