package data

type UpdateUserProfileDetailsBody struct {
	Fullname     string `json:"fullname"`
	Username     string `json:"username"`
	Description  string `json:"description"`
	SessionToken string `json:"sessionToken"`
}

type UpdateUserProfileDetailsResponse struct {
	Fullname    string `json:"fullname"`
	Username    string `json:"username"`
	Description string `json:"description"`
}
