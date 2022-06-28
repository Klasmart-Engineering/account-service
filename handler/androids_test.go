package handler

import (
	"context"
	"encoding/json"
	"fmt"
	db "kidsloop/account-service/database"
	api_errors "kidsloop/account-service/errors"
	"kidsloop/account-service/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetAndroid200(t *testing.T) {
	account, _ := db.Database.CreateAccount(nil, context.Background())
	androidGroup, _ := db.Database.CreateAndroidGroup(nil, context.Background(), account.ID)
	android, _ := db.Database.CreateAndroid(nil, context.Background(), androidGroup.ID)

	url := fmt.Sprintf("/androids/%s", android.ID)
	request, _ := http.NewRequest("GET", url, nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	var data model.Account
	_ = json.Unmarshal([]byte(response.Body.Bytes()), &data)

	assert.Equal(t, 200, response.Code)
	assert.Equal(t, android.ID, data.ID)
}

func TestGetAndroid400(t *testing.T) {
	url := fmt.Sprintf("/androids/%s", "not-a-uuid")
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

func TestGetAndroid404(t *testing.T) {
	id := uuid.New().String()
	url := fmt.Sprintf("/androids/%s", id)
	request, _ := http.NewRequest("GET", url, nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	var data model.ErrorResponse
	_ = json.Unmarshal([]byte(response.Body.Bytes()), &data)

	assert.Equal(t, 404, response.Code)
	assert.Equal(t, len(data.Errors), 1)

	err := data.Errors[0]
	assert.Equal(t, err.Code, api_errors.ErrCodeNotFound)
	assert.Equal(t, err.Message, fmt.Sprintf(api_errors.ErrMsgNotFound, "android", id))
}
