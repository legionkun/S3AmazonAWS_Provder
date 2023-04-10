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

// func (service *Prod_Service) AllProduct(p int, g int, f string, sort string) *[]repositories.Result {
// 	return service.ProRepo.AllProduct(p, g, f, sort)
// }

func (service *Prod_Service) InsertProductByRequest(p dto.ProductInsert_Request) *models.Product {
	return service.ProductInter.InsertProductByRequest(p)
}

// func (service *Prod_Service) Update(p dto.ProductUpdate) *models.Product {
// 	book := &models.Product{}
// 	err := smapping.FillStruct(&book, smapping.MapFields(&p))
// 	if err != nil {
// 		log.Fatalf("Failed Map %v: ", err)
// 	}
// 	updated := service.ProRepo.UpdateProduct(p)
// 	return updated
// }

// func (service *Prod_Service) Delete(p models.Product) *models.Product {
// 	return service.ProRepo.DeleteProduct(p)
// }
// func (service *Prod_Service) FindByID(proID int) *models.Product {
// 	return service.ProRepo.FindByID(proID)
// }

// func (ser *Prod_Service) FindByNameGetID(id int) *[]models.Product {
// 	return ser.ProRepo.FindByNameGetbyID(id)
// }

// func (service *Prod_Service) ChangeNameToTypeID(Name string) int {
// 	if Name == "Tình yêu" {
// 		return int(1)
// 	} else if Name == "Kinh dị" {
// 		return int(2)
// 	} else if Name == "Trinh tham" {
// 		return int(3)
// 	} else if Name == "Tiểu Thuyết" {
// 		return int(4)
// 	}
// 	return 0
// }
