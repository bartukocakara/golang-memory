package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

var API_PORT string
var BASE_URL string

func init() {
	API_PORT = os.Getenv("API_PORT")
	if API_PORT == "" {
		API_PORT = "8080"
	}
	BASE_URL = fmt.Sprintf("http://localhost:%s", API_PORT)
}

func TestSet(t *testing.T) {
	requestBody := map[string]string{
		"key":   "test-key",
		"value": "test-value",
	}

	wantStatus := 201
	bodyData, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal("TestSet Error when json.Marshal: ", err)
	}
	bufferData := bytes.NewBuffer(bodyData)

	req, err := http.NewRequest(http.MethodPost, BASE_URL+"/set", bufferData)
	if err != nil {
		t.Fatal("TestSet Error when NewRequest: ", err)
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal("TestSet Error when client.Do: ", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != wantStatus {
		t.Fatalf("Expected to get status %d but instead got %d", wantStatus, resp.StatusCode)
	}

	t.Log("Test Success.")
}

func TestSetKeyError(t *testing.T) {
	requestBody := map[string]string{
		"value": "test-value",
	}

	wantStatus := 400
	wantError := "The 'key' is required."
	var response map[string]string
	bodyData, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal("TestSet Error when json.Marshal: ", err)
	}
	bufferData := bytes.NewBuffer(bodyData)

	req, err := http.NewRequest(http.MethodPost, BASE_URL+"/set", bufferData)
	if err != nil {
		t.Fatal("TestSet Error when NewRequest: ", err)
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal("TestSet Error when client.Do: ", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != wantStatus {
		t.Fatalf("Expected to get status %d but instead got %d", wantStatus, resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("TestSet Error when ioutil.ReadAll: ", err)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Fatal("TestSet Error when response unmarshall: ", err)
	}

	if response["error"] != wantError {
		t.Fatalf("Expected to get error %s but instead got %s", wantError, response["error"])
	}

	t.Log("Test Success.")
}

func TestSetValueError(t *testing.T) {
	requestBody := map[string]string{
		"key": "test-key",
	}

	wantStatus := 400
	wantError := "The 'value' is required."
	var response map[string]string

	bodyData, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal("TestSet Error when json.Marshal: ", err)
	}
	bufferData := bytes.NewBuffer(bodyData)
	req, err := http.NewRequest(http.MethodPost, BASE_URL+"/set", bufferData)
	if err != nil {
		t.Fatal("TestSet Error when NewRequest: ", err)
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal("TestSet Error when client.Do: ", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != wantStatus {
		t.Fatalf("Expected to get status %d but instead got %d", wantStatus, resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("TestSet Error when ioutil.ReadAll: ", err)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Fatal("TestSet Error when response unmarshall: ", err)
	}

	if response["error"] != wantError {
		t.Fatalf("Expected to get error %s but instead got %s", wantError, response["error"])
	}

	t.Log("Test Success.")
}

func TestSetMethodNotAllowed(t *testing.T) {
	requestParams := map[string]string{
		"key":   "test-key",
		"value": "test-value",
	}

	wantStatus := 405
	req, err := http.NewRequest(http.MethodDelete, BASE_URL+"/set", nil)
	if err != nil {
		t.Fatal("TestSet Error when NewRequest: ", err)
	}
	q := req.URL.Query()
	for key, value := range requestParams {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal("TestSet Error when client.Do: ", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != wantStatus {
		t.Fatalf("Expected to get status %d but instead got %d", wantStatus, resp.StatusCode)
	}

	t.Log("Test Success.")
}

func TestGet(t *testing.T) {
	requestParams := map[string]string{
		"key": "test-key",
	}
	want := "test-value"
	wantStatus := 200
	var response map[string]string
	req, err := http.NewRequest(http.MethodGet, BASE_URL+"/get", nil)
	if err != nil {
		t.Fatal("TestSet Error when NewRequest: ", err)
	}
	q := req.URL.Query()
	for key, value := range requestParams {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal("TestSet Error when client.Do: ", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != wantStatus {
		t.Fatalf("Expected to get status %d but instead got %d", wantStatus, resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("TestSet Error when ioutil.ReadAll: ", err)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Fatal("TestSet Error when response unmarshall: ", err)
	}

	if response["result"] != want {
		t.Fatalf("Expected to get result %s but instead got %s", want, response["result"])
	}

	t.Log("Test Success.")
}

func TestGetNotFoundError(t *testing.T) {
	requestParams := map[string]string{
		"key": "test-keys",
	}
	wantError := "The key 'test-keys' could not be found."
	wantStatus := 404
	var response map[string]string
	req, err := http.NewRequest(http.MethodGet, BASE_URL+"/get", nil)
	if err != nil {
		t.Fatal("TestSet Error when NewRequest: ", err)
	}
	q := req.URL.Query()
	for key, value := range requestParams {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal("TestSet Error when client.Do: ", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != wantStatus {
		t.Fatalf("Expected to get status %d but instead got %d", wantStatus, resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("TestSet Error when ioutil.ReadAll: ", err)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Fatal("TestSet Error when response unmarshall: ", err)
	}

	if response["error"] != wantError {
		t.Fatalf("Expected to get error %s but instead got %s", wantError, response["error"])
	}

	t.Log("Test Success.")
}

func TestHome(t *testing.T) {
	wantStatus := 200
	req, err := http.NewRequest(http.MethodGet, BASE_URL+"/", nil)
	if err != nil {
		t.Fatal("TestSet Error when NewRequest: ", err)
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal("TestSet Error when client.Do: ", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != wantStatus {
		t.Fatalf("Expected to get status %d but instead got %d", wantStatus, resp.StatusCode)
	}
	t.Log("Test Success.")
}

func TestFlush(t *testing.T) {
	wantStatus := 204
	req, err := http.NewRequest(http.MethodDelete, BASE_URL+"/flush", nil)
	if err != nil {
		t.Fatal("TestSet Error when NewRequest: ", err)
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal("TestSet Error when client.Do: ", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != wantStatus {
		t.Fatalf("Expected to get status %d but instead got %d", wantStatus, resp.StatusCode)
	}
	t.Log("Test Success.")
}
