package usecase

import (
	"sen-global-api/internal/data/repository"
	"sen-global-api/internal/domain/entity"
	"sen-global-api/internal/domain/request"
)

type GetRoleClaimPermissionUseCase struct {
	*repository.RoleClaimPermissionRepository
}

func (receiver *GetRoleClaimPermissionUseCase) GetAllRoleClaimPermission() ([]entity.SRoleClaimPermission, error) {
	return receiver.GetAll()
}

func (receiver *GetRoleClaimPermissionUseCase) GetRoleClaimPermissionById(req request.GetRoleClaimPermissionByIdRequest) (*entity.SRoleClaimPermission, error) {
	return receiver.GetByID(req)
}

func (receiver *GetRoleClaimPermissionUseCase) GetRoleClaimPermissionByName(req request.GetRoleClaimPermissionByNameRequest) (*entity.SRoleClaimPermission, error) {
	return receiver.GetByName(req)
}
