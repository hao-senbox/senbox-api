package router

import (
	"sen-global-api/config"
	"sen-global-api/internal/controller"
	"sen-global-api/internal/data/repository"
	"sen-global-api/internal/domain/usecase"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupUserRoutes(engine *gin.Engine, dbConn *gorm.DB, config config.AppConfig) {
	sessionRepository := repository.SessionRepository{
		AuthorizeEncryptKey: config.AuthorizeEncryptKey,

		TokenExpireTimeInHour: time.Duration(config.TokenExpireDurationInHour),
	}

	userEntityController := &controller.UserEntityController{
		GetUserEntityUseCase: &usecase.GetUserEntityUseCase{
			UserEntityRepository: &repository.UserEntityRepository{DBConn: dbConn},
		},
		CreateUserEntityUseCase: &usecase.CreateUserEntityUseCase{
			UserEntityRepository: &repository.UserEntityRepository{DBConn: dbConn},
		},
		UpdateUserEntityUseCase: &usecase.UpdateUserEntityUseCase{
			UserEntityRepository: &repository.UserEntityRepository{DBConn: dbConn},
		},
		AuthorizeUseCase: &usecase.AuthorizeUseCase{
			UserEntityRepository: &repository.UserEntityRepository{DBConn: dbConn},
			DeviceRepository:     &repository.DeviceRepository{DBConn: dbConn},
			SessionRepository:    sessionRepository,
		},
	}

	userRoleController := &controller.RoleController{
		GetRoleUseCase: &usecase.GetRoleUseCase{
			RoleRepository: &repository.RoleRepository{DBConn: dbConn},
		},
		CreateRoleUseCase: &usecase.CreateRoleUseCase{
			RoleRepository: &repository.RoleRepository{DBConn: dbConn},
		},
		UpdateRoleUseCase: &usecase.UpdateRoleUseCase{
			RoleRepository: &repository.RoleRepository{DBConn: dbConn},
		},
		DeleteRoleUseCase: &usecase.DeleteRoleUseCase{
			RoleRepository: &repository.RoleRepository{DBConn: dbConn},
		},
	}

	roleClaimController := &controller.RoleClaimController{
		GetRoleClaimUseCase: &usecase.GetRoleClaimUseCase{
			RoleClaimRepository: &repository.RoleClaimRepository{DBConn: dbConn},
		},
		CreateRoleClaimUseCase: &usecase.CreateRoleClaimUseCase{
			RoleClaimRepository: &repository.RoleClaimRepository{DBConn: dbConn},
		},
		UpdateRoleClaimUseCase: &usecase.UpdateRoleClaimUseCase{
			RoleClaimRepository: &repository.RoleClaimRepository{DBConn: dbConn},
		},
		DeleteRoleClaimUseCase: &usecase.DeleteRoleClaimUseCase{
			RoleClaimRepository: &repository.RoleClaimRepository{DBConn: dbConn},
		},
	}

	roleClaimPermissionController := &controller.RoleClaimPermissionController{
		GetRoleClaimPermissionUseCase: &usecase.GetRoleClaimPermissionUseCase{
			RoleClaimPermissionRepository: &repository.RoleClaimPermissionRepository{DBConn: dbConn},
		},
		CreateRoleClaimPermissionUseCase: &usecase.CreateRoleClaimPermissionUseCase{
			RoleClaimPermissionRepository: &repository.RoleClaimPermissionRepository{DBConn: dbConn},
		},
		UpdateRoleClaimPermissionUseCase: &usecase.UpdateRoleClaimPermissionUseCase{
			RoleClaimPermissionRepository: &repository.RoleClaimPermissionRepository{DBConn: dbConn},
		},
		DeleteRoleClaimPermissionUseCase: &usecase.DeleteRoleClaimPermissionUseCase{
			RoleClaimPermissionRepository: &repository.RoleClaimPermissionRepository{DBConn: dbConn},
		},
	}

	userAccess := engine.Group("v1/")
	{
		loginController := &controller.LoginController{DBConn: dbConn,
			AuthorizeUseCase: usecase.AuthorizeUseCase{
				UserEntityRepository: &repository.UserEntityRepository{DBConn: dbConn},
				DeviceRepository:     &repository.DeviceRepository{DBConn: dbConn},
				SessionRepository:    sessionRepository,
			},
		}
		userAccess.POST("/login", loginController.UserLogin)
	}

	user := engine.Group("v1/user")
	{
		user.GET("/all", userEntityController.GetAllUserEntity)
		user.GET("/:id", userEntityController.GetUserEntityById)
		user.GET("/name/:username", userEntityController.GetUserEntityByName)
		user.GET("/:id/children", userEntityController.GetChildrenOfGuardian)

		user.POST("/init", userEntityController.CreateUserEntity)
		user.POST("/update", userEntityController.UpdateUserEntity)
		user.POST("/role/update", userEntityController.UpdateUserRole)
	}

	userRole := engine.Group("v1/user-role")
	{
		userRole.GET("/:organization_id/all", userRoleController.GetAllRoleByOrganization)
		userRole.GET("/:id", userRoleController.GetRoleById)
		userRole.GET("/name/:role_name", userRoleController.GetRoleByName)

		userRole.POST("/init", userRoleController.CreateRole)
		userRole.POST("/", userRoleController.UpdateRole)

		userRole.DELETE("/:id", userRoleController.DeleteRole)
	}

	roleClaim := engine.Group("v1/role-claim")
	{
		roleClaim.GET("/all", roleClaimController.GetAllRoleClaim)
		roleClaim.GET("/all/:role_id", roleClaimController.GetAllRoleClaimByRole)
		roleClaim.GET("/:id", roleClaimController.GetRoleClaimById)
		roleClaim.GET("/name/:claim_name", roleClaimController.GetRoleClaimByName)

		roleClaim.POST("/init", roleClaimController.CreateRoleClaim)
		roleClaim.POST("/", roleClaimController.UpdateRoleClaim)

		roleClaim.DELETE("/:id", roleClaimController.DeleteRoleClaim)
	}

	roleClaimPermission := engine.Group("v1/role-policy")
	{
		roleClaimPermission.GET("/all", roleClaimPermissionController.GetAllRoleClaimPermission)
		roleClaimPermission.GET("/:id", roleClaimPermissionController.GetRoleClaimPermissionById)
		roleClaimPermission.GET("/name/:policy_name", roleClaimPermissionController.GetRoleClaimPermissionByName)

		roleClaimPermission.POST("/init", roleClaimPermissionController.CreateRoleClaimPermission)
		roleClaimPermission.POST("/", roleClaimPermissionController.UpdateRoleClaimPermission)

		roleClaimPermission.DELETE("/:id", roleClaimPermissionController.DeleteRoleClaimPermission)
	}
}
