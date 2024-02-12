package exerciseService

import (
	model "go-lenguage/models"

	"fmt"
	"os"
	"os/exec"
	"strings"
)

func createYAMLString(lenguage string, arrays []model.Example, imageName string) (string, error) {
	yamlContent := ""
	for i, arr := range arrays {
		yamlContent += fmt.Sprintf(`
apiVersion: batch/v1
kind: Job
metadata:
  name: function-%d%s
spec:
  template:
    spec:
      containers:
      - name: function
        image: %s
        env:`, i+1, lenguage, imageName)

		// Set environment variables for each input in the array
		for j, input := range arr.Input {
			yamlContent += fmt.Sprintf(`
        - name: MY_VARIABLE_%d
          value: "%s"`, j+1, input)
		}

		yamlContent += `
      restartPolicy: Never
---
`
	}
	return yamlContent, nil
}

func runCommand(command string) (string, error) {
	fmt.Print("\ncomnand:\n", command)
	cmd := exec.Command("cmd", "/C", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Print("errr", err)
		return "", err
	}
	return string(output), nil
}

func compareOutputs(examples []model.Example, lenguage string) (bool, error) {
	for i, arr := range examples {
		podName, err := runCommand(fmt.Sprintf("kubectl get pods --selector=job-name=function-%d%s -o=jsonpath='{.items[0].metadata.name}'", i+1, lenguage))
		if err != nil {
			return false, err
		}
		podNameWithoutHyphen := strings.Replace(podName, "'", "", -1)
		logs, err := runCommand(fmt.Sprintf("kubectl logs %s", podNameWithoutHyphen))
		fmt.Print("\nlog--\n", logs)
		if err != nil {
			return false, err
		}
		if lenguage == "go" {
			if !strings.EqualFold(logs, arr.Output) {
				return false, nil
			}
		}
		if lenguage == "node.js" {
			result := strings.Split(logs, " ")
			if result[len(result)] != arr.Output {
				return false, nil
			}
		}

	}
	return true, nil
}

func createAndRunYmlFile(lenguage string, examples []model.Example, imageName string) (string, error) {
	yamlContent, err := createYAMLString(lenguage, examples, imageName)
	fmt.Print("\nyamlContent", yamlContent)
	if err != nil {
		fmt.Println("Error creating YAML string:", err)
		return "err", err
	}

	tmpfile, err := os.CreateTemp("", "temp-*.yaml")
	if err != nil {
		fmt.Println("Error creating temporary file:", err)
		return "err", err
	}

	defer os.Remove(tmpfile.Name())
	_, err = tmpfile.WriteString(yamlContent)
	if err != nil {
		fmt.Println("Error writing to temporary file:", err)
		return "err", err
	}
	tmpfile.Close()

	_, err = runCommand(fmt.Sprintf("kubectl apply -f %s ", tmpfile.Name()))
	if err != nil {
		fmt.Println("Error applying YAML from temporary file:", err)
		return "err", err
	}

	outputEqual, err := compareOutputs(examples, lenguage)
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
