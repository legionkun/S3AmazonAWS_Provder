package inter

import (
	"ImageS3-Service/phuoc/dto"
	"ImageS3-Service/phuoc/models"

	"github.com/gin-gonic/gin"
)

type ProductRepository interface {
	InsertProductByRequest(p dto.ProductInsert_Request) (*models.Product, error)
}

type ProController interface {
	InsertProductByRequest(context *gin.Context)
}

type ProService interface {
	InsertProductByRequest(p dto.ProductInsert_Request) (*models.Product, error)
}

type ImaRepository interface {
	AddNewImage(i dto.GetDataFormImage) error
	GetImageByIdProduct(id uint64, proId uint64) (*dto.GetDatabyIdAndProID, error)
}

type ImageService interface {
	AddNewImage(i dto.GetDataFormImage) error
	GetData(id uint64, proId uint64) (*dto.GetDatabyIdAndProID, error)
}

type ImageController interface {
	AddNewImage(context *gin.Context)
	GetImageByID(context *gin.Context)
}

type S3Provider interface {
	SetUps3ProviderConfig()
}

type S3ProviderService interface {
	SaveFileUpload(data []byte, dst string) map[string]interface{}
}
