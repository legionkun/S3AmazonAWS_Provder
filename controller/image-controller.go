package controller

import (
	"ImageS3-Service/phuoc/dto"
	"ImageS3-Service/phuoc/helper"
	"ImageS3-Service/phuoc/inter"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/olahol/go-imageupload"
)

type Image_Handle struct {
	ImaService inter.ImageService
	Pro_s3     inter.S3ProviderService
}

func NewImageController(Ima_Service inter.ImageService, S3Service inter.S3ProviderService) inter.ImageController {
	return &Image_Handle{
		ImaService: Ima_Service,
		Pro_s3:     S3Service,
	}
}

func (p *Image_Handle) AddNewImage(context *gin.Context) {
	img, err := imageupload.Process(context.Request, "file")
	if err != nil {
		panic(err)
	}
	imgConvertSmaller, err := imageupload.ThumbnailJPEG(img, 270, 270, 60)
	if err != nil {
		panic(err)
	}
	getID := context.Request.FormValue("productId")
	ProID, err := strconv.Atoi(getID)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request ", "Cannot convert your data - Please check again", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		info := dto.GetDataFormImage{}
		info.Id = uint64(ProID)
		info.Filename = img.Filename
		info.ContentType = imgConvertSmaller.ContentType
		domain := os.Getenv("S3Domain")
		info.Url = fmt.Sprintf(domain + "/" + info.Filename)
		info.Size = int64(imgConvertSmaller.Size)
		erro := p.Pro_s3.SaveFileUpload(imgConvertSmaller.Data, info.Filename)
		if erro != nil {
			res := helper.BuildErrorResponse("Failed to process request ", "Cannot save image to database or S3 - Please check again", map[string]interface{}{
				"Problem at S3": erro,
			})
			context.JSON(http.StatusBadRequest, res)
		} else {
			err := p.ImaService.AddNewImage(info)
			if err != nil {
				res := helper.BuildErrorResponse("Failed to process request ", "Cannot save image to database or S3 - Please check again", map[string]interface{}{
					"Problem at DB": err.Error(),
				})
				context.JSON(http.StatusBadRequest, res)
			} else {
				response := helper.BuildResponse(true, "OK", map[string]interface{}{
					"Product ID": getID,
					"URL":        info.Url,
				})
				context.JSON(http.StatusCreated, response)
			}
		}
	}
}

func (p *Image_Handle) GetImageByID(context *gin.Context) {
	ProID, err := strconv.ParseUint(context.Param("proid"), 0, 0)
	ID, erro := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get Product id", "No param Product were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	} else if erro != nil {
		response := helper.BuildErrorResponse("Failed to get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	} else {
		data, _ := p.ImaService.GetData(ID, ProID)
		if data == nil {
			response := helper.BuildErrorResponse("Failed to get id", "No data were found", helper.EmptyObj{})
			context.JSON(http.StatusOK, response)
		} else {
			context.JSON(http.StatusOK, data)
		}
	}
}
