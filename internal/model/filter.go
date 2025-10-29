package model

type Filter struct {
	Status      Status `json:"status" binding:"required"`
	CreatedDate string `json:"createdDate" binding:"required"`
}
