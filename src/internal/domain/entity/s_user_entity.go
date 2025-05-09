package entity

import (
	"html"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type SUserEntity struct {
	ID        uuid.UUID `gorm:"type:char(36);primary_key"`
	Username  string    `gorm:"type:varchar(255);not null;default:''"`
	Fullname  string    `gorm:"type:varchar(255);not null;default:''"`
	Phone     string    `gorm:"type:varchar(255);not null;default:''"`
	Email     string    `gorm:"type:varchar(255);not null;default:''"`
	Birthday  time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Password  string    `gorm:"type:varchar(255);not null;default:''"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP;not null"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP;not null"`

	// Many-to-many relationship with roles
	Roles []SRole `gorm:"many2many:s_user_roles;foreignKey:id;joinForeignKey:user_id;references:id;joinReferences:role_id"`

	Organizations []SOrganization `gorm:"many2many:s_user_organizations;foreignKey:id;joinForeignKey:user_id;references:id;joinReferences:organization_id"`

	Guardians []SUserEntity `gorm:"many2many:s_user_guardians;foreignKey:id;joinForeignKey:user_id;references:id;joinReferences:guardian_id"`
	Devices   []SDevice     `gorm:"many2many:s_user_devices;foreignKey:id;joinForeignKey:user_id;references:id;joinReferences:device_id"`
}

func (user *SUserEntity) BeforeCreate(tx *gorm.DB) (err error) {
	id, err := uuid.NewUUID()
	if err == nil {
		user.ID = id
	}

	if user.Password != "" {
		encryptedPwdData, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err == nil {
			user.Password = string(encryptedPwdData)
		}
	}

	user.Username = strings.ToLower(html.EscapeString(strings.TrimSpace(user.Username)))

	return err
}
