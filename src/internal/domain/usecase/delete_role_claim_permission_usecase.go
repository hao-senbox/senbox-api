package usecase

import (
	"sen-global-api/internal/data/repository"
	"sen-global-api/internal/domain/request"
)

type DeleteRoleClaimPermissionUseCase struct {
	*repository.RoleClaimPermissionRepository
}

func (receiver *DeleteRoleClaimPermissionUseCase) DeleteRoleClaimPermission(req request.DeleteRoleClaimPermissionRequest) error {
	return receiver.RoleClaimPermissionRepository.DeleteRoleClaimPermission(req)
}
