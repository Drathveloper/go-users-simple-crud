package dto

type GetUserResponse struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	BirthDate string `json:"birth_date"`
}

type RegisterUserRequest struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	BirthDate string `json:"birth_date"`
}

type RegisterUserResponse struct {
	ID int64 `json:"id"`
}

type UpdateUserRequest struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	BirthDate string `json:"birth_date"`
}

type CountResponse struct {
	Message string
}
