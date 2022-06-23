package handler

import (
	"encoding/json"
	"fmt"
	db "kidsloop/account-service/database"
	api_errors "kidsloop/account-service/errors"
	"kidsloop/account-service/model"
	"kidsloop/account-service/test_util"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateAndroid200(t *testing.T) {
	account, _ := db.Database.CreateAccount(nil)
	androidGroup, _ := db.Database.CreateAndroidGroup(nil, account.ID)

	url := fmt.Sprintf("/accounts/%s/android_groups/%s", account.ID, androidGroup.ID)
	request, _ := http.NewRequest("POST", url, nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	var data model.Android
	_ = json.Unmarshal([]byte(response.Body.Bytes()), &data)

	assert.Equal(t, 200, response.Code)
	assert.True(t, test_util.IsValidUUID(data.ID))
	assert.Equal(t, androidGroup.ID, data.AndroidGroupID)

	android, err := db.Database.GetAndroid(nil, data.ID)
	assert.Nil(t, err)
	assert.Equal(t, android.ID, data.ID)
}

func TestCreateAndroid404Account(t *testing.T) {
	account, _ := db.Database.CreateAccount(nil)
	androidGroup, _ := db.Database.CreateAndroidGroup(nil, account.ID)
	wrongUUID := uuid.New().String()

	url := fmt.Sprintf("/accounts/%s/android_groups/%s", wrongUUID, androidGroup.ID)
	request, _ := http.NewRequest("POST", url, nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	var data model.ErrorResponse
	_ = json.Unmarshal([]byte(response.Body.Bytes()), &data)

	assert.Equal(t, 404, response.Code)
	assert.Equal(t, len(data.Errors), 1)
	err := data.Errors[0]
	assert.Equal(t, err.Code, api_errors.ErrCodeNotFound)
	assert.Equal(t, err.Message, fmt.Sprintf(api_errors.ErrMsgNotFound, "account", wrongUUID))
}

func TestCreateAndroid404AndroidGroup(t *testing.T) {
	account, _ := db.Database.CreateAccount(nil)
	wrongUUID := uuid.New().String()

	url := fmt.Sprintf("/accounts/%s/android_groups/%s", account.ID, wrongUUID)
	request, _ := http.NewRequest("POST", url, nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	var data model.ErrorResponse
	_ = json.Unmarshal([]byte(response.Body.Bytes()), &data)

	assert.Equal(t, 404, response.Code)
	assert.Equal(t, len(data.Errors), 1)
	err := data.Errors[0]
	assert.Equal(t, err.Code, api_errors.ErrCodeNotFound)
	assert.Equal(t, err.Message, fmt.Sprintf(api_errors.ErrMsgNotFound, "android_group", wrongUUID))
}

func TestCreateAndroid404GroupInAccount(t *testing.T) {
	// query a valid android_group that belongs to a different account than the one specified
	account1, _ := db.Database.CreateAccount(nil)
	account2, _ := db.Database.CreateAccount(nil)
	androidGroup, _ := db.Database.CreateAndroidGroup(nil, account2.ID)

	url := fmt.Sprintf("/accounts/%s/android_groups/%s", account1.ID, androidGroup.ID)
	request, _ := http.NewRequest("POST", url, nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	var data model.ErrorResponse
	_ = json.Unmarshal([]byte(response.Body.Bytes()), &data)

	assert.Equal(t, 404, response.Code)
	assert.Equal(t, len(data.Errors), 1)
	err := data.Errors[0]
	assert.Equal(t, err.Code, api_errors.ErrCodeNotFound)
	assert.Equal(t, err.Message, fmt.Sprintf(api_errors.ErrMsgNotFound, "android_group", androidGroup.ID))
}
