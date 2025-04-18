package request

type UpdateRoleClaimPermissionRequest struct {
	ID             uint   `json:"id" binding:"required"`
	PermissionName string `json:"permission_name" binding:"required"`
	Description    string `json:"desciption" binding:"required"`
	RoleClaimId    int64  `json:"role_claim_id" binding:"required"`
}
