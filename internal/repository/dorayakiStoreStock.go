package repository

import (
	"backend/internal/schema"

	"gorm.io/gorm"
)

type IDorayakiStoreStockRepository interface {
	GetStocks(dorayakiID, storeID *int) ([]schema.DorayakiStoreStock, error)
	UpdateStock(stock, stockID int) (schema.DorayakiStoreStock, error)
}

type dorayakiStoreStockRepository struct {
	db *gorm.DB
}

func (r *dorayakiStoreStockRepository) GetStocks(dorayakiID, storeID *int) ([]schema.DorayakiStoreStock, error) {
	var dorayakiStoreStocks []schema.DorayakiStoreStock
	query := r.db.Preload("Dorayaki").Preload("Store")

	if dorayakiID != nil {
		query = query.Where("dorayaki_id = ?", *dorayakiID)
	}

	if storeID != nil {
		query = query.Where("store_id = ?", *storeID)
	}

	err := query.Find(&dorayakiStoreStocks).Error

	return dorayakiStoreStocks, err
}

func (r *dorayakiStoreStockRepository) UpdateStock(stock, stockID int) (schema.DorayakiStoreStock, error) {
	tx := r.db.Begin()
	var stockItem schema.DorayakiStoreStock
	updateFields := map[string]interface{}{
		"Stock": stock,
	}
	result := tx.Model(&schema.DorayakiStoreStock{}).Where("id = ?", stockID).Updates(updateFields)
	if result.Error != nil {
		tx.Rollback()
		return stockItem, result.Error
	}

	if err := tx.Preload("Dorayaki").Preload("Store").Where("id = ?", stockID).Find(&stockItem).Error; err != nil {
		tx.Rollback()
		return stockItem, err
	}

	if err := tx.Commit().Error; err != nil {
		return stockItem, err
	}

	return stockItem, nil
}

func NewDorayakiStoreStockRepository(db *gorm.DB) IDorayakiStoreStockRepository {
	return &dorayakiStoreStockRepository{
		db: db,
	}
}
