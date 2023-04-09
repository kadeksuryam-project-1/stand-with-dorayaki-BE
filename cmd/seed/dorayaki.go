package seed

import (
	"backend/config"
	"backend/internal/schema"

	"gorm.io/gorm"
)

func SeedDorayaki(db *gorm.DB) {
	bucketAddress := config.C.BucketAddress

	var seedDorayaki = []schema.Dorayaki{
		{Flavor: "chocolate", Description: "It has a soft texture", Image: bucketAddress + "assets/dorayaki/chocolate.png"},
		{Flavor: "oreo", Description: "It's a mix of some Oreos", Image: bucketAddress + "assets/dorayaki/oreo.png"},
		{Flavor: "strawberry", Description: "It's a mix of strawberries", Image: bucketAddress + "assets/dorayaki/strawberry.png"},
		{Flavor: "ice cream", Description: "It's a mixture of several ice creams", Image: bucketAddress + "assets/dorayaki/ice-cream.png"},
		{Flavor: "matcha", Description: "It has a thick matcha taste", Image: bucketAddress + "assets/dorayaki/matcha.png"},
	}

	tx := db.Begin()

	for _, dorayaki := range seedDorayaki {
		if err := tx.Create(&dorayaki).Error; err != nil {
			tx.Rollback()
			return
		}
	}

	tx.Commit()
}
