package repository

import (
	"backend/internal/schema"

	"gorm.io/gorm"
)

type IDorayakiStoreStockRepository interface {
	GetAll() ([]schema.DorayakiStoreStock, error)
	UpdateStock(stock, dorayakiID, storeID int) (schema.DorayakiStoreStock, error)
}

type dorayakiStoreStockRepository struct {
	db *gorm.DB
}

func (r *dorayakiStoreStockRepository) GetAll() ([]schema.DorayakiStoreStock, error) {
	var dorayakiStoreStocks []schema.DorayakiStoreStock

	err := r.db.Preload("Dorayaki").Preload("Store").Find(&dorayakiStoreStocks).Error

	return dorayakiStoreStocks, err
}

func (r *dorayakiStoreStockRepository) UpdateStock(stock, dorayakiID, storeID int) (schema.DorayakiStoreStock, error) {
	tx := r.db.Begin()
	var stockItem schema.DorayakiStoreStock
	updateFields := map[string]interface{}{
		"Stock": stock,
	}
	result := tx.Model(&schema.DorayakiStoreStock{}).Where("dorayaki_id = ? AND store_id = ?", dorayakiID, storeID).Updates(updateFields)
	if result.Error != nil {
		tx.Rollback()
		return stockItem, result.Error
	}

	if err := tx.Preload("Dorayaki").Preload("Store").Where("dorayaki_id = ? AND store_id = ?", dorayakiID, storeID).Find(&stockItem).Error; err != nil {
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
