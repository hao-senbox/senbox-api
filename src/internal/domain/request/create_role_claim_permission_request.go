package request

type CreateRoleClaimPermissionRequest struct {
	PermissionName string `json:"permission_name" binding:"required"`
	Description    string `json:"desciption" default:"" binding:"required"`
	RoleClaimId    int64  `json:"role_claim_id" binding:"required"`
}
