package tests

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type userResponse struct {
	ID    string  `json:"ID"`
	Name  string  `json:"Name"`
	Total float64 `json:"Total"`
}

func TestInvalidPayload(t *testing.T) {
	response, err := http.Get("http://localhost:8080/user")

	assert.Nil(t, err)
	assert.NotNil(t, response)

	responseError, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}
	assert.EqualValues(t, http.StatusBadRequest, response.StatusCode)
	assert.EqualValues(t, "invalid payload\n", responseError)
}

func TestUserNotFound(t *testing.T) {
	body := map[string]string{
		"ID": "invaliduser",
	}
	jsonBody, err := json.Marshal(body)

	if err != nil {
		t.Fatal(err)
	}
	request, err := http.NewRequest("GET", "http://localhost:8080/user", bytes.NewBuffer(jsonBody))
	client := &http.Client{}
	response, err := client.Do(request)

	assert.Nil(t, err)
	assert.NotNil(t, response)

	responseError, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}

	assert.EqualValues(t, http.StatusInternalServerError, response.StatusCode)
	assert.EqualValues(t, "failed to fetch user\n", responseError)
}

func TestGetUser(t *testing.T) {
	body := map[string]string{
		"ID": "testuid",
	}
	jsonBody, err := json.Marshal(body)

	if err != nil {
		t.Fatal(err)
	}
	request, err := http.NewRequest("GET", "http://localhost:8080/user", bytes.NewBuffer(jsonBody))
	client := &http.Client{}
	response, err := client.Do(request)

	assert.Nil(t, err)
	assert.NotNil(t, response)

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}

	var userResponse userResponse
	err = json.Unmarshal([]byte(responseBody), &userResponse)
	if err != nil {
		t.Fatal(err)
	}

	assert.EqualValues(t, http.StatusOK, response.StatusCode)
	assert.EqualValues(t, "testuid", userResponse.ID)
	assert.EqualValues(t, "testname", userResponse.Name)
	assert.EqualValues(t, 0, userResponse.Total)
}
