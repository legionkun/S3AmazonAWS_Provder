package server

import (
	"ImageS3-Service/phuoc/controller"
	"ImageS3-Service/phuoc/inter"
	"ImageS3-Service/phuoc/repositories"
	"ImageS3-Service/phuoc/usercase"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                = SetupDatabaseConnection()
	Pro_Repository inter.ProductRepository = repositories.NewProductRepository(db)
	Pro_Service    inter.ProService        = usercase.NewProductService(Pro_Repository)
	Pro_Controller inter.ProController     = controller.NewProductController(Pro_Service)

	Image_Repository inter.ImaRepository     = repositories.NewImageRepository(db)
	Image_Service    inter.ImageService      = usercase.NewImageService(Image_Repository)
	Pro_s3           inter.S3ProviderService = usercase.SetUps3ProviderConfig()
	Image_Controller inter.ImageController   = controller.NewImageController(Image_Service, Pro_s3)
)

func ServerLineUser() {
	defer CloseDatabaseConnection(db)
	r := gin.Default()

	ProRoutes := r.Group("/v1/product")
	{
		// ProRoutes.GET("", Pro_Controller.All)
		ProRoutes.POST("", Pro_Controller.InsertProductByRequest)
		// ProRoutes.PUT(":id", Pro_Controller.Update)
		// ProRoutes.DELETE(":id", Pro_Controller.Delete)
		// ProRoutes.GET(":id", Pro_Controller.GetByID)
	}
	ImageRoutes := r.Group("/v1/image")
	{
		ImageRoutes.POST("upload", Image_Controller.AddNewImage)
		ImageRoutes.GET(":id/:proid", Image_Controller.GetImageByID)
	}
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err := r.Run(":" + os.Getenv("SERVICE_PORT"))
	if err != nil {
		panic("[Error] failed to start Gin server due to: " + err.Error())
	}
}
