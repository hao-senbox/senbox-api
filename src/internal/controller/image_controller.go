package controller

import (
	"bufio"
	"net/http"
	"sen-global-api/internal/domain/request"
	"sen-global-api/internal/domain/response"
	"sen-global-api/internal/domain/usecase"
	"sen-global-api/pkg/randx"

	"github.com/gin-gonic/gin"
)

type ImageController struct {
	*usecase.GetImageUseCase
	*usecase.UploadImageUseCase
	*usecase.DeleteImageUseCase
}

type getUrlByKeyRequest struct {
	Key string `json:"key" binding:"required"`
}

func (receiver *ImageController) GetUrlByKey(context *gin.Context) {
	var req getUrlByKeyRequest
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, response.FailedResponse{
			Code:  http.StatusBadRequest,
			Error: err.Error(),
		})
		return
	}

	url, err := receiver.GetImageUseCase.GetUrlByKey(req.Key)
	if err != nil {
		context.JSON(http.StatusBadRequest, response.FailedResponse{
			Code:  http.StatusBadRequest,
			Error: err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, response.SucceedResponse{
		Code:    http.StatusOK,
		Message: "image was get successfully",
		Data:    *url,
	})
}

func (receiver *ImageController) CreateImage(context *gin.Context) {
	fileHeader, err := context.FormFile("file")
	if err != nil {
		context.JSON(http.StatusBadRequest, response.FailedResponse{
			Code:  http.StatusBadRequest,
			Error: err.Error(),
		})
		return
	}

	folder := context.DefaultPostForm("folder", "img")
	fileName := context.DefaultPostForm("file_name", randx.GenString(10))

	file, err := fileHeader.Open()
	if err != nil {
		context.JSON(http.StatusBadRequest, response.FailedResponse{
			Code:  http.StatusBadRequest,
			Error: err.Error(),
		})
		return
	}

	defer file.Close()

	dataBytes := make([]byte, fileHeader.Size)
	if _, err := bufio.NewReader(file).Read(dataBytes); err != nil {
		context.JSON(http.StatusBadRequest, response.FailedResponse{
			Code:  http.StatusBadRequest,
			Error: err.Error(),
		})
		return
	}

	url, img, err := receiver.UploadImageUseCase.UploadImage(dataBytes, folder, fileHeader.Filename, fileName)
	if err != nil {
		context.JSON(http.StatusBadRequest, response.FailedResponse{
			Code:  http.StatusBadRequest,
			Error: err.Error(),
		})
		return
	}

	if url == nil {
		context.JSON(http.StatusBadRequest, response.FailedResponse{
			Code:  http.StatusBadRequest,
			Error: "image was not created",
		})
		return
	}

	context.JSON(http.StatusOK, response.SucceedResponse{
		Code:    http.StatusOK,
		Message: "image was create successfully",
		Data: response.ImageResponse{
			ImageName: img.ImageName,
			Key:       img.Key,
			Extension: img.Extension,
			Url:       *url,
			Width:     img.Width,
			Height:    img.Height,
		},
	})
}

func (receiver *ImageController) DeleteImage(context *gin.Context) {
	var req request.DeleteImageRequest
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, response.FailedResponse{
			Code:  http.StatusBadRequest,
			Error: "key is required",
		})
		return
	}

	err := receiver.DeleteImageUseCase.DeleteImage(req.Key)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.FailedResponse{
			Code:  http.StatusInternalServerError,
			Error: err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, response.SucceedResponse{
		Code:    http.StatusOK,
		Message: "image deleted successfully",
	})
}
