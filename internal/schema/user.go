package schema

import "time"

type User struct {
	ID        int       `json:"id" gorm:"primaryKey;column:id;type:integer"`
	Email     string    `json:"email" gorm:"primaryKey;column:email;type:varchar(255)"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;not null"`
}
