package user

type User struct {
	ID           string `json:"id" bson:"_id,omitempty"`
	Username     string `json:"username" bson:"username,omitempty"`
	Email        string `json:"email" bson:"email,omitempty"`
	PasswordHash string `json:"-" bson:"password"`
}

type CreateUserDTO struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
