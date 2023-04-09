package schema

import "mime/multipart"

type DorayakiForm struct {
	Flavor      string                `json:"flavor" validate:"required"`
	Description string                `json:"description" validate:"required"`
	Image       *multipart.FileHeader `json:"image"`
}

type StoreForm struct {
	Name        string                `json:"name" validate:"required"`
	Street      string                `json:"street" validate:"required"`
	Subdistrict string                `json:"subdistrict" validate:"required"`
	District    string                `json:"district" validate:"required"`
	Province    string                `json:"province" validate:"required"`
	Image       *multipart.FileHeader `json:"image"`
}
