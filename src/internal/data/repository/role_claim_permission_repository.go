package repository

import (
	"errors"
	"sen-global-api/internal/domain/entity"
	"sen-global-api/internal/domain/request"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type RoleClaimPermissionRepository struct {
	DBConn *gorm.DB
}

func NewRoleClaimPermissionRepository(dbConn *gorm.DB) *RoleClaimPermissionRepository {
	return &RoleClaimPermissionRepository{DBConn: dbConn}
}

func (receiver *RoleClaimPermissionRepository) GetAll() ([]entity.SRoleClaimPermission, error) {
	var policies []entity.SRoleClaimPermission
	err := receiver.DBConn.Table("s_role_permission").Find(&policies).Error
	if err != nil {
		log.Error("RoleClaimPermissionRepository.GetAll: " + err.Error())
		return nil, errors.New("failed to get all policies")
	}

	return policies, err
}

func (receiver *RoleClaimPermissionRepository) GetByID(req request.GetRoleClaimPermissionByIdRequest) (*entity.SRoleClaimPermission, error) {
	var permission entity.SRoleClaimPermission
	err := receiver.DBConn.Where("id = ?", req.ID).First(&permission).Error
	if err != nil {
		log.Error("RoleClaimPermissionRepository.GetByID: " + err.Error())
		return nil, errors.New("failed to get permission")
	}
	return &permission, nil
}

func (receiver *RoleClaimPermissionRepository) GetByName(req request.GetRoleClaimPermissionByNameRequest) (*entity.SRoleClaimPermission, error) {
	var permission entity.SRoleClaimPermission
	err := receiver.DBConn.Where("permission_name = ?", req.PermissionName).First(&permission).Error
	if err != nil {
		log.Error("RoleClaimPermissionRepository.GetByName: " + err.Error())
		return nil, errors.New("failed to get permission")
	}
	return &permission, nil
}

func (receiver *RoleClaimPermissionRepository) CreateRoleClaimPermission(req request.CreateRoleClaimPermissionRequest) error {
	permission, _ := receiver.GetByName(request.GetRoleClaimPermissionByNameRequest{PermissionName: req.PermissionName})

	if permission != nil {
		log.Error("RoleClaimPermissionRepository.CreateRoleClaimPermission: " + permission.PermissionName)
		return errors.New("permission already existed")
	}

	var roleClaimCount int64
	receiver.DBConn.Model(&entity.SRoleClaim{}).Where("id = ?", req.RoleClaimId).Count(&roleClaimCount)

	if roleClaimCount == 0 {
		log.Error("RoleClaimPermissionRepository.CreateRoleClaimPermission: " + "role claim not found")
		return errors.New("role claim not found")
	}

	permissionReq := entity.SRoleClaimPermission{
		PermissionName: req.PermissionName,
		Description:    req.Description,
		RoleClaimId:    req.RoleClaimId,
	}
	permissionResult := receiver.DBConn.Create(&permissionReq)

	if permissionResult.Error != nil {
		log.Error("RoleClaimPermissionRepository.CreateRoleClaimPermission: " + permissionResult.Error.Error())
		return errors.New("failed to create permission")
	}

	return nil
}

func (receiver *RoleClaimPermissionRepository) UpdateRoleClaimPermission(req request.UpdateRoleClaimPermissionRequest) error {
	updateResult := receiver.DBConn.Model(&entity.SRoleClaimPermission{}).Where("id = ?", req.ID).
		Updates(map[string]interface{}{
			"permission_name": req.PermissionName,
			"description":     req.Description,
			"role_claim_id":   req.RoleClaimId,
		})

	if updateResult.Error != nil {
		log.Error("RoleClaimPermissionRepository.UpdateRoleClaimPermission: " + updateResult.Error.Error())
		return errors.New("failed to update permission")
	}

	return nil
}

func (receiver *RoleClaimPermissionRepository) DeleteRoleClaimPermission(req request.DeleteRoleClaimPermissionRequest) error {
	result := receiver.DBConn.Delete(&entity.SRoleClaimPermission{}, req.ID)
	if result.Error != nil {
		log.Error("RoleClaimPermissionRepository.DeleteRoleClaimPermission: " + result.Error.Error())
		return errors.New("failed to delete permission")
	}
	return nil
}
