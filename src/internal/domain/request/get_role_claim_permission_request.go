package request

type GetRoleClaimPermissionByIdRequest struct {
	ID uint `json:"id" binding:"required"`
}

type GetRoleClaimPermissionByNameRequest struct {
	PermissionName string `json:"permission_name" binding:"required"`
}
