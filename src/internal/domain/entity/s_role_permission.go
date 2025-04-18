package entity

import (
	"time"
)

type SRoleClaimPermission struct {
	ID             int64      `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	PermissionName string     `gorm:"type:varchar(255);not null;"`
	Description    string     `gorm:"type:varchar(255);not null;default:''"`
	RoleClaimId    int64      `gorm:"column:role_claim_id;"`
	RoleClaim      SRoleClaim `gorm:"foreignKey:RoleClaimId;references:id;constraint:OnDelete:CASCADE"`
	CreatedAt      time.Time  `gorm:"default:CURRENT_TIMESTAMP;not null"`
	UpdatedAt      time.Time  `gorm:"default:CURRENT_TIMESTAMP;not null"`
}
