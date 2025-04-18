package usecase

import (
	"sen-global-api/internal/data/repository"
	"sen-global-api/internal/domain/request"
)

type UpdateRoleClaimPermissionUseCase struct {
	*repository.RoleClaimPermissionRepository
}

func (receiver *UpdateRoleClaimPermissionUseCase) UpdateRoleClaimPermission(req request.UpdateRoleClaimPermissionRequest) error {
	return receiver.RoleClaimPermissionRepository.UpdateRoleClaimPermission(req)
}
