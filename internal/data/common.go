package data

type ErrorResponse struct {
	Status     int    `json:"status"`
	StatusText string `json:"statusText"`
	Error      string `json:"error"`
}
