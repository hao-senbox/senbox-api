package response

type RoleClaimPermissionResponse struct {
	ID             int64  `json:"id"`
	PermissionName string `json:"permission_name"`
	Description    string `json:"description"`
}

type RoleClaimPermissionListResponseData struct {
	ID             int64  `json:"id"`
	PermissionName string `json:"permission_name"`
}

type RoleClaimPermissionListResponse struct {
	Data []RoleClaimPermissionListResponseData `json:"data"`
}
