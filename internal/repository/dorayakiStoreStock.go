package repository

import (
	"backend/internal/schema"
	"errors"

	"gorm.io/gorm"
)

type IDorayakiStoreStockRepository interface {
	GetStocks(dorayakiID, storeID *int) ([]schema.DorayakiStoreStock, error)
	UpdateStock(stock, stockID int) (schema.DorayakiStoreStock, error)
	TransferStock(srcID, destID, amount int) error
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

func (r *dorayakiStoreStockRepository) TransferStock(srcID, destID, amount int) error {
	tx := r.db.Begin()

	var srcStock, destStock schema.DorayakiStoreStock
	if err := tx.Preload("Dorayaki").Preload("Store").Where("id = ?", srcID).Find(&srcStock).Error; err != nil {
		tx.Rollback()
		return err
	}
	if srcStock.Stock < amount {
		tx.Rollback()
		return errors.New("src stock cannot be less than transfer amount")
	}
	if err := tx.Preload("Dorayaki").Preload("Store").Where("id = ?", destID).Find(&destStock).Error; err != nil {
		tx.Rollback()
		return err
	}

	srcUpdatedStock := srcStock.Stock - amount
	destUpdatedStock := destStock.Stock + amount

	resultSrc := tx.Model(&schema.DorayakiStoreStock{}).Where("id = ?", srcID).Update("Stock", srcUpdatedStock)
	if resultSrc.Error != nil {
		tx.Rollback()
		return resultSrc.Error
	}

	resultDest := tx.Model(&schema.DorayakiStoreStock{}).Where("id = ?", destID).Update("Stock", destUpdatedStock)
	if resultDest.Error != nil {
		tx.Rollback()
		return resultDest.Error
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}

func NewDorayakiStoreStockRepository(db *gorm.DB) IDorayakiStoreStockRepository {
	return &dorayakiStoreStockRepository{
		db: db,
	}
}
