// +build e2e

package e2e

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yaroslavnayug/go-payment-system/internal/handler/v1.0"
)

var host = "http://localhost:8080"

var customerID = ""

func TestMain(m *testing.M) {
	if len(os.Getenv("HOST")) > 0 {
		host = os.Getenv("HOST")
	}
	code := m.Run()
	os.Exit(code)
}

func TestCreate(t *testing.T) {
	// arrange
	var requestBody = []byte(`{
		"first_name": "foo",
		"last_name": "too",
		"phone": "+7993",
		"address": {
			"country": "R",
			"region": "R",
			"city": "R",
			"street": "R",
			"building": "105"
		},
		"passport": {
			"number": "9876543210",
			"birth_date": "01-01-2000",
			"birth_place": "R",
			"issuer": "MMM",
			"issue_date": "01-01-2000"
		}
	}`)

	request, err := http.NewRequest("POST", fmt.Sprintf("%s/customer", host), bytes.NewBuffer(requestBody))
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	client := &http.Client{}

	// act
	response, err := client.Do(request)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	responseJSON := v1.CustomerBody{}
	_ = json.Unmarshal(body, &responseJSON)

	// assert
	assert.Equal(t, http.StatusCreated, response.StatusCode)
	assert.NotNil(t, responseJSON.CustomerID)

	customerID = responseJSON.CustomerID
}

func TestFind(t *testing.T) {
	// arrange
	request, err := http.NewRequest("GET", fmt.Sprintf("%s/customer/%s", host, customerID), bytes.NewBuffer(nil))
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	client := &http.Client{}

	// act
	response, err := client.Do(request)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	responseJSON := v1.CustomerBody{}
	_ = json.Unmarshal(body, &responseJSON)

	// assert
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.NotNil(t, responseJSON.CustomerID)
}

func TestUpdate(t *testing.T) {
	// arrange
	var requestBody = []byte(`{
		"first_name": "foo2",
		"last_name": "too2",
		"phone": "+79932",
		"address": {
			"country": "R2",
			"region": "R2",
			"city": "R2",
			"street": "R2",
			"building": "1052"
		},
		"passport": {
			"number": "9876543210",
			"birth_date": "01-01-2001",
			"birth_place": "R2",
			"issuer": "MMM2",
			"issue_date": "01-01-2020"
		}
	}`)

	request, err := http.NewRequest("PUT", fmt.Sprintf("%s/customer/%s", host, customerID), bytes.NewBuffer(requestBody))
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	client := &http.Client{}

	// act
	response, err := client.Do(request)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	responseJSON := v1.CustomerBody{}
	_ = json.Unmarshal(body, &responseJSON)

	// assert
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, []byte{}, body)
}

func TestDelete(t *testing.T) {
	// arrange
	request, err := http.NewRequest("DELETE", fmt.Sprintf("%s/customer/%s", host, customerID), bytes.NewBuffer(nil))
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	client := &http.Client{}

	// act
	response, err := client.Do(request)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	// assert
	assert.Equal(t, http.StatusNoContent, response.StatusCode)
	assert.Equal(t, []byte{}, body)
}
