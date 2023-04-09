package schema

import "time"

type Dorayaki struct {
	ID          int       `json:"id" gorm:"primaryKey;column:id;type:integer"`
	Flavor      string    `json:"flavor" gorm:"column:flavor;not null;type:varchar(255)"`
	Description string    `json:"description" gorm:"column:description;type:varchar(1000)"`
	Image       string    `json:"image" gorm:"column:image;not null;type:varchar(1000)"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at;not null"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at;not null"`
}

func (Dorayaki) TableName() string {
	return "dorayakis"
}
