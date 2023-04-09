package schema

type CreateDorayakiResponseDTO struct {
	Response
	Data Dorayaki `json:"data"`
}

type GetDorayakisResponseDTO struct {
	Response
	Data []Dorayaki `json:"data"`
}

type GetDorayakiResponseDTO struct {
	Response
	Data Dorayaki `json:"data"`
}

type UpdateDorayakiResponseDTO struct {
	Response
	Data Dorayaki `json:"data"`
}

type DeleteDorayakiResponseDTO struct {
	Response
}
