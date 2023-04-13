package services

import (
	"github.com/alvinmdj/mygram-api/models"
	"github.com/alvinmdj/mygram-api/repositories"
)

type SocialMediaSvcInterface interface {
	GetAll() (socialMedias []models.SocialMedia, err error)
	GetOneById(id int) (socialMedia models.SocialMedia, err error)
	Create(socialMediaInput models.SocialMediaCreateInput) (socialMedia models.SocialMedia, err error)
	// Update()
	// Delete()
}

type SocialMediaSvc struct {
	socialMediaRepo repositories.SocialMediaRepoInterface
}

func NewSocialMediaSvc(socialMediaRepo repositories.SocialMediaRepoInterface) SocialMediaSvcInterface {
	return &SocialMediaSvc{
		socialMediaRepo: socialMediaRepo,
	}
}

func (s *SocialMediaSvc) GetAll() (socialMedias []models.SocialMedia, err error) {
	socialMedias, err = s.socialMediaRepo.FindAll()
	return
}

func (s *SocialMediaSvc) GetOneById(id int) (socialMedia models.SocialMedia, err error) {
	socialMedia, err = s.socialMediaRepo.FindById(id)
	return
}

func (s *SocialMediaSvc) Create(socialMediaInput models.SocialMediaCreateInput) (socialMedia models.SocialMedia, err error) {
	socialMedia = models.SocialMedia{
		Name:           socialMediaInput.Name,
		SocialMediaURL: socialMediaInput.SocialMediaURL,
		UserID:         socialMediaInput.UserID,
	}

	socialMedia, err = s.socialMediaRepo.Save(socialMedia)
	return
}
