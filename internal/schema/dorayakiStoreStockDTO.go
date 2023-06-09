package schema

type StockRequestDTO struct {
	Stock int `json:"stock" validate:"required,gte=0"`
}

type GetStocksResponseDTO struct {
	Response
	Data []DorayakiStoreStock `json:"data"`
}

type UpdateStockResponseDTO struct {
	Response
	Data DorayakiStoreStock `json:"data"`
}
