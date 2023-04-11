package repositories

import (
	"ImageS3-Service/phuoc/dto"
	"ImageS3-Service/phuoc/models"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ImageConnnection struct {
	connection *gorm.DB
}

func NewImageRepository(db *gorm.DB) *ImageConnnection {
	return &ImageConnnection{
		connection: db,
	}
}

func (db *ImageConnnection) AddNewImage(i dto.GetDataFormImage) error {
	img := models.Image{}
	img.ProductID = i.Id
	img.ContentType = i.ContentType
	img.Size = i.Size
	img.Url = i.Url
	img.Name = i.Filename
	err := db.connection.Create(&img)
	if err != nil {
		return err.Error
	}
	return nil
}

func (db *ImageConnnection) GetImageByIdProduct(id uint64, proId uint64) (*dto.GetDatabyIdAndProID, error) {
	var img models.Image
	var data dto.GetDatabyIdAndProID
	db.connection.Where("product_id = ?", proId).Find(&img, id)
	data.Id = img.Id
	data.ContentType = img.ContentType
	data.ProId = img.ProductID
	data.Url = img.Url
	if data.Id == 0 {
		return nil, logger.ErrRecordNotFound
	}
	return &data, nil
}
