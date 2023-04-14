package helpers

import (
	"context"
	"log"
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
)

var (
	cld *cloudinary.Cloudinary
	err error
)

func InitCloudinary() {
	cloudName := os.Getenv("CLOUDINARY_CLOUD_NAME")
	apiKey := os.Getenv("CLOUDINARY_API_KEY")
	apiSecret := os.Getenv("CLOUDINARY_API_SECRET")

	cld, err = cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	if err != nil {
		log.Fatal("error connecting to cloudinary:", err.Error())
	}

	log.Println("cloudinary connected successfully")
}

func UploadToCloudinary(file multipart.File) (string, error) {
	ctx := context.Background()
	publicId := uuid.New()

	resp, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
		PublicID: "photos/" + publicId.String(), // folder-name/file-name
	})
	if err != nil {
		log.Printf("error uploading file to cloudinary: %v", err.Error())
	}

	return resp.SecureURL, err
}
