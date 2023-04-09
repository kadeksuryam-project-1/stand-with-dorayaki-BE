package service

import (
	"backend/internal/repository"
	"backend/internal/schema"
)

type IDorayakiStoreStockService interface {
	GetStocks() ([]schema.DorayakiStoreStock, error)
	UpdateStock(stock, dorayakiID, storeID int) (schema.DorayakiStoreStock, error)
}

type dorayakiStoreStockService struct {
	dorayakiStoreStockRepository repository.IDorayakiStoreStockRepository
}

func NewDorayakiStoreStockService(dorayakiStoreStockRepository repository.IDorayakiStoreStockRepository) IDorayakiStoreStockService {
	return &dorayakiStoreStockService{
		dorayakiStoreStockRepository: dorayakiStoreStockRepository,
	}
}

func (s *dorayakiStoreStockService) GetStocks() ([]schema.DorayakiStoreStock, error) {
	dorayakis, err := s.dorayakiStoreStockRepository.GetAll()

	return dorayakis, err
}

func (s *dorayakiStoreStockService) UpdateStock(stock, dorayakiID, storeID int) (schema.DorayakiStoreStock, error) {
	dorayaki, err := s.dorayakiStoreStockRepository.UpdateStock(stock, dorayakiID, storeID)

	return dorayaki, err
}
