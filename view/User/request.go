package user

type InsertUserRequest struct {
	Nama     string `json:"nama" `
	Email    string `json:"email" `
	Password string `json:"password" `
}

type UpdateUserRequest struct {
	Nama     string `json:"nama"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
}
