package usercase

import (
	"ImageS3-Service/phuoc/dto"
	"ImageS3-Service/phuoc/inter"
	"ImageS3-Service/phuoc/models"
)

type Prod_Service struct {
	ProductInter inter.ProductRepository
}

func NewProductService(prodRepo inter.ProductRepository) inter.ProService {
	return &Prod_Service{
		ProductInter: prodRepo,
	}
}

func (service *Prod_Service) InsertProductByRequest(p dto.ProductInsert_Request) (*models.Product, error) {
	return service.ProductInter.InsertProductByRequest(p)
}
