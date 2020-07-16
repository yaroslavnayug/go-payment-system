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
	"github.com/yaroslavnayug/go-payment-system/internal/api"
	e2e "github.com/yaroslavnayug/go-payment-system/test/utils"
)

var host = "http://localhost:8080"

func TestMain(m *testing.M) {
	if len(os.Getenv("HOST")) > 0 {
		host = os.Getenv("HOST")
	}
	code := m.Run()
	os.Exit(code)
}

func TestCreateAccount(t *testing.T) {
	// arrange
	var requestBody = []byte(fmt.Sprintf(`{
	"first_name": "foo",
	"last_name": "too",
	"passport_data": "%d",
	"phone": "+7904",
	"country": "Russia",
	"region": "Tatarstan",
	"city": "Kazan",
	"street": "Marks"
}`, e2e.GeneratePassportData()))

	request, err := http.NewRequest("POST", fmt.Sprintf("%s/createAccount", host), bytes.NewBuffer(requestBody))
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

	responseJSON := api.Response{}
	_ = json.Unmarshal(body, &responseJSON)

	// assert
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, http.StatusText(http.StatusOK), responseJSON.StatusMessage)
	assert.Contains(t, responseJSON.Response, "account_id")
}
