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

// Social Media GetAll godoc
// @Summary Get all social media
// @Description Get all social media
// @Tags socialMedias
// @Param Authorization header string true "format: Bearer token-here"
// @Produce json
// @Success 200 {object} []models.SocialMediaGetOutput{}
// @Failure 400 {object} map[string]string{}
// @Failure 500 {object} map[string]string{}
// @Router /api/v1/social-medias [get]
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

// Social Media GetOneById godoc
// @Summary Get one social media by id
// @Description Get one social media by id
// @Tags socialMedias
// @Param socialMediaId path string true "get social media by id"
// @Param Authorization header string true "format: Bearer token-here"
// @Produce json
// @Success 200 {object} models.SocialMediaGetOutput{}
// @Failure 404 {object} map[string]string{}
// @Failure 500 {object} map[string]string{}
// @Router /api/v1/social-medias/:socialMediaId [get]
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

// Social Media Create godoc
// @Summary Create social media
// @Description Create social media
// @Tags socialMedias
// @Accept json,mpfd
// @Produce json
// @Param models.SocialMediaCreateInput body models.SocialMediaCreateInput{} true "create social media"
// @Param Authorization header string true "format: Bearer token-here"
// @Success 201 {object} models.SocialMediaCreateOutput{}
// @Failure 400 {object} map[string]string{}
// @Failure 500 {object} map[string]string{}
// @Router /api/v1/social-medias [post]
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
