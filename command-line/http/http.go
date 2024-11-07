package httpCall

import (
	"bytes"
	"command-line/modules"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var url = "http://localhost:8080/exercises"

func AddExercise(exercise modules.Exercise) (string, error) {
	postBody, _ := json.Marshal(exercise)
	responseBody := bytes.NewBuffer(postBody)
	response, err := http.Post(url+"/create", "application/json", responseBody)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func DeleteExercise(exerciseID string) (string, error) {
	request, err := http.NewRequest("DELETE", url+"/delete/"+exerciseID, nil)
	if err != nil {
		return "", err
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	return response.Status, nil
}

func UpdateExercise(exerciseID string, data map[string]string) (string, error) {
	payload, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	request, err := http.NewRequest("PUT", url+"/update/"+exerciseID, bytes.NewBuffer(payload))
	if err != nil {
		return "", err
	}
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	return response.Status, nil
}

func GetExerciseByID(exerciseID string) (modules.Exercise, error) {
	var exercise modules.Exercise

	response, err := http.Get(url + "/getById/" + exerciseID)
	if err != nil {
		return exercise, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return exercise, fmt.Errorf("Failed to get item by ID. Status code: %d", response.StatusCode)
	}

	err = json.NewDecoder(response.Body).Decode(&exercise)
	if err != nil {
		return exercise, fmt.Errorf("Failed to get item by ID.")
	}

	return exercise, nil
}

func GetAllExercises() ([]modules.Exercise, error) {
	var exercises []modules.Exercise

	response, err := http.Get(url + "/getAll")
	if err != nil {
		return exercises, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return exercises, fmt.Errorf("Failed to get items. Status code: %d", response.StatusCode)
	}

	err = json.NewDecoder(response.Body).Decode(&exercises)
	if err != nil {
		return exercises, err
	}

	return exercises, nil
}

func CheckExercise(id string, functionCode, lenguage string) (string, error) {
	postBody, _ := json.Marshal(map[string]string{
		"function": functionCode,
		"lenguage": lenguage,
	})
	responseBody := bytes.NewBuffer(postBody)
	response, err := http.Post(url+"/check/"+id, "application/json", responseBody)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
