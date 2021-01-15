package tests

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: Mock Server
func TestCreateTransactionInvalidPayload(t *testing.T) {
	body := map[string]interface{}{
		"ID": "invalidid",
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}

	response, err := http.Post("http://localhost:8080/create_transaction", "application/json", bytes.NewBuffer(jsonBody))
	assert.Nil(t, err)
	assert.NotNil(t, response)

	responseError, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}

	assert.EqualValues(t, http.StatusBadRequest, response.StatusCode)
	assert.EqualValues(t, "invalid payload\n", responseError)
}

func TestCreateTransactionInvalidAccount(t *testing.T) {
	body := map[string]interface{}{
		"ID":        "testuidt",
		"AccountID": "invalidaccount",
		"Amount":    1,
		"CreatedAt": 1,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}

	response, err := http.Post("http://localhost:8080/create_transaction", "application/json", bytes.NewBuffer(jsonBody))
	assert.Nil(t, err)
	assert.NotNil(t, response)

	responseError, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}

	assert.EqualValues(t, http.StatusInternalServerError, response.StatusCode)
	assert.EqualValues(t, "failed to fetch account\n", string(responseError))
}

func TestCreateTransactionTotalUpdated(t *testing.T) {
	body := map[string]interface{}{
		"ID":        "testuidt",
		"AccountID": "testaid1",
		"Amount":    10,
		"CreatedAt": 1,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}

	response, err := http.Post("http://localhost:8080/create_transaction", "application/json", bytes.NewBuffer(jsonBody))
	assert.Nil(t, err)
	assert.NotNil(t, response)

	assert.EqualValues(t, http.StatusCreated, response.StatusCode)

	userBody := map[string]string{
		"ID": "testuid",
	}
	useJsonBody, err := json.Marshal(userBody)

	if err != nil {
		t.Fatal(err)
	}
	request, err := http.NewRequest("GET", "http://localhost:8080/user", bytes.NewBuffer(useJsonBody))
	client := &http.Client{}
	response, err = client.Do(request)

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
	assert.EqualValues(t, 90, userResponse.Total)
}

func TestCreateTransaction(t *testing.T) {
	body := map[string]interface{}{
		"ID":        "testuidt",
		"AccountID": "testaid1",
		"Amount":    1,
		"CreatedAt": 1,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}

	response, err := http.Post("http://localhost:8080/create_transaction", "application/json", bytes.NewBuffer(jsonBody))
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.EqualValues(t, http.StatusCreated, response.StatusCode)
}
