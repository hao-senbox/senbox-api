package controller

import (
	"net/http"
	"sen-global-api/internal/domain/request"
	"sen-global-api/internal/domain/response"
	"sen-global-api/internal/domain/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RoleClaimPermissionController struct {
	*usecase.GetRoleClaimPermissionUseCase
	*usecase.CreateRoleClaimPermissionUseCase
	*usecase.UpdateRoleClaimPermissionUseCase
	*usecase.DeleteRoleClaimPermissionUseCase
}

func (receiver *RoleClaimPermissionController) GetAllRoleClaimPermission(context *gin.Context) {
	policies, err := receiver.GetRoleClaimPermissionUseCase.GetAllRoleClaimPermission()
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.FailedResponse{
			Code:  http.StatusInternalServerError,
			Error: err.Error(),
		})

		return
	}

	var permissionListResponse []response.RoleClaimPermissionListResponseData
	for _, permission := range policies {
		permissionListResponse = append(permissionListResponse, response.RoleClaimPermissionListResponseData{
			ID:             permission.ID,
			PermissionName: permission.PermissionName,
		})
	}

	context.JSON(http.StatusOK, response.RoleClaimPermissionListResponse{
		Data: permissionListResponse,
	})
}

func (receiver *RoleClaimPermissionController) GetRoleClaimPermissionById(context *gin.Context) {
	permissionId := context.Param("id")
	if permissionId == "" {
		context.JSON(
			http.StatusBadRequest, response.FailedResponse{
				Code:  http.StatusBadRequest,
				Error: "permission id is required",
			},
		)
		return
	}

	id, err := strconv.ParseUint(permissionId, 10, 32)
	if err != nil {
		context.JSON(
			http.StatusBadRequest, response.FailedResponse{
				Code:  http.StatusBadRequest,
				Error: "permission id is invalid",
			},
		)
		return
	}

	userRoleClaimPermission, err := receiver.GetRoleClaimPermissionUseCase.GetRoleClaimPermissionById(request.GetRoleClaimPermissionByIdRequest{ID: uint(id)})
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.FailedResponse{
			Code:  http.StatusInternalServerError,
			Error: err.Error(),
		})

		return
	}

	context.JSON(http.StatusOK, response.SucceedResponse{
		Data: response.RoleClaimPermissionResponse{
			ID:             userRoleClaimPermission.ID,
			PermissionName: userRoleClaimPermission.PermissionName,
			Description:    userRoleClaimPermission.Description,
		},
	})
}

func (receiver *RoleClaimPermissionController) GetRoleClaimPermissionByName(context *gin.Context) {
	permissionName := context.Param("permission_name")
	if permissionName == "" {
		context.JSON(
			http.StatusBadRequest, response.FailedResponse{
				Code:  http.StatusBadRequest,
				Error: "permission name is required",
			},
		)
		return
	}

	userRoleClaimPermission, err := receiver.GetRoleClaimPermissionUseCase.GetRoleClaimPermissionByName(request.GetRoleClaimPermissionByNameRequest{PermissionName: permissionName})
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.FailedResponse{
			Code:  http.StatusInternalServerError,
			Error: err.Error(),
		})

		return
	}

	context.JSON(http.StatusOK, response.SucceedResponse{
		Data: response.RoleClaimPermissionResponse{
			ID:             userRoleClaimPermission.ID,
			PermissionName: userRoleClaimPermission.PermissionName,
			Description:    userRoleClaimPermission.Description,
		},
	})
}

func (receiver *RoleClaimPermissionController) CreateRoleClaimPermission(context *gin.Context) {
	var req request.CreateRoleClaimPermissionRequest
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, response.FailedResponse{
			Code:  http.StatusBadRequest,
			Error: err.Error(),
		})
		return
	}

	err := receiver.Create(req)
	if err != nil {
		context.JSON(http.StatusBadRequest, response.FailedResponse{
			Code:  http.StatusBadRequest,
			Error: err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, response.SucceedResponse{
		Code:    http.StatusOK,
		Message: "role claim permission was create successfully",
	})
}

func (receiver *RoleClaimPermissionController) UpdateRoleClaimPermission(context *gin.Context) {
	var req request.UpdateRoleClaimPermissionRequest
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, response.FailedResponse{
			Code:  http.StatusBadRequest,
			Error: err.Error(),
		})
		return
	}

	err := receiver.UpdateRoleClaimPermissionUseCase.UpdateRoleClaimPermission(req)
	if err != nil {
		context.JSON(http.StatusBadRequest, response.FailedResponse{
			Code:  http.StatusBadRequest,
			Error: err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, response.SucceedResponse{
		Code:    http.StatusOK,
		Message: "role claim permission was update successfully",
	})
}

func (receiver *RoleClaimPermissionController) DeleteRoleClaimPermission(context *gin.Context) {
	permissionId := context.Param("id")
	if permissionId == "" {
		context.JSON(
			http.StatusBadRequest, response.FailedResponse{
				Code:  http.StatusBadRequest,
				Error: "permission id is required",
			},
		)
		return
	}

	id, err := strconv.ParseUint(permissionId, 10, 32)
	if err != nil {
		context.JSON(
			http.StatusBadRequest, response.FailedResponse{
				Code:  http.StatusBadRequest,
				Error: "permission id is invalid",
			},
		)
		return
	}

	err = receiver.DeleteRoleClaimPermissionUseCase.DeleteRoleClaimPermission(request.DeleteRoleClaimPermissionRequest{ID: uint(id)})
	if err != nil {
		context.JSON(http.StatusBadRequest, response.FailedResponse{
			Code:  http.StatusBadRequest,
			Error: err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, response.SucceedResponse{
		Code:    http.StatusOK,
		Message: "role claim permission was delete successfully",
	})
}
