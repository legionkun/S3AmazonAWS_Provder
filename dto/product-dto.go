package dto

type ProductInsert_Request struct {
	Name       string `json:"name" form:"title" binding:"required"`
	Desciption string `json:"desciption" form:"description" binding:"required"`
}
