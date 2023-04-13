package handlers

import (
	"net/http"
	"strconv"

	"github.com/alvinmdj/mygram-api/helpers"
	"github.com/alvinmdj/mygram-api/models"
	"github.com/alvinmdj/mygram-api/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type SocialMediaHdlInterface interface {
	GetAll(c *gin.Context)
	GetOneById(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type SocialMediaHandler struct {
	socialMediaSvc services.SocialMediaSvcInterface
}

func NewSocialMediaHdl(socialMediaSvc services.SocialMediaSvcInterface) SocialMediaHdlInterface {
	return &SocialMediaHandler{
		socialMediaSvc: socialMediaSvc,
	}
}

func (s *SocialMediaHandler) GetAll(c *gin.Context) {
	socialMedias, err := s.socialMediaSvc.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
		return
	}

	socialMediasResponse := []models.SocialMediaGetOutput{}
	for _, socialMedia := range socialMedias {
		socialMediaOutput := models.SocialMediaGetOutput{
			Base:           socialMedia.Base,
			Name:           socialMedia.Name,
			SocialMediaURL: socialMedia.SocialMediaURL,
			User: models.UserRegisterOutput{
				Base:     socialMedia.User.Base,
				Username: socialMedia.User.Username,
				Email:    socialMedia.User.Email,
				Age:      socialMedia.User.Age,
			},
		}
		socialMediasResponse = append(socialMediasResponse, socialMediaOutput)
	}
	c.JSON(http.StatusOK, socialMediasResponse)
}

func (s *SocialMediaHandler) GetOneById(c *gin.Context) {
	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))

	socialMedia, err := s.socialMediaSvc.GetOneById(socialMediaId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "not found",
			"message": err.Error(),
		})
		return
	}

	socialMediaResponse := models.SocialMediaGetOutput{
		Base:           socialMedia.Base,
		Name:           socialMedia.Name,
		SocialMediaURL: socialMedia.SocialMediaURL,
		User: models.UserRegisterOutput{
			Base:     socialMedia.User.Base,
			Username: socialMedia.User.Username,
			Email:    socialMedia.User.Email,
			Age:      socialMedia.User.Age,
		},
	}
	c.JSON(http.StatusOK, socialMediaResponse)
}

func (s *SocialMediaHandler) Create(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	socialMediaInput := models.SocialMediaCreateInput{}

	// get token claims in userData context from authentication middleware
	// and cast the data type from any to jwt.MapClaims
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	socialMediaInput.UserID = userId

	if contentType == helpers.AppJson {
		c.ShouldBindJSON(&socialMediaInput)
	} else {
		c.ShouldBind(&socialMediaInput)
	}

	socialMedia, err := s.socialMediaSvc.Create(socialMediaInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "BAD REQUEST",
			"message": err.Error(),
		})
		return
	}

	socialMediaResponse := models.SocialMediaCreateOutput{
		Base:           socialMedia.Base,
		Name:           socialMedia.Name,
		SocialMediaURL: socialMedia.SocialMediaURL,
		UserID:         socialMedia.UserID,
	}
	c.JSON(http.StatusCreated, socialMediaResponse)
}

func (s *SocialMediaHandler) Update(c *gin.Context) {}

func (s *SocialMediaHandler) Delete(c *gin.Context) {}
