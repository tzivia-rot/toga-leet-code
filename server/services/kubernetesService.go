package exerciseService

import (
	model "go-lenguage/models"
	"time"

	"fmt"
	"os"
	"os/exec"
	"strings"
)

func createYAMLString(language string, examples []model.Example, imageName string, UUID string) (string, error) {
	yamlContent := ""
	for i, example := range examples {
		yamlContent += fmt.Sprintf(`
apiVersion: batch/v1
kind: Job
metadata:
  name: funcc%d%s
spec:
  template:
    spec:
      containers:
      - name: function
        image: %s
        env:`, i+1, UUID, imageName)

		// Set environment variables for each input in the example
		for j, input := range example.Input {
			yamlContent += fmt.Sprintf(`
        - name: MY_VARIABLE_INPUT_%d
          value: "%s"`, j+1, input)
		}

		// Set environment variable for output
		yamlContent += fmt.Sprintf(`
        - name: MY_VARIABLE_OUTPUT
          value: "%s"`, example.Output)

		yamlContent += `
      restartPolicy: Never
---
`
	}
	return yamlContent, nil
}
func waitForPodReady(podName string) (error, string) {
	fmt.Println("Waiting for pod", podName, "to be ready...")
	for {
		cmd := exec.Command("kubectl", "get", "pod", podName, "-o", "jsonpath='{.status.phase}'")
		output, err := cmd.CombinedOutput()
		if err != nil {
			return err, "nil"
		}
		podPhase := strings.Trim(string(output), "'")
		if podPhase == "Succeeded" {
			fmt.Println("Pod", podName, "is now ready")
			return nil, "Succeeded"
		}
		if podPhase == "Failed" {
			fmt.Println("Pod", podName, "is now ready")
			return nil, "Failed"
		}
		time.Sleep(5 * time.Second) // Wait for 5 seconds before checking again
	}
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

func compareOutputs(examples []model.Example, lenguage string, UUID string) (bool, error) {
	var podStatus string
	for i, arr := range examples {
		podName, err := runCommand(fmt.Sprintf("kubectl get pods --selector=job-name=funcc%d%s -o=jsonpath='{.items[0].metadata.name}'", i+1, UUID))
		if err != nil {
			return false, err
		}
		podNameWithoutHyphen := strings.Replace(podName, "'", "", -1)

		// Wait for the pod to be ready
		if err, podStatus = waitForPodReady(podNameWithoutHyphen); err != nil {
			return false, err
		}
		fmt.Print(arr)
		// podStatus, err := runCommand(fmt.Sprintf("kubectl get pod %s -o jsonpath='{.status.containerStatuses[0].state.terminated.exitCode}'", podNameWithoutHyphen))
		fmt.Print("\npodStatus\n", podStatus)
		// logs, err := runCommand(fmt.Sprintf("kubectl logs %s", podNameWithoutHyphen))
		// fmt.Print("\nlog--\n", logs)
		// if err != nil {
		// 	return false, err
		// }
		// fmt.Print("\narr\n", arr)
		if podStatus != "Succeeded" {
			return false, nil
		}
		// if lenguage == "go" {
		// 	if !strings.EqualFold(logs, arr.Output) {
		// 		return false, nil
		// 	}
		// }
		// if lenguage == "node.js" {
		// 	result := strings.Split(logs, " ")
		// 	if result[len(result)] != arr.Output {
		// 		return false, nil
		// 	}
		// }

	}
	return true, nil
}

func createAndRunJobs(lenguage string, examples []model.Example, imageName string, UUID string) (string, error) {
	yamlContent, err := createYAMLString(lenguage, examples, imageName, UUID)
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

	outputEqual, err := compareOutputs(examples, lenguage, UUID)
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
