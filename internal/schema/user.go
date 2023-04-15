package schema

type User struct {
	ID    int    `json:"id" gorm:"primaryKey;column:id;type:integer"`
	Email string `json:"email" gorm:"primaryKey;column:email;type:varchar(255)"`
}
