package dto

type BadRequestResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
