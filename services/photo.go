package services

import (
	"log"
	"mime/multipart"

	"github.com/alvinmdj/mygram-api/helpers"
	"github.com/alvinmdj/mygram-api/models"
	"github.com/alvinmdj/mygram-api/repositories"
)

type PhotoSvcInterface interface {
	GetAll() (photos []models.Photo, err error)
	// GetOneById(id int) (socialMedia models.SocialMedia, err error)
	Create(photoInput models.PhotoCreateInput, photoFileHeader *multipart.FileHeader) (photo models.Photo, err error)
	// Update(socialMediaInput models.SocialMediaUpdateInput) (socialMedia models.SocialMedia, err error)
	// Delete(id int) (err error)
}

type PhotoSvc struct {
	photoRepo repositories.PhotoRepoInterface
}

func NewPhotoSvc(photoRepo repositories.PhotoRepoInterface) PhotoSvcInterface {
	return &PhotoSvc{
		photoRepo: photoRepo,
	}
}

func (p *PhotoSvc) GetAll() (photos []models.Photo, err error) {
	photos, err = p.photoRepo.FindAll()
	return
}

func (p *PhotoSvc) Create(photoInput models.PhotoCreateInput, photoFileHeader *multipart.FileHeader) (photo models.Photo, err error) {
	// open the file and get its content
	photoFile, err := photoFileHeader.Open()
	if err != nil {
		log.Printf("error opening file: %v", err)
	}
	defer photoFile.Close()

	photoUrl, err := helpers.UploadToCloudinary(photoFile)
	if err != nil {
		return
	}

	photo = models.Photo{
		Title:    photoInput.Title,
		Caption:  photoInput.Caption,
		UserID:   photoInput.UserID,
		PhotoURL: photoUrl,
	}

	photo, err = p.photoRepo.Save(photo)
	return
}
