package data

// Used in the refresh session token related routes.

type RefreshUserSessionTokenCredentialsBody struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	SessionToken string `json:"sessionToken"`
}

type RefreshUserSessionTokenResponse struct {
	NewSessionToken string `json:"newSessionToken"`
}
