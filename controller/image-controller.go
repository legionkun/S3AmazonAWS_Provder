package controller

import (
	"ImageS3-Service/phuoc/dto"
	"ImageS3-Service/phuoc/helper"
	"ImageS3-Service/phuoc/inter"
	"net/http"
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
		data := imgConvertSmaller.DataURI()
		info.Data = data
		info.Size = int64(imgConvertSmaller.Size)
		p.ImaService.AddNewImage(info)
		err := p.Pro_s3.SaveFileUpload(imgConvertSmaller.Data, info.Filename)
		if err != nil {
			res := helper.BuildErrorResponse("Failed to process request ", "Cannot upload your data to S3 - Please check again", helper.EmptyObj{})
			context.JSON(http.StatusBadRequest, res)
		}
		response := helper.BuildResponse(true, "OK", map[string]interface{}{
			"Product ID": getID,
		})
		context.JSON(http.StatusCreated, response)
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
		data := p.ImaService.GetData(ID, ProID)
		context.JSON(http.StatusOK, data)
	}
}
