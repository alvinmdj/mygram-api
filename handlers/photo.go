package handlers

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/alvinmdj/mygram-api/helpers"
	"github.com/alvinmdj/mygram-api/models"
	"github.com/alvinmdj/mygram-api/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type PhotoHdlInterface interface {
	GetAll(c *gin.Context)
	GetOneById(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type PhotoHandler struct {
	photoSvc services.PhotoSvcInterface
}

func NewPhotoHdl(photoSvc services.PhotoSvcInterface) PhotoHdlInterface {
	return &PhotoHandler{
		photoSvc: photoSvc,
	}
}

// Photo GetAll godoc
// @Summary Get all photos
// @Description Get all photos
// @Tags photos
// @Param Authorization header string true "format: Bearer token-here"
// @Produce json
// @Success 200 {object} []models.PhotoGetOutput{}
// @Failure 400 {object} map[string]string{}
// @Failure 500 {object} map[string]string{}
// @Router /api/v1/photos [get]
func (p *PhotoHandler) GetAll(c *gin.Context) {
	photos, err := p.photoSvc.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "BAD REQUEST",
			"message": err.Error(),
		})
		return
	}

	photosResponse := []models.PhotoGetOutput{}
	for _, photo := range photos {
		photoOutput := models.PhotoGetOutput{
			Base:     photo.Base,
			Title:    photo.Title,
			Caption:  photo.Caption,
			PhotoURL: photo.PhotoURL,
			User: models.UserRegisterOutput{
				Base:     photo.User.Base,
				Username: photo.User.Username,
				Email:    photo.User.Email,
				Age:      photo.User.Age,
			},
		}
		photosResponse = append(photosResponse, photoOutput)
	}
	c.JSON(http.StatusOK, photosResponse)
}

func (p *PhotoHandler) GetOneById(c *gin.Context) {

}

// Photo Create godoc
// @Summary Create photos
// @Description Create photos
// @Tags photos
// @Accept mpfd
// @Produce json
// @Param models.PhotoCreateInput body models.PhotoCreateInput{} true "create photo"
// @Param Authorization header string true "format: Bearer token-here"
// @Success 201 {object} models.PhotoCreateOutput{}
// @Failure 400 {object} map[string]string{}
// @Failure 413 {object} map[string]string{}
// @Failure 500 {object} map[string]string{}
// @Router /api/v1/photos [post]
func (p *PhotoHandler) Create(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	photoInput := models.PhotoCreateInput{}

	// get token claims in userData context from authentication middleware
	// and cast the data type from any to jwt.MapClaims
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	photoInput.UserID = userId

	// only accept multipart/form-data
	if contentType == helpers.AppJson {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "BAD REQUEST",
			"message": "invalid content type",
		})
		return
	} else {
		c.ShouldBind(&photoInput)
	}

	// photo source, check if photo is uploaded
	photoFileHeader, err := c.FormFile("photo")
	if err != nil {
		log.Printf("get form err - %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "BAD REQUEST",
			"message": "no photo file uploaded",
		})
		return
	}

	// Check if the file is an image
	ext := filepath.Ext(photoFileHeader.Filename)
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".webp" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "BAD REQUEST",
			"message": "invalid file type",
		})
		return
	}

	photo, err := p.photoSvc.Create(photoInput, photoFileHeader)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "BAD REQUEST",
			"message": err.Error(),
		})
		return
	}

	photoResponse := models.PhotoCreateOutput{
		Base:     photo.Base,
		Title:    photo.Title,
		Caption:  photo.Caption,
		PhotoURL: photo.PhotoURL,
		UserID:   photo.UserID,
	}
	c.JSON(http.StatusCreated, photoResponse)
}

func (p *PhotoHandler) Update(c *gin.Context) {

}

func (p *PhotoHandler) Delete(c *gin.Context) {

}
