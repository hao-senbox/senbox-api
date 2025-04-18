package request

type CreateRoleRequest struct {
	RoleName       string `json:"role_name" binding:"required"`
	Description    string `json:"description" default:"" binding:"required"`
	OrganizationId int64  `json:"organization_id" binding:"required"`
}
