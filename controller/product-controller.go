package controller

import (
	"ImageS3-Service/phuoc/dto"
	"ImageS3-Service/phuoc/helper"
	"ImageS3-Service/phuoc/inter"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Pro_Controller struct {
	Service inter.ProService
}

func NewProductController(ProService inter.ProService) inter.ProController {
	return &Pro_Controller{
		Service: ProService,
	}
}

func (p *Pro_Controller) InsertProductByRequest(context *gin.Context) {
	var createProduct dto.ProductInsert_Request
	errDTO := context.ShouldBindJSON(&createProduct)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request ", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		result, err := p.Service.InsertProductByRequest(createProduct)
		if err != nil {
			context.JSON(http.StatusInternalServerError, err.Error())
		} else {
			response := helper.BuildResponse(true, "OK", result)
			context.JSON(http.StatusOK, response)
		}
	}
}
