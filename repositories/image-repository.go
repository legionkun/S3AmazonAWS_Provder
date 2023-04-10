package repositories

import (
	"ImageS3-Service/phuoc/dto"
	"ImageS3-Service/phuoc/models"

	"gorm.io/gorm"
)

type ImageConnnection struct {
	connection *gorm.DB
}

func NewImageRepository(db *gorm.DB) *ImageConnnection {
	return &ImageConnnection{
		connection: db,
	}
}

func (db *ImageConnnection) AddNewImage(i dto.GetDataFormImage) {
	img := models.Image{}
	img.ProductID = i.Id
	img.ContentType = i.ContentType
	img.Size = i.Size
	img.Url = i.Filename
	img.Data = i.Data
	db.connection.Create(&img)
}

func (db *ImageConnnection) GetImageByIdProduct(id uint64, proId uint64) (data dto.GetDatabyIdAndProID) {
	var img models.Image
	db.connection.Where("product_id = ?", proId).Find(&img, id)
	data.Id = img.Id
	data.ContentType = img.ContentType
	data.ProId = img.ProductID
	data.Data = img.Data
	return data
}
