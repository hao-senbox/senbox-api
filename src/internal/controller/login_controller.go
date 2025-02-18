package controller

import (
	"net/http"
	"sen-global-api/internal/domain/request"
	"sen-global-api/internal/domain/response"
	"sen-global-api/internal/domain/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginController struct {
	DBConn *gorm.DB
	usecase.AuthorizeUseCase
}

func (receiver LoginController) Login(c *gin.Context) {
	var req request.UserLoginRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest, response.FailedResponse{
				Error: response.Cause{
					Code:    http.StatusBadRequest,
					Message: err.Error(),
				},
			},
		)
		return
	}

	data, err := receiver.AuthorizeUseCase.LoginInputDao(req)

	if err != nil {
		c.JSON(
			http.StatusBadRequest, response.FailedResponse{
				Error: response.Cause{
					Code:    http.StatusBadRequest,
					Message: err.Error(),
				},
			},
		)
		return
	}

	c.JSON(http.StatusOK, response.SucceedResponse{
		Data: data,
	})
}

// Login godoc
// @Summary      Retrieve a token
// @Description  login using username and password
// @Tags         Admin
// @Accept       json
// @Produce      json
// @Param req body request.LoginInputReq true "Login Params"
// @Success      200  {object}  response.LoginResponse
// @Failure      400  {object}  response.FailedResponse
// @Failure      404  {object}  response.FailedResponse
// @Failure      500  {object}  response.FailedResponse
// @Router       /v1/login [post]
func (receiver LoginController) UserLogin(c *gin.Context) {
	var req request.UserLoginRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest, response.FailedResponse{
				Error: response.Cause{
					Code:    http.StatusBadRequest,
					Message: err.Error(),
				},
			},
		)
		return
	}

	data, err := receiver.AuthorizeUseCase.UserLoginUsecase(req)

	if err != nil {
		c.JSON(
			http.StatusBadRequest, response.FailedResponse{
				Error: response.Cause{
					Code:    http.StatusBadRequest,
					Message: err.Error(),
				},
			},
		)
		return
	}

	c.JSON(http.StatusOK, response.LoginResponse{
		Data: data,
	})
}
