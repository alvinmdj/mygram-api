package test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/alvinmdj/mygram-api/helpers"
	"github.com/alvinmdj/mygram-api/models"
	"github.com/alvinmdj/mygram-api/routers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateSocialMediaSuccess(t *testing.T) {
	// Setup
	db := SetupTestDB()
	TruncateUsersTable(db)
	gin.SetMode(gin.ReleaseMode)
	router := routers.StartApp(db)

	// Create user
	RegisterTestUser(db)
	userData, _ := GetTestUser(db)
	jwt := helpers.GenerateToken(userData.ID, userData.Email)

	// Create an instance of SocialMediaCreateInput struct
	input := models.SocialMediaCreateInput{
		UserID:         userData.ID,
		Name:           "My Youtube",
		SocialMediaURL: "https://www.youtube.com",
	}

	// Marshal the struct into a JSON string
	jsonBody, _ := json.Marshal(input)

	// Create an io.Reader from the JSON string
	requestBody := strings.NewReader(string(jsonBody))
	request, _ := http.NewRequest(http.MethodPost, "/api/v1/social-medias", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %v", jwt))

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 201, response.StatusCode)

	// Read the response body & unmarshal the response into SocialMediaCreateOutput struct
	body, _ := io.ReadAll(response.Body)
	var responseBody models.SocialMediaCreateOutput
	json.Unmarshal(body, &responseBody)

	assert.NotNil(t, responseBody.ID)
	assert.NotNil(t, responseBody.CreatedAt)
	assert.NotNil(t, responseBody.UpdatedAt)
	assert.Equal(t, "My Youtube", responseBody.Name)
	assert.Equal(t, "https://www.youtube.com", responseBody.SocialMediaURL)
	assert.Equal(t, userData.ID, responseBody.UserID)
}

func TestCreateSocialMediaFailedBadRequest(t *testing.T) {
	// Setup
	db := SetupTestDB()
	TruncateUsersTable(db)
	gin.SetMode(gin.ReleaseMode)
	router := routers.StartApp(db)

	// Create user
	RegisterTestUser(db)
	userData, _ := GetTestUser(db)
	jwt := helpers.GenerateToken(userData.ID, userData.Email)

	// Create an instance of SocialMediaCreateInput struct
	input := models.SocialMediaCreateInput{
		// no user id & empty field
		Name:           "",
		SocialMediaURL: "",
	}

	// Marshal the struct into a JSON string
	jsonBody, _ := json.Marshal(input)

	// Create an io.Reader from the JSON string
	requestBody := strings.NewReader(string(jsonBody))
	request, _ := http.NewRequest(http.MethodPost, "/api/v1/social-medias", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %v", jwt))

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

func TestCreateSocialMediaFailedUnauthenticated(t *testing.T) {
	// Setup
	db := SetupTestDB()
	TruncateUsersTable(db)
	gin.SetMode(gin.ReleaseMode)
	router := routers.StartApp(db)

	// Create an instance of SocialMediaCreateInput struct
	input := models.SocialMediaCreateInput{
		UserID:         1,
		Name:           "My Youtube",
		SocialMediaURL: "https://www.youtube.com",
	}

	// Marshal the struct into a JSON string
	jsonBody, _ := json.Marshal(input)

	// Create an io.Reader from the JSON string
	requestBody := strings.NewReader(string(jsonBody))
	request, _ := http.NewRequest(http.MethodPost, "/api/v1/social-medias", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer wrong-token")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 401, response.StatusCode)

	// Read the response body & unmarshal the response into ErrorResponse struct
	body, _ := io.ReadAll(response.Body)
	var responseBody models.ErrorResponse
	json.Unmarshal(body, &responseBody)

	assert.NotNil(t, responseBody.Error)
	assert.Equal(t, "UNAUTHENTICATED", responseBody.Error)
	assert.NotNil(t, responseBody.Message)
}
