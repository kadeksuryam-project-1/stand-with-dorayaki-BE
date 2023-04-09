package schema

type Response struct {
	Success bool `json:"success"`
}
type Error struct {
	Response
	Message string `json:"message"`
}

func NewError(message string) Error {
	return Error{
		Response: Response{
			Success: false,
		},
		Message: message,
	}
}
