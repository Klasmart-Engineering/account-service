package handler

import (
	"context"
	"encoding/json"
	"fmt"
	db "kidsloop/account-service/database"
	api_errors "kidsloop/account-service/errors"
	"kidsloop/account-service/model"
	"kidsloop/account-service/test_util"
	"net/http"
	"net/http/httptest"
	"sort"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateAndroid200(t *testing.T) {
	ctx := context.Background()
	account, _ := db.Database.CreateAccount(nil, ctx)
	androidGroup, _ := db.Database.CreateAndroidGroup(nil, ctx, account.ID)

	url := fmt.Sprintf("/android_groups/%s/androids", androidGroup.ID)
	request, _ := http.NewRequest("POST", url, nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	var data model.Android
	_ = json.Unmarshal([]byte(response.Body.Bytes()), &data)

	assert.Equal(t, 200, response.Code)
	assert.True(t, test_util.IsValidUUID(data.ID))
	assert.Equal(t, androidGroup.ID, data.AndroidGroupID)

	android, err := db.Database.GetAndroid(nil, ctx, data.ID)
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

func TestGetAndroidsByGroup200(t *testing.T) {
	ctx := context.Background()
	account, _ := db.Database.CreateAccount(nil, ctx)
	androidGroup, _ := db.Database.CreateAndroidGroup(nil, ctx, account.ID)
	android, _ := db.Database.CreateAndroid(nil, ctx, androidGroup.ID)

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

func TestGetAndroidsByGroup200Pages(t *testing.T) {
	ctx := context.Background()
	account, _ := db.Database.CreateAccount(nil, ctx)
	androidGroup, _ := db.Database.CreateAndroidGroup(nil, ctx, account.ID)

	ids := make([]string, 0)
	for i := 0; i < 10; i++ {
		android, _ := db.Database.CreateAndroid(nil, ctx, androidGroup.ID)
		ids = append(ids, android.ID)
	}
	// pagination should order by ID by default
	sort.Strings(ids)

	firstPage := fmt.Sprintf("/android_groups/%s/androids?limit=5", androidGroup.ID)
	secondPage := fmt.Sprintf("/android_groups/%s/androids?offset=5", androidGroup.ID)

	for index, url := range []string{firstPage, secondPage} {
		request, _ := http.NewRequest("GET", url, nil)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		var data []model.Android
		_ = json.Unmarshal([]byte(response.Body.Bytes()), &data)

		assert.Equal(t, 200, response.Code)
		assert.Equal(t, 5, len(data))
		for i := 0; i < 5; i++ {
			assert.Equal(t, ids[i+index*5], data[i].ID)
		}
	}
}

func TestDeleteAndroid200(t *testing.T) {
	account, _ := db.Database.CreateAccount(nil, context.Background())
	androidGroup, _ := db.Database.CreateAndroidGroup(nil, context.Background(), account.ID)
	android, _ := db.Database.CreateAndroid(nil, context.Background(), androidGroup.ID)

	url := fmt.Sprintf("/androids/%s", android.ID)
	request, _ := http.NewRequest("DELETE", url, nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	var data model.Account
	_ = json.Unmarshal([]byte(response.Body.Bytes()), &data)

	assert.Equal(t, 200, response.Code)
	assert.Equal(t, android.ID, data.ID)
}

func TestDeleteAndroid400(t *testing.T) {
	url := fmt.Sprintf("/androids/%s", "not-a-uuid")
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

func TestDeleteAndroid404(t *testing.T) {
	id := uuid.New().String()
	url := fmt.Sprintf("/androids/%s", id)
	request, _ := http.NewRequest("DELETE", url, nil)
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
