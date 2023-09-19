package test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alvinmdj/mygram-api/routers"
	"github.com/stretchr/testify/assert"
)

func TestRegisterSuccess(t *testing.T) {
	router := routers.StartApp()

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/healthcheck", nil)

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "ok", responseBody["status"])
}
