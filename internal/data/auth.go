package data

// Used on the user register API related routes.

type RegisterNewUserBodyCredentialsBody struct {
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterNewUserSuccessfulResponse struct {
	Username     string `json:"username"`
	PublicID     string `json:"publicId"`
	SessionToken string `json:"sessionToken"`
}

// Used in the refresh session token related routes.

type RefreshUserSessionTokenCredentialsBody struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	SessionToken string `json:"sessionToken"`
}

type RefreshUserSessionTokenResponse struct {
	NewSessionToken string `json:"newSessionToken"`
}
