package seed

import (
	"backend/internal/schema"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

func SeedDorayakiStoreStock(db *gorm.DB) {
	var dorayakis []schema.Dorayaki
	var stores []schema.Store
	var dorayakiStoreStocks []schema.DorayakiStoreStock

	db.Find(&dorayakis)
	db.Find(&stores)

	rand.Seed(time.Now().UnixNano())
	for _, dorayaki := range dorayakis {
		for _, store := range stores {
			dorayakiStoreStock := schema.DorayakiStoreStock{
				DorayakiId: dorayaki.ID,
				StoreId:    store.ID,
				Stock:      rand.Intn(101),
			}
			dorayakiStoreStocks = append(dorayakiStoreStocks, dorayakiStoreStock)
		}
	}

	tx := db.Begin()

	for _, dorayakiStoreStock := range dorayakiStoreStocks {
		if err := tx.Create(&dorayakiStoreStock).Error; err != nil {
			tx.Rollback()
			return
		}
	}

	tx.Commit()
}
