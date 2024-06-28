package utils_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kevinmarquesp/go-postr/internal/data"
	"github.com/kevinmarquesp/go-postr/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestWriteJsonError(t *testing.T) {
	w := httptest.NewRecorder()
	err := fmt.Errorf("Something went wrong")
	status := http.StatusInternalServerError

	utils.WriteJsonError(w, status, err)

	if w.Code != status {
		t.Errorf("Expected status code %d, but got %d", status, w.Code)
	}

	var received data.ErrorResponse

	json.Unmarshal([]byte(w.Body.String()), &received)

	expected := data.ErrorResponse{
		Status:     500,
		StatusText: http.StatusText(http.StatusInternalServerError),
		Error:      err.Error(),
	}

	assert.Equal(t, received.Status, expected.Status)
	assert.Equal(t, received.StatusText, expected.StatusText)
	assert.Equal(t, received.Error, expected.Error)
}
