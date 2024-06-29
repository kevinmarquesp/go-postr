package data

type RegisterNewUserBodyCredentialsBody struct {
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterNewUserSuccessfulResponse struct {
	Username     string `json:"username"`
	SessionToken string `json:"sessionToken"`
}

type UpdateUserSessionTokenCredentialsBody struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	SessionToken string `json:"sessionToken"`
}

type UpdateUserSessionTokenResponse struct {
	NewSessionToken string `json:"newSessionToken"`
}
