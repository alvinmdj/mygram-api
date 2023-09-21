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
	TruncateSocialMediasTable(db)
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
	TruncateSocialMediasTable(db)
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
	TruncateSocialMediasTable(db)
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

func TestGetAllSocialMediaSuccess(t *testing.T) {
	// Setup
	db := SetupTestDB()
	TruncateUsersTable(db)
	TruncateSocialMediasTable(db)
	gin.SetMode(gin.ReleaseMode)
	router := routers.StartApp(db)

	// Create user
	RegisterTestUser(db)
	userData, _ := GetTestUser(db)
	jwt := helpers.GenerateToken(userData.ID, userData.Email)

	// Create social media
	CreateTestSocialMedia(db)

	// Send request
	request, _ := http.NewRequest(http.MethodGet, "/api/v1/social-medias", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %v", jwt))

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	// Read the response body & unmarshal the response into SocialMediaGetOutput struct
	body, _ := io.ReadAll(response.Body)
	var responseBody []models.SocialMediaGetOutput
	json.Unmarshal(body, &responseBody)

	assert.Len(t, responseBody, 1)
	assert.Equal(t, "My Youtube", responseBody[0].Name)
	assert.Equal(t, "https://www.youtube.com", responseBody[0].SocialMediaURL)
	assert.Equal(t, userData.ID, responseBody[0].User.ID)
	assert.Equal(t, userData.CreatedAt, responseBody[0].User.CreatedAt)
	assert.Equal(t, userData.UpdatedAt, responseBody[0].User.UpdatedAt)
	assert.Equal(t, userData.Username, responseBody[0].User.Username)
	assert.Equal(t, userData.Email, responseBody[0].User.Email)
	assert.Equal(t, userData.Age, responseBody[0].User.Age)
}

func TestGetAllSocialMediaFailedUnauthenticated(t *testing.T) {
	// Setup
	db := SetupTestDB()
	TruncateUsersTable(db)
	TruncateSocialMediasTable(db)
	gin.SetMode(gin.ReleaseMode)
	router := routers.StartApp(db)

	// Send request
	request, _ := http.NewRequest(http.MethodGet, "/api/v1/social-medias", nil)
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

func TestGetSocialMediaByIDSuccess(t *testing.T) {
	// Setup
	db := SetupTestDB()
	TruncateUsersTable(db)
	TruncateSocialMediasTable(db)
	gin.SetMode(gin.ReleaseMode)
	router := routers.StartApp(db)

	// Create user
	RegisterTestUser(db)
	userData, _ := GetTestUser(db)
	jwt := helpers.GenerateToken(userData.ID, userData.Email)

	// Create social media
	socialMedia := CreateTestSocialMedia(db)

	// Send request
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/social-medias/%d", socialMedia.ID), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %v", jwt))

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	// Read the response body & unmarshal the response into SocialMediaGetOutput struct
	body, _ := io.ReadAll(response.Body)
	var responseBody models.SocialMediaGetOutput
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "My Youtube", responseBody.Name)
	assert.Equal(t, "https://www.youtube.com", responseBody.SocialMediaURL)
	assert.Equal(t, userData.ID, responseBody.User.ID)
	assert.Equal(t, userData.CreatedAt, responseBody.User.CreatedAt)
	assert.Equal(t, userData.UpdatedAt, responseBody.User.UpdatedAt)
	assert.Equal(t, userData.Username, responseBody.User.Username)
	assert.Equal(t, userData.Email, responseBody.User.Email)
	assert.Equal(t, userData.Age, responseBody.User.Age)
}

func TestGetSocialMediaByIDFailedNotFound(t *testing.T) {
	// Setup
	db := SetupTestDB()
	TruncateUsersTable(db)
	TruncateSocialMediasTable(db)
	gin.SetMode(gin.ReleaseMode)
	router := routers.StartApp(db)

	// Create user
	userData := RegisterTestUser(db)
	jwt := helpers.GenerateToken(userData.ID, userData.Email)

	// Send request
	request, _ := http.NewRequest(http.MethodGet, "/api/v1/social-medias/123", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %v", jwt))

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	// Read the response body & unmarshal the response into ErrorResponse struct
	body, _ := io.ReadAll(response.Body)
	var responseBody models.ErrorResponse
	json.Unmarshal(body, &responseBody)

	assert.NotNil(t, responseBody.Error)
	assert.Equal(t, "NOT FOUND", responseBody.Error)
	assert.NotNil(t, responseBody.Message)
}

func TestGetSocialMediaByIDFailedUnauthenticated(t *testing.T) {
	// Setup
	db := SetupTestDB()
	TruncateUsersTable(db)
	TruncateSocialMediasTable(db)
	gin.SetMode(gin.ReleaseMode)
	router := routers.StartApp(db)

	// Create user
	RegisterTestUser(db)

	// Create social media
	socialMedia := CreateTestSocialMedia(db)

	// Send request
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/social-medias/%d", socialMedia.ID), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer hehe-token")

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

// TODO:
// UPDATE
// DELETE
