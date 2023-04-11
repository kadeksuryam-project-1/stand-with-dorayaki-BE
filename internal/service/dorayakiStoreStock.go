package service

import (
	"backend/internal/repository"
	"backend/internal/schema"
)

type IDorayakiStoreStockService interface {
	GetStocks(dorayakiID, storeID *int) ([]schema.DorayakiStoreStock, error)
	UpdateStock(stock, stockID int) (schema.DorayakiStoreStock, error)
}

type dorayakiStoreStockService struct {
	dorayakiStoreStockRepository repository.IDorayakiStoreStockRepository
}

func NewDorayakiStoreStockService(dorayakiStoreStockRepository repository.IDorayakiStoreStockRepository) IDorayakiStoreStockService {
	return &dorayakiStoreStockService{
		dorayakiStoreStockRepository: dorayakiStoreStockRepository,
	}
}

func (s *dorayakiStoreStockService) GetStocks(dorayakiID, storeID *int) ([]schema.DorayakiStoreStock, error) {
	dorayaki, err := s.dorayakiStoreStockRepository.GetStocks(dorayakiID, storeID)

	return dorayaki, err
}

func (s *dorayakiStoreStockService) UpdateStock(stock, stockID int) (schema.DorayakiStoreStock, error) {
	dorayaki, err := s.dorayakiStoreStockRepository.UpdateStock(stock, stockID)

	return dorayaki, err
}
