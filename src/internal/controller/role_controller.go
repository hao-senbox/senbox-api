package controller

import (
	"net/http"
	"sen-global-api/internal/domain/request"
	"sen-global-api/internal/domain/response"
	"sen-global-api/internal/domain/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	*usecase.GetRoleUseCase
	*usecase.CreateRoleUseCase
	*usecase.UpdateRoleUseCase
	*usecase.DeleteRoleUseCase
}

func (receiver *RoleController) GetAllRole(context *gin.Context) {
	roles, err := receiver.GetRoleUseCase.GetAllRole()
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.FailedResponse{
			Error: response.Cause{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		})

		return
	}

	var roleListResponse []response.RoleListResponseData
	for _, role := range roles {
		roleListResponse = append(roleListResponse, response.RoleListResponseData{
			ID:       role.ID,
			RoleName: role.RoleName,
		})
	}

	context.JSON(http.StatusOK, response.RoleListResponse{
		Data: roleListResponse,
	})
}

func (receiver *RoleController) GetRoleById(context *gin.Context) {
	roleId := context.Param("id")
	if roleId == "" {
		context.JSON(
			http.StatusBadRequest, response.FailedResponse{
				Error: response.Cause{
					Code:    http.StatusBadRequest,
					Message: "Role Id is required",
				},
			},
		)
		return
	}

	id, err := strconv.ParseUint(roleId, 10, 32)
	if err != nil {
		context.JSON(
			http.StatusBadRequest, response.FailedResponse{
				Error: response.Cause{
					Code:    http.StatusBadRequest,
					Message: "Role Id is invalid",
				},
			},
		)
		return
	}

	userRole, err := receiver.GetRoleUseCase.GetRoleById(request.GetRoleByIdRequest{ID: uint(id)})
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.FailedResponse{
			Error: response.Cause{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		})

		return
	}

	context.JSON(http.StatusOK, response.SucceedResponse{
		Data: response.RoleResponse{
			ID:          userRole.ID,
			RoleName:    userRole.RoleName,
			Description: userRole.Description,
		},
	})
}

func (receiver *RoleController) GetRoleByName(context *gin.Context) {
	roleName := context.Param("role_name")
	if roleName == "" {
		context.JSON(
			http.StatusBadRequest, response.FailedResponse{
				Error: response.Cause{
					Code:    http.StatusBadRequest,
					Message: "role name is required",
				},
			},
		)
		return
	}

	userRole, err := receiver.GetRoleUseCase.GetRoleByName(request.GetRoleByNameRequest{RoleName: roleName})
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.FailedResponse{
			Error: response.Cause{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		})

		return
	}

	context.JSON(http.StatusOK, response.SucceedResponse{
		Data: response.RoleResponse{
			ID:          userRole.ID,
			RoleName:    userRole.RoleName,
			Description: userRole.Description,
		},
	})
}

func (receiver *RoleController) CreateRole(context *gin.Context) {
	var req request.CreateRoleRequest
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, response.FailedResponse{
			Error: response.Cause{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			},
		})
		return
	}

	err := receiver.CreateRoleUseCase.Create(req)
	if err != nil {
		context.JSON(http.StatusBadRequest, response.FailedResponse{
			Error: response.Cause{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			},
		})
		return
	}
	context.JSON(http.StatusOK, response.SucceedResponse{
		Data: response.Cause{
			Code:    http.StatusOK,
			Message: "user role was create successfully",
		},
	})
}

func (receiver *RoleController) UpdateRole(context *gin.Context) {
	var req request.UpdateRoleRequest
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, response.FailedResponse{
			Error: response.Cause{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			},
		})
		return
	}

	err := receiver.UpdateRoleUseCase.UpdateRole(req)
	if err != nil {
		context.JSON(http.StatusBadRequest, response.FailedResponse{
			Error: response.Cause{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			},
		})
		return
	}

	context.JSON(http.StatusOK, response.SucceedResponse{
		Data: response.Cause{
			Code:    http.StatusOK,
			Message: "role was update successfully",
		},
	})
}

func (receiver *RoleController) DeleteRole(context *gin.Context) {
	roleId := context.Param("id")
	if roleId == "" {
		context.JSON(
			http.StatusBadRequest, response.FailedResponse{
				Error: response.Cause{
					Code:    http.StatusBadRequest,
					Message: "Role Id is required",
				},
			},
		)
		return
	}

	id, err := strconv.ParseUint(roleId, 10, 32)
	if err != nil {
		context.JSON(
			http.StatusBadRequest, response.FailedResponse{
				Error: response.Cause{
					Code:    http.StatusBadRequest,
					Message: "Role Id is invalid",
				},
			},
		)
		return
	}

	err = receiver.DeleteRoleUseCase.DeleteRole(request.DeleteRoleRequest{ID: uint(id)})
	if err != nil {
		context.JSON(http.StatusBadRequest, response.FailedResponse{
			Error: response.Cause{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			},
		})
		return
	}

	context.JSON(http.StatusOK, response.SucceedResponse{
		Data: response.Cause{
			Code:    http.StatusOK,
			Message: "role was delete successfully",
		},
	})
}
