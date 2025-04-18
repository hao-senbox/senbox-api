package usecase

import (
	"sen-global-api/internal/data/repository"
	"sen-global-api/internal/domain/request"
)

type CreateRoleClaimPermissionUseCase struct {
	*repository.RoleClaimPermissionRepository
}

func (receiver *CreateRoleClaimPermissionUseCase) Create(req request.CreateRoleClaimPermissionRequest) error {
	return receiver.CreateRoleClaimPermission(req)
}
