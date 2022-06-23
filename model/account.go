package model

import (
	api_errors "kidsloop/account-service/errors"
)

type Account struct {
	ID string `json:"id"`
}

type Android struct {
	ID             string `json:"id"`
	AndroidGroupID string `json:"android_group_id"`
}

type AndroidGroup struct {
	ID        string `json:"id"`
	AccountID string `json:"account_id"`
}

type CreateAccountResponse struct {
	Account      Account      `json:"account"`
	Android      Android      `json:"android"`
	AndroidGroup AndroidGroup `json:"android_group"`
}

type ErrorResponse struct {
	Errors []api_errors.APIErrorResponse `json:"errors"`
}
