package types

type CreateUserRequest struct {
	Username 		string `json:"username"`
	Email 			string `json:"email"`
}

type UserRequestHeader struct {
	ID 					int `json:"id"`
	Username 		string `json:"username"`
	Email 			string `json:"email"`
	ImageURL 		string `json:"image_url"`
	// Role 				string `json:"role"`
	ExpairAt 		int64 `json:"expair_at"`
}

type User struct {
	ID 					int `json:"id"`
	Username 		string `json:"username"`
	Email 			string `json:"email"`
	ImageURL 		string `json:"image_url"`
	CreatedAt 	string `json:"created_at"`
	UpdatedAt 	string `json:"updated_at"`
}
