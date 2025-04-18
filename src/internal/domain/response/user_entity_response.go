package response

type UserEntityResponse struct {
	ID           string   `json:"id"`
	Username     string   `json:"username"`
	Fullname     string   `json:"fullname"`
	Phone        string   `json:"phone"`
	Email        string   `json:"email"`
	Dob          string   `json:"dob"`
	Organization []string `json:"organizations"`
	CreatedAt    string   `json:"created_at"`

	Roles     *[]RoleListResponseData   `json:"roles"`
	Guardians *[]UserEntityResponseData `json:"guardians"`
	Devices   *[]string                 `json:"devices"`
}

type UserEntityResponseData struct {
	ID       string   `json:"id"`
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
}

type UserEntityDataResponse struct {
	Data []UserEntityResponseData `json:"data"`
}
