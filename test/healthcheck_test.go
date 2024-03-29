package test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alvinmdj/mygram-api/internal/routers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealthcheckOK(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
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
