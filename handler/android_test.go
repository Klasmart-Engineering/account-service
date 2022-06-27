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

	url := fmt.Sprintf("/android_groups/%s/androids", androidGroup.ID)
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

func TestCreateAndroid404(t *testing.T) {
	wrongUUID := uuid.New().String()
	url := fmt.Sprintf("/android_groups/%s/androids", wrongUUID)
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

func TestGetPaginatedAndroidsByGroup200(t *testing.T) {
	account, _ := db.Database.CreateAccount(nil)
	androidGroup, _ := db.Database.CreateAndroidGroup(nil, account.ID)
	android, _ := db.Database.CreateAndroid(nil, androidGroup.ID)

	url := fmt.Sprintf("/android_groups/%s/androids", androidGroup.ID)
	request, _ := http.NewRequest("GET", url, nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	var data []model.Android
	_ = json.Unmarshal([]byte(response.Body.Bytes()), &data)

	assert.Equal(t, 200, response.Code)
	assert.Equal(t, 1, len(data))
	assert.Equal(t, android.ID, data[0].ID)
}
