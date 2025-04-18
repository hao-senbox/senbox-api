package repository

import (
	"errors"
	"sen-global-api/internal/domain/entity"
	"sen-global-api/internal/domain/request"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type OrganizationRepository struct {
	DBConn *gorm.DB
}

func NewOrganizationRepository(dbConn *gorm.DB) *OrganizationRepository {
	return &OrganizationRepository{DBConn: dbConn}
}

func (receiver *OrganizationRepository) GetByID(id uint) (*entity.SOrganization, error) {
	var organization entity.SOrganization
	err := receiver.DBConn.Where("id = ?", id).First(&organization).Error
	if err != nil {
		log.Error("OrganizationRepository.GetByID: " + err.Error())
		return nil, errors.New("failed to get organization")
	}
	return &organization, nil
}

func (receiver *OrganizationRepository) CreateOrganization(req request.CreateOrganizationRequest) error {
	result := receiver.DBConn.Create(&entity.SOrganization{
		OrganizationName: req.OrganizationName,
		Address:     req.Address,
		Description: req.Description,
	})

	if result.Error != nil {
		log.Error("OrganizationRepository.CreateOrganization: " + result.Error.Error())
		return errors.New("failed to create organization")
	}

	return nil
}

func (receiver *OrganizationRepository) UpdateOrganization(req request.UpdateOrganizationRequest) error {
	updateResult := receiver.DBConn.Model(&entity.SOrganization{}).Where("id = ?", req.ID).
		Updates(map[string]interface{}{
			"organization_name": req.OrganizationName,
			"address":      req.Address,
			"description":  req.Description,
		})

	if updateResult.Error != nil {
		log.Error("OrganizationRepository.UpdateOrganization: " + updateResult.Error.Error())
		return errors.New("failed to update organization")
	}

	return nil
}
