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

func (db *ProductConnection) InsertProductByRequest(p dto.ProductInsert_Request) (*models.Product, error) {
	producing := &models.Product{}
	producing.Name = p.Name
	producing.Description = p.Desciption
	err := db.connection.Create(&producing)
	if err != nil {
		return nil, err.Error
	}
	return producing, nil
}
