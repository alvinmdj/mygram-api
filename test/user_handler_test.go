package test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/alvinmdj/mygram-api/internal/models"
	"github.com/alvinmdj/mygram-api/internal/routers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRegisterSuccess(t *testing.T) {
	// Setup
	db := SetupTestDB()
	TruncateUsersTable(db)
	gin.SetMode(gin.ReleaseMode)
	router := routers.StartApp(db)

	// Create an instance of UserRegisterInput struct
	input := models.UserRegisterInput{
		Username: "alvinmdj",
		Email:    "alvinmdj@mygram.com",
		Password: "password",
		Age:      21,
	}

	// Marshal the struct into a JSON string
	jsonBody, _ := json.Marshal(input)

	// Create an io.Reader from the JSON string
	requestBody := strings.NewReader(string(jsonBody))
	request, _ := http.NewRequest(http.MethodPost, "/api/v1/users/register", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 201, response.StatusCode)

	// Read the response body & unmarshal the response into UserRegisterOutput struct
	body, _ := io.ReadAll(response.Body)
	var responseBody models.UserRegisterOutput
	json.Unmarshal(body, &responseBody)

	assert.NotNil(t, responseBody.ID)
	assert.NotNil(t, responseBody.CreatedAt)
	assert.NotNil(t, responseBody.UpdatedAt)
	assert.Equal(t, "alvinmdj", responseBody.Username)
	assert.Equal(t, "alvinmdj@mygram.com", responseBody.Email)
	assert.Equal(t, 21, responseBody.Age)
}

func TestRegisterFailed(t *testing.T) {
	// Setup
	db := SetupTestDB()
	TruncateUsersTable(db)
	gin.SetMode(gin.ReleaseMode)
	router := routers.StartApp(db)

	// Empty input
	input := models.UserRegisterInput{}

	// Marshal the struct into a JSON string
	jsonBody, _ := json.Marshal(input)

	// Create an io.Reader from the JSON string
	requestBody := strings.NewReader(string(jsonBody))
	request, _ := http.NewRequest(http.MethodPost, "/api/v1/users/register", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	// Read the response body & unmarshal the response into ErrorResponse struct
	body, _ := io.ReadAll(response.Body)
	var responseBody models.ErrorResponse
	json.Unmarshal(body, &responseBody)

	assert.NotNil(t, responseBody.Error)
	assert.Equal(t, "BAD REQUEST", responseBody.Error)
	assert.NotNil(t, responseBody.Message)
}

func TestLoginSuccess(t *testing.T) {
	// Setup
	db := SetupTestDB()
	TruncateUsersTable(db)
	gin.SetMode(gin.ReleaseMode)
	router := routers.StartApp(db)

	// Register a new test user
	RegisterTestUser(db)

	// Create an instance of UserLoginInput struct
	input := models.UserLoginInput{
		Email:    "alvinmdj@mygram.com",
		Password: "password",
	}

	// Marshal the struct into a JSON string
	jsonBody, _ := json.Marshal(input)

	// Create an io.Reader from the JSON string
	requestBody := strings.NewReader(string(jsonBody))
	request, _ := http.NewRequest(http.MethodPost, "/api/v1/users/login", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 201, response.StatusCode)

	// Read the response body & unmarshal the response into UserLoginOutput struct
	body, _ := io.ReadAll(response.Body)
	var responseBody models.UserLoginOutput
	json.Unmarshal(body, &responseBody)

	assert.NotNil(t, responseBody.Token)
}

func TestLoginFailed(t *testing.T) {
	// Setup
	db := SetupTestDB()
	TruncateUsersTable(db)
	gin.SetMode(gin.ReleaseMode)
	router := routers.StartApp(db)

	// Register a new test user
	RegisterTestUser(db)

	// Create an instance of UserLoginInput struct with wrong credentials
	input := models.UserLoginInput{
		Email:    "wrong@test.com",
		Password: "hahahaha",
	}

	// Marshal the struct into a JSON string
	jsonBody, _ := json.Marshal(input)

	// Create an io.Reader from the JSON string
	requestBody := strings.NewReader(string(jsonBody))
	request, _ := http.NewRequest(http.MethodPost, "/api/v1/users/login", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 401, response.StatusCode)

	// Read the response body & unmarshal the response into ErrorResponse struct
	body, _ := io.ReadAll(response.Body)
	var responseBody models.ErrorResponse
	json.Unmarshal(body, &responseBody)

	assert.NotNil(t, responseBody.Error)
	assert.Equal(t, "UNAUTHORIZED", responseBody.Error)
	assert.NotNil(t, responseBody.Message)
}
