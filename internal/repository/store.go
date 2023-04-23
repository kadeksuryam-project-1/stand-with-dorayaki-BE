package repository

import (
	"backend/internal/schema"

	"gorm.io/gorm"
)

type IStoreRepository interface {
	CreateOne(store schema.Store) (schema.Store, error)
	UpdateOne(store schema.Store, id int) (schema.Store, error)
	GetAll() ([]schema.Store, error)
	GetOne(id int) (schema.Store, error)
	DeleteOne(id int) error
}

type storeRepository struct {
	db *gorm.DB
}

// Contructor
func NewStoreRepository(db *gorm.DB) IStoreRepository {
	return &storeRepository{
		db: db,
	}
}

// CREATE A STORE
func (r *storeRepository) CreateOne(store schema.Store) (schema.Store, error) {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := createStore(tx, &store); err != nil {
		tx.Rollback()
		return schema.Store{}, err
	}

	dorayakis, err := findAllDorayakis(tx)
	if err != nil {
		tx.Rollback()
		return schema.Store{}, err
	}

	newDorayakiStoreStocks := createDorayakiStoreStocksFromAStore(store, dorayakis)

	if err := createDorayakiStoreStocksInBatches(tx, newDorayakiStoreStocks); err != nil {
		tx.Rollback()
		return schema.Store{}, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return schema.Store{}, err
	}

	return store, nil
}

func createStore(tx *gorm.DB, store *schema.Store) error {
	return tx.Table("stores").Create(store).Error
}

func findAllDorayakis(tx *gorm.DB) ([]schema.Dorayaki, error) {
	var dorayakis []schema.Dorayaki
	if err := tx.Table("dorayakis").Order("id ASC").Find(&dorayakis).Error; err != nil {
		return nil, err
	}
	return dorayakis, nil
}

func createDorayakiStoreStocksFromAStore(store schema.Store, dorayakis []schema.Dorayaki) []schema.DorayakiStoreStock {
	var newDorayakiStoreStocks []schema.DorayakiStoreStock
	for _, dorayaki := range dorayakis {
		newdorayakiStoreStock := schema.DorayakiStoreStock{
			DorayakiId: dorayaki.ID,
			StoreId:    store.ID,
			Stock:      0,
		}
		newDorayakiStoreStocks = append(newDorayakiStoreStocks, newdorayakiStoreStock)
	}
	return newDorayakiStoreStocks
}

// UPDATE A STORE
func (r *storeRepository) UpdateOne(store schema.Store, id int) (schema.Store, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return schema.Store{}, tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var existingStore schema.Store
	if err := tx.First(&existingStore, id).Error; err != nil {
		tx.Rollback()
		return schema.Store{}, err
	}

	updates := map[string]interface{}{
		"Name":        store.Name,
		"Street":      store.Street,
		"Subdistrict": store.Subdistrict,
		"District":    store.District,
		"Province":    store.Province,
		"Image":       store.Image,
	}

	if err := tx.Model(&existingStore).Updates(updates).Error; err != nil {
		tx.Rollback()
		return schema.Store{}, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return schema.Store{}, err
	}

	return existingStore, nil
}

// Get All Stores
func (r *storeRepository) GetAll() ([]schema.Store, error) {
	var stores []schema.Store

	err := r.db.Session(&gorm.Session{}).Order("id ASC").Find(&stores).Error

	return stores, err
}

// Get One Store
func (r *storeRepository) GetOne(id int) (schema.Store, error) {
	var store schema.Store

	err := r.db.Session(&gorm.Session{}).First(&store, id).Error

	return store, err
}

// Delete One Store
func (r *storeRepository) DeleteOne(id int) error {
	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := deleteAllCorrespondingStocksByStore(tx, id); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where("id = ?", id).Delete(&schema.Store{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func deleteAllCorrespondingStocksByStore(tx *gorm.DB, storeID int) error {
	var store schema.Store
	if err := tx.First(&store, storeID).Error; err != nil {
		return err
	}

	if err := tx.Where("store_id = ?", storeID).Delete(&schema.DorayakiStoreStock{}).Error; err != nil {
		return err
	}

	return nil
}
