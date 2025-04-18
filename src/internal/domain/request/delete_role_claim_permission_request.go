package request

type DeleteRoleClaimPermissionRequest struct {
	ID uint `json:"id" binding:"required"`
}
