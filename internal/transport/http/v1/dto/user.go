package httpdto

type CreateUserDTO struct {
	Username       string `json:"username,omitempty"`
	Email          string `json:"email,omitempty"`
	Password       string `json:"password,omitempty"`
	RepeatPassword string `json:"repeatpassword,omitempty"`
}

type AuthDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
