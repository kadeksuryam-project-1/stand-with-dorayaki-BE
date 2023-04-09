package schema

type CreateStoreResponseDTO struct {
	Response
	Data Store `json:"data"`
}

type GetStoresResponseDTO struct {
	Response
	Data []Store `json:"data"`
}

type GetStoreResponseDTO struct {
	Response
	Data Store `json:"data"`
}

type UpdateStoreResponseDTO struct {
	Response
	Data Store `json:"data"`
}

type DeleteStoreResponseDTO struct {
	Response
}
