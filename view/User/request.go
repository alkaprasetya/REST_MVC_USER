package user

type InsertUserRequest struct {
	Nama     string `json:"nama" `
	Email    string `json:"email" `
	Password string `json:"password" `
}

type UpdateUserRequest struct {
	Nama     string `json:"nama" validate:"required"`
	Email    string `json:"Email" validate:"required"`
	Password string `json:"Password" validate:"required"`
}
