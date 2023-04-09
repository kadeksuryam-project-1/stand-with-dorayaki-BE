package service

import (
	"backend/config"
	"backend/internal/repository"
	"backend/internal/schema"
	"backend/pkg/lib"
	"context"
	"strings"
)

type IStoreService interface {
	CreateStore(form schema.StoreForm) (schema.Store, error)
	GetStores() ([]schema.Store, error)
	GetStore(id int) (schema.Store, error)
	UpdateStore(form schema.StoreForm, id int) (schema.Store, error)
	DeleteStore(id int) error
}

type storeService struct {
	storeRepository repository.IStoreRepository
}

func NewStoreService(storeRepository repository.IStoreRepository) IStoreService {
	return &storeService{
		storeRepository: storeRepository,
	}
}

func (s *storeService) CreateStore(form schema.StoreForm) (schema.Store, error) {
	emptyStore := schema.Store{}
	bucketAddress := config.C.BucketAddress
	newStore := schema.Store{
		Name:        form.Name,
		Street:      form.Street,
		Subdistrict: form.Subdistrict,
		District:    form.District,
		Province:    form.Province,
		Image:       bucketAddress + "assets/store/default.png",
	}

	if form.Image != nil {
		src, err := form.Image.Open()
		if err != nil {
			return emptyStore, err
		}
		defer src.Close()

		imgPath := "assets/store/"
		imgName := imgPath + generateFilename(form.Image.Filename)

		if err := lib.GStorageUploader.UploadFile(context.Background(), imgName, src); err != nil {
			return emptyStore, err
		}

		newStore.Image = bucketAddress + imgName
	}

	savedStore, err := s.storeRepository.CreateOne(newStore)

	if err != nil {
		if form.Image != nil {
			lib.GStorageUploader.DeleteFile(context.Background(), strings.TrimPrefix(newStore.Image, config.C.BucketAddress))
		}
		return emptyStore, err
	}

	return savedStore, nil
}

func (s *storeService) GetStores() ([]schema.Store, error) {
	stores, err := s.storeRepository.GetAll()

	return stores, err
}

func (s *storeService) GetStore(id int) (schema.Store, error) {
	store, err := s.storeRepository.GetOne(id)

	return store, err
}

func (s *storeService) UpdateStore(form schema.StoreForm, id int) (schema.Store, error) {
	emptyStore := schema.Store{}
	bucketAddress := config.C.BucketAddress
	store, err := s.storeRepository.GetOne(id)
	if err != nil {
		return emptyStore, err
	}

	updateStore := schema.Store{
		Name:        form.Name,
		Street:      form.Street,
		Subdistrict: form.Subdistrict,
		District:    form.District,
		Province:    form.Province,
		Image:       store.Image,
	}

	if form.Image != nil {
		src, err := form.Image.Open()
		if err != nil {
			return emptyStore, err
		}
		defer src.Close()

		imgPath := "assets/store/"
		imgName := imgPath + generateFilename(form.Image.Filename)

		if err := lib.GStorageUploader.UploadFile(context.Background(), imgName, src); err != nil {
			return emptyStore, err
		}

		updateStore.Image = bucketAddress + imgName
	}

	updatedStore, err := s.storeRepository.UpdateOne(updateStore, id)

	if err != nil {
		if form.Image != nil {
			lib.GStorageUploader.DeleteFile(context.Background(), strings.TrimPrefix(updatedStore.Image, config.C.BucketAddress))
		}
		return emptyStore, err
	}

	defaultImage := bucketAddress + "assets/store/default.png"
	if form.Image != nil && store.Image != defaultImage {
		lib.GStorageUploader.DeleteFile(context.Background(), strings.TrimPrefix(store.Image, config.C.BucketAddress))
	}

	return updatedStore, nil
}

func (s *storeService) DeleteStore(id int) error {
	store, err := s.storeRepository.GetOne(id)
	if err != nil {
		return err
	}
	bucketAddress := config.C.BucketAddress
	defaultImage := bucketAddress + "assets/store/default.png"

	if store.Image != defaultImage {
		if err := lib.GStorageUploader.DeleteFile(context.Background(), strings.TrimPrefix(store.Image, config.C.BucketAddress)); err != nil {
			return err
		}
	}

	return s.storeRepository.DeleteOne(id)
}
