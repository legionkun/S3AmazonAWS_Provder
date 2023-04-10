package dto

type GetIdProduct struct {
	ProductId uint64 `json:"product_id" form:"product_id"`
}

type GetDataFormImage struct {
	Id          uint64
	Filename    string
	ContentType string
	Data        string
	Size        int64
}

type GetDatabyIdAndProID struct {
	Id          uint64 `json:"ID"`
	ProId       uint64 `json:"ProductID"`
	ContentType string `json:"ContentType"`
	Data        string `json:"Data"`
}
