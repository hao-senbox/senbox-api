package usecase

import (
	"sen-global-api/internal/data/repository"
	"sen-global-api/internal/domain/entity"
	"sen-global-api/internal/domain/request"
)

type GetRoleUseCase struct {
	*repository.RoleRepository
}

func (receiver *GetRoleUseCase) GetAllRoleByOrganization(organizationId int64) ([]entity.SRole, error) {
	return receiver.GetAllByOrganization(organizationId)
}

func (receiver *GetRoleUseCase) GetRoleById(req request.GetRoleByIdRequest) (*entity.SRole, error) {
	return receiver.GetByID(req)
}

func (receiver *GetRoleUseCase) GetRoleByName(req request.GetRoleByNameRequest) (*entity.SRole, error) {
	return receiver.GetByName(req)
}
