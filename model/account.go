package model

type Account struct {
	ID string `json:"id"`
}

type Android struct {
	ID             string `json:"id"`
	AndroidGroupID string `json:"android_group_id"`
}

type AndroidGroup struct {
	ID string `json:"id"`
}

type APIError struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details"`
}

type CreateAccountResponse struct {
	Account      Account      `json:"account"`
	Android      Android      `json:"android"`
	AndroidGroup AndroidGroup `json:"android_group"`
}
