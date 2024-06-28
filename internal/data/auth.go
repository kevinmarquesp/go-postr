package data

type RegisterCredentialsIncome struct {
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterSuccessfulSessionTokenResponse struct {
	Username     string `json:"username"`
	SessionToken string `json:"sessionToken"`
}
