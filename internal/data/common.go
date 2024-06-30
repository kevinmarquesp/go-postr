package data

type GenericErrorResponse struct {
	StatusText string `json:"statusText"`
	Error      string `json:"error"`
}
