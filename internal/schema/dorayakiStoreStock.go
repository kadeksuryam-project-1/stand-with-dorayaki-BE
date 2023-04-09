package schema

import "time"

type DorayakiStoreStock struct {
	ID         int       `json:"id" gorm:"primaryKey;column:id;type:integer"`
	DorayakiId int       `json:"dorayaki_id" gorm:"column:dorayaki_id;not null;type:integer"`
	StoreId    int       `json:"store_id" gorm:"column:store_id;not null;type:integer"`
	Stock      int       `json:"stock" gorm:"column:stock;not null;type:integer"`
	CreatedAt  time.Time `json:"created_at" gorm:"column:created_at;not null"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"column:updated_at;not null"`

	Dorayaki Dorayaki `gorm:"foreignKey:DorayakiId"`
	Store    Store    `gorm:"foreignKey:StoreId"`
}

func (DorayakiStoreStock) TableName() string {
	return "dorayaki_store_stocks"
}
