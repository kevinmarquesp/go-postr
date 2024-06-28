package data

type ErrorResponse struct {
	StatusText string `json:"statusText"`
	Error      string `json:"error"`
}
