package entity

import (
	"github.com/google/uuid"
)

type SUserRoles struct {
	UserId uuid.UUID   `gorm:"column:user_id;primary_key"`
	User   SUserEntity `gorm:"foreignKey:UserId;references:id;constraint:OnDelete:CASCADE;"`
	RoleId int64       `gorm:"column:role_id;primary_key"`
	Role   SRole       `gorm:"foreignKey:RoleId;references:id;constraint:OnDelete:CASCADE"`
	// RoleClaimId           int64                `gorm:"column:role_claim_id;primary_key"`
	// RoleClaim             SRoleClaim           `gorm:"foreignKey:RoleClaimId;references:id;constraint:OnDelete:CASCADE"`
	// RoleClaimPermissionId int64                `gorm:"column:role_claim_permission_id;primary_key"`
	// RoleClaimPermission   SRoleClaimPermission `gorm:"foreignKey:RoleClaimPermissionId;references:id;constraint:OnDelete:CASCADE"`
}
