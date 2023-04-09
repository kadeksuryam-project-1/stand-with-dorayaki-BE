package repository

import (
	"backend/internal/schema"

	"gorm.io/gorm"
)

type IDorayakiRepository interface {
	CreateOne(dorayaki schema.Dorayaki) (schema.Dorayaki, error)
	UpdateOne(dorayaki schema.Dorayaki, id int) (schema.Dorayaki, error)
	GetAll() ([]schema.Dorayaki, error)
	GetOne(id int) (schema.Dorayaki, error)
	DeleteOne(id int) error
}

type dorayakiRepository struct {
	db *gorm.DB
}

// Contructor
func NewDorayakiRepository(db *gorm.DB) IDorayakiRepository {
	return &dorayakiRepository{
		db: db,
	}
}

// CREATE A DORAYAKI
func (r *dorayakiRepository) CreateOne(dorayaki schema.Dorayaki) (schema.Dorayaki, error) {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := createDorayaki(tx, &dorayaki); err != nil {
		tx.Rollback()
		return schema.Dorayaki{}, err
	}

	stores, err := findAllStores(tx)
	if err != nil {
		tx.Rollback()
		return schema.Dorayaki{}, err
	}

	newDorayakiStoreStocks := createDorayakiStoreStocksFromADorayaki(dorayaki, stores)

	if err := createDorayakiStoreStocksInBatches(tx, newDorayakiStoreStocks); err != nil {
		tx.Rollback()
		return schema.Dorayaki{}, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return schema.Dorayaki{}, err
	}

	return dorayaki, nil
}

func createDorayaki(tx *gorm.DB, dorayaki *schema.Dorayaki) error {
	return tx.Table("dorayakis").Create(dorayaki).Error
}

func findAllStores(tx *gorm.DB) ([]schema.Store, error) {
	var stores []schema.Store
	if err := tx.Table("stores").Find(&stores).Error; err != nil {
		return nil, err
	}
	return stores, nil
}

func createDorayakiStoreStocksFromADorayaki(dorayaki schema.Dorayaki, stores []schema.Store) []schema.DorayakiStoreStock {
	var newDorayakiStoreStocks []schema.DorayakiStoreStock
	for _, store := range stores {
		newdorayakiStoreStock := schema.DorayakiStoreStock{
			DorayakiId: dorayaki.ID,
			StoreId:    store.ID,
			Stock:      0,
		}
		newDorayakiStoreStocks = append(newDorayakiStoreStocks, newdorayakiStoreStock)
	}
	return newDorayakiStoreStocks
}

func createDorayakiStoreStocksInBatches(tx *gorm.DB, dorayakiStoreStocks []schema.DorayakiStoreStock) error {
	return tx.Table("dorayaki_store_stocks").CreateInBatches(&dorayakiStoreStocks, 10).Error
}

// UPDATE A DORAYAKI
func (r *dorayakiRepository) UpdateOne(dorayaki schema.Dorayaki, id int) (schema.Dorayaki, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return schema.Dorayaki{}, tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var existingDorayaki schema.Dorayaki
	if err := tx.First(&existingDorayaki, id).Error; err != nil {
		tx.Rollback()
		return schema.Dorayaki{}, err
	}

	updates := map[string]interface{}{
		"Flavor":      dorayaki.Flavor,
		"Description": dorayaki.Description,
		"Image":       dorayaki.Image,
	}

	if err := tx.Model(&existingDorayaki).Updates(updates).Error; err != nil {
		tx.Rollback()
		return schema.Dorayaki{}, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return schema.Dorayaki{}, err
	}

	return existingDorayaki, nil
}

// Get All Dorayakis
func (r *dorayakiRepository) GetAll() ([]schema.Dorayaki, error) {
	var dorayakis []schema.Dorayaki

	err := r.db.Session(&gorm.Session{}).Find(&dorayakis).Error

	return dorayakis, err
}

// Get One Dorayaki
func (r *dorayakiRepository) GetOne(id int) (schema.Dorayaki, error) {
	var dorayaki schema.Dorayaki

	err := r.db.Session(&gorm.Session{}).First(&dorayaki, id).Error

	return dorayaki, err
}

// Delete One Dorayaki
func (r *dorayakiRepository) DeleteOne(id int) error {
	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := deleteAllCorrespondingStocksByDorayaki(tx, id); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where("id = ?", id).Delete(&schema.Dorayaki{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func deleteAllCorrespondingStocksByDorayaki(tx *gorm.DB, dorayakiID int) error {
	var dorayaki schema.Dorayaki
	if err := tx.First(&dorayaki, dorayakiID).Error; err != nil {
		return err
	}

	if err := tx.Where("dorayaki_id = ?", dorayakiID).Delete(&schema.DorayakiStoreStock{}).Error; err != nil {
		return err
	}

	return nil
}
