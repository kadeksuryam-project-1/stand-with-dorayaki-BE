package service

import (
	"backend/config"
	"backend/internal/repository"
	"backend/internal/schema"
	"backend/pkg/lib"
	"context"
	"fmt"
	"math/rand"
	"path/filepath"
	"strings"
	"time"
)

type IDorayakiService interface {
	CreateDorayaki(form schema.DorayakiForm) (schema.Dorayaki, error)
	GetDorayakis() ([]schema.Dorayaki, error)
	GetDorayaki(id int) (schema.Dorayaki, error)
	UpdateDorayaki(form schema.DorayakiForm, id int) (schema.Dorayaki, error)
	DeleteDorayaki(id int) error
}

type dorayakiService struct {
	dorayakiRepository repository.IDorayakiRepository
}

func NewDorayakiService(dorayakiRepository repository.IDorayakiRepository) IDorayakiService {
	return &dorayakiService{
		dorayakiRepository: dorayakiRepository,
	}
}

func (s *dorayakiService) CreateDorayaki(form schema.DorayakiForm) (schema.Dorayaki, error) {
	emptyDorayaki := schema.Dorayaki{}
	bucketAddress := config.C.BucketAddress
	newDorayaki := schema.Dorayaki{
		Flavor:      form.Flavor,
		Description: form.Description,
		Image:       bucketAddress + "assets/dorayaki/default.png",
	}

	if form.Image != nil {
		src, err := form.Image.Open()
		if err != nil {
			return emptyDorayaki, err
		}
		defer src.Close()

		imgPath := "assets/dorayaki/"
		imgName := imgPath + generateFilename(form.Image.Filename)

		if err := lib.GStorageUploader.UploadFile(context.Background(), imgName, src); err != nil {
			return emptyDorayaki, err
		}

		newDorayaki.Image = bucketAddress + imgName
	}

	savedDorayaki, err := s.dorayakiRepository.CreateOne(newDorayaki)

	if err != nil {
		if form.Image != nil {
			lib.GStorageUploader.DeleteFile(context.Background(), strings.TrimPrefix(newDorayaki.Image, config.C.BucketAddress))
		}
		return emptyDorayaki, err
	}

	return savedDorayaki, nil
}

func generateFilename(originalFilename string) string {
	ext := filepath.Ext(originalFilename)
	name := originalFilename[:len(originalFilename)-len(ext)]
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	randomString := randString(6) // generate a random 6-character string
	return fmt.Sprintf("%s_%s_%s%s", name, timestamp, randomString, ext)
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func (s *dorayakiService) GetDorayakis() ([]schema.Dorayaki, error) {
	dorayakis, err := s.dorayakiRepository.GetAll()

	return dorayakis, err
}

func (s *dorayakiService) GetDorayaki(id int) (schema.Dorayaki, error) {
	dorayaki, err := s.dorayakiRepository.GetOne(id)

	return dorayaki, err
}

func (s *dorayakiService) UpdateDorayaki(form schema.DorayakiForm, id int) (schema.Dorayaki, error) {
	emptyDorayaki := schema.Dorayaki{}
	bucketAddress := config.C.BucketAddress
	dorayaki, err := s.dorayakiRepository.GetOne(id)
	if err != nil {
		return emptyDorayaki, err
	}

	updateDorayaki := schema.Dorayaki{
		Flavor:      form.Flavor,
		Description: form.Description,
		Image:       dorayaki.Image,
	}

	if form.Image != nil {
		src, err := form.Image.Open()
		if err != nil {
			return emptyDorayaki, err
		}
		defer src.Close()

		imgPath := "assets/dorayaki/"
		imgName := imgPath + generateFilename(form.Image.Filename)

		if err := lib.GStorageUploader.UploadFile(context.Background(), imgName, src); err != nil {
			return emptyDorayaki, err
		}

		updateDorayaki.Image = bucketAddress + imgName
	}

	updatedDorayaki, err := s.dorayakiRepository.UpdateOne(updateDorayaki, id)
	if err != nil {
		if form.Image != nil {
			lib.GStorageUploader.DeleteFile(context.Background(), strings.TrimPrefix(updateDorayaki.Image, config.C.BucketAddress))
		}
		return emptyDorayaki, err
	}

	defaultImage := bucketAddress + "assets/dorayaki/default.png"
	if form.Image != nil && dorayaki.Image != defaultImage {
		lib.GStorageUploader.DeleteFile(context.Background(), strings.TrimPrefix(dorayaki.Image, config.C.BucketAddress))
	}

	return updatedDorayaki, nil
}

func (s *dorayakiService) DeleteDorayaki(id int) error {
	dorayaki, err := s.dorayakiRepository.GetOne(id)
	if err != nil {
		return err
	}
	bucketAddress := config.C.BucketAddress
	defaultImage := bucketAddress + "assets/dorayaki/default.png"

	if dorayaki.Image != defaultImage {
		if err := lib.GStorageUploader.DeleteFile(context.Background(), strings.TrimPrefix(dorayaki.Image, config.C.BucketAddress)); err != nil {
			return err
		}
	}
	return s.dorayakiRepository.DeleteOne(id)
}
