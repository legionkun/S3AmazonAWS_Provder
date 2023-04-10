package repositories

import (
	"ImageS3-Service/phuoc/dto"
	"ImageS3-Service/phuoc/inter"
	"ImageS3-Service/phuoc/models"

	"gorm.io/gorm"
)

type ProductConnection struct {
	connection   *gorm.DB
	ProductInter inter.ProductRepository
}

func NewProductRepository(db *gorm.DB) *ProductConnection {
	return &ProductConnection{
		connection: db,
	}
}

// func (db *ProductConnection) AllProduct(p int, g int, f string, sort string) *[]Result {
// 	var pros []Result
// 	//var typ []models.Images
// 	a := p * g
// 	if sort == "id[DESC]" {
// 		db.connection.Limit(g).Offset(a-g).Select("products.id ,products.title, product_types.name").
// 			Joins("left join products on product_types.id = products.type_id").Where("product_types.name =? ", f).Order("id desc").Scan(&pros)
// 		return &pros
// 	} else if sort == "name[DESC]" {
// 		db.connection.Limit(g).Offset(a-g).Select("products.id ,products.title, product_types.name").
// 			Joins("left join products on product_types.id = products.type_id").Where("product_types.name =? ", f).Order("title desc").Scan(&pros)
// 		return &pros
// 	} else if sort == "name[ASC]" {
// 		db.connection.Limit(g).Offset(a-g).Select("products.id ,products.title, product_types.name").
// 			Joins("left join products on product_types.id = products.type_id").Where("product_types.name =? ", f).Order("title asc").Scan(&pros)
// 		return &pros
// 	} else {
// 		db.connection.Limit(g).Offset(a-g).Select("products.id ,products.title, product_types.name").
// 			Joins("left join products on product_types.id = products.type_id").Where("product_types.name =? ", f).Scan(&pros)
// 		return &pros
// 	}
// }

func (db *ProductConnection) InsertProductByRequest(p dto.ProductInsert_Request) *models.Product {
	producing := &models.Product{}
	producing.Name = p.Name
	producing.Description = p.Desciption
	db.connection.Create(&producing)
	return producing
}

// func (db *ProductConnection) UpdateProduct(p dto.ProductUpdate) *models.Product {
// 	var product *models.Product = db.FindByID(int(p.ID))
// 	db.connection.First(&product)
// 	product.Title = p.Title
// 	//product.TypeID = p.TypeID
// 	db.connection.Save(&product)
// 	return product
// }

// func (db *ProductConnection) DeleteProduct(p models.Product) *models.Product {
// 	db.connection.Delete(&p)
// 	return &p
// }

// func (db *ProductConnection) FindByID(proID int) *models.Product {
// 	var product models.Product
// 	db.connection.Find(&product, proID)
// 	return &product
// }
// func (db *ProductConnection) FindByNameGetbyID(typeID int) *[]models.Product {
// 	var typ []models.Product
// 	db.connection.Where("type_id = ?", typeID).Find(&typ)
// 	return &typ
// }
