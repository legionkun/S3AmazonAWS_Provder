package usercase

import (
	"ImageS3-Service/phuoc/dto"
	"ImageS3-Service/phuoc/inter"
)

type Ima_Service struct {
	ImageRepo inter.ImaRepository
}

func NewImageService(Imarepo inter.ImaRepository) inter.ImageService {
	return &Ima_Service{
		ImageRepo: Imarepo,
	}
}

func (service *Ima_Service) AddNewImage(i dto.GetDataFormImage) {
	service.ImageRepo.AddNewImage(i)
}

func (service *Ima_Service) GetData(id uint64, proId uint64) (data dto.GetDatabyIdAndProID) {
	return service.ImageRepo.GetImageByIdProduct(id, proId)
}
