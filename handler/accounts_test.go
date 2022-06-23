package handler

import (
	"encoding/json"
	"fmt"
	db "kidsloop/account-service/database"
	api_errors "kidsloop/account-service/errors"
	"kidsloop/account-service/model"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"kidsloop/account-service/test_util"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var router *gin.Engine

func TestMain(m *testing.M) {
	test_util.LoadTestEnv("../")

	err := db.InitDB()
	if err != nil {
		log.Println("Failed to connect to postgres:")
		log.Fatal(err)
	}

	router = SetUpRouter()

	code := m.Run()
	os.Exit(code)
}

func TestCreateAccount200(t *testing.T) {
	request, _ := http.NewRequest("POST", "/accounts", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	var data model.CreateAccountResponse
	_ = json.Unmarshal([]byte(response.Body.Bytes()), &data)

	assert.Equal(t, 200, response.Code)
	assert.NotEqual(t, "", data.Account.ID)
	assert.NotEqual(t, "", data.Android.ID)
	assert.NotEqual(t, "", data.AndroidGroup.ID)
}

func TestGetAccount200(t *testing.T) {
	account, _ := db.Database.CreateAccount(nil)

	url := fmt.Sprintf("/accounts/%s", account.ID)
	request, _ := http.NewRequest("GET", url, nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	var data model.Account
	_ = json.Unmarshal([]byte(response.Body.Bytes()), &data)

	assert.Equal(t, 200, response.Code)
	assert.Equal(t, account.ID, data.ID)
}

func TestGetAccount400(t *testing.T) {
	url := fmt.Sprintf("/accounts/%s", "not-a-uuid")
	request, _ := http.NewRequest("GET", url, nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	var data model.ErrorResponse
	_ = json.Unmarshal([]byte(response.Body.Bytes()), &data)

	assert.Equal(t, 400, response.Code)
	assert.Equal(t, len(data.Errors), 1)

	err := data.Errors[0]
	assert.Equal(t, err.Code, api_errors.ErrCodeBadRequest)
	assert.Equal(t, err.Message, api_errors.ErrMsgBadRequest)
}

func TestGetAccount404(t *testing.T) {
	id := uuid.New().String()
	url := fmt.Sprintf("/accounts/%s", id)
	request, _ := http.NewRequest("GET", url, nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	var data model.ErrorResponse
	_ = json.Unmarshal([]byte(response.Body.Bytes()), &data)

	assert.Equal(t, 404, response.Code)
	assert.Equal(t, len(data.Errors), 1)

	err := data.Errors[0]
	assert.Equal(t, err.Code, api_errors.ErrCodeNotFound)
	assert.Equal(t, err.Message, fmt.Sprintf(api_errors.ErrMsgNotFound, "account", id))
}

func TestDeleteAccount200(t *testing.T) {
	account, _ := db.Database.CreateAccount(nil)

	url := fmt.Sprintf("/accounts/%s", account.ID)
	request, _ := http.NewRequest("DELETE", url, nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	var data model.Account
	_ = json.Unmarshal([]byte(response.Body.Bytes()), &data)

	assert.Equal(t, response.Code, 200)
	assert.Equal(t, response.Body.String(), "Success")
}

func TestDeleteAccount400(t *testing.T) {
	url := fmt.Sprintf("/accounts/%s", "not-a-uuid")
	request, _ := http.NewRequest("DELETE", url, nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	var data model.ErrorResponse
	_ = json.Unmarshal([]byte(response.Body.Bytes()), &data)

	assert.Equal(t, 400, response.Code)
	assert.Equal(t, len(data.Errors), 1)

	err := data.Errors[0]
	assert.Equal(t, err.Code, api_errors.ErrCodeBadRequest)
	assert.Equal(t, err.Message, api_errors.ErrMsgBadRequest)
}

func TestDeleteAccount404(t *testing.T) {
	id := uuid.New().String()
	url := fmt.Sprintf("/accounts/%s", id)
	request, _ := http.NewRequest("DELETE", url, nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	var data model.ErrorResponse
	_ = json.Unmarshal([]byte(response.Body.Bytes()), &data)

	assert.Equal(t, 404, response.Code)
	assert.Equal(t, len(data.Errors), 1)

	err := data.Errors[0]
	assert.Equal(t, err.Code, api_errors.ErrCodeNotFound)
	assert.Equal(t, err.Message, fmt.Sprintf(api_errors.ErrMsgNotFound, "account", id))
}
