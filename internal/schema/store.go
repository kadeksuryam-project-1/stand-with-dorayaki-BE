package schema

import "time"

type Store struct {
	ID          int       `json:"id" gorm:"primaryKey;column:id;type:integer"`
	Name        string    `json:"name" gorm:"column:name;not null;type:varchar(255)"`
	Street      string    `json:"street" gorm:"column:street;not null;type:varchar(1000)"`
	Subdistrict string    `json:"subdistrict" gorm:"column:subdistrict;not null;type:varchar(1000)"`
	District    string    `json:"district" gorm:"column:district;not null;type:varchar(1000)"`
	Province    string    `json:"province" gorm:"column:province;not null;type:varchar(1000)"`
	Image       string    `json:"image" gorm:"column:image;not null;type:varchar(1000)"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at;not null"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at;not null"`
}

func (Store) TableName() string {
	return "stores"
}
