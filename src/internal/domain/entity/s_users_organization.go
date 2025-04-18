package entity

import (
	"github.com/google/uuid"
)

type SUsersOrganization struct {
	UserId         uuid.UUID     `gorm:"column:user_id;primary_key"`
	User           SUserEntity   `gorm:"foreignKey:UserId;references:id;constraint:OnDelete:CASCADE;"`
	OrganizationID int64         `gorm:"column:organization_id;primary_key"`
	Organization   SOrganization `gorm:"foreignKey:OrganizationID;references:id;constraint:OnDelete:CASCADE"`
}
