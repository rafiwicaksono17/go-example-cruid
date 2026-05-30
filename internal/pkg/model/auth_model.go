package model

type (
	// Auth Models
	RegisterRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=6"`
		Name     string `json:"name" validate:"required"`
		Phone    string `json:"phone" validate:"required"`
	}

	LoginRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	AuthResponse struct {
		Token string       `json:"token"`
		User  UserResponse `json:"user"`
	}

	// User Models
	UserResponse struct {
		ID      uint   `json:"id"`
		Email   string `json:"email"`
		Phone   string `json:"phone"`
		Name    string `json:"name"`
		IsAdmin bool   `json:"is_admin"`
	}

	UpdateUserRequest struct {
		Name  string `json:"name"`
		Phone string `json:"phone"`
		Email string `json:"email"`
	}

	GetUsersResponse struct {
		Data  []UserResponse `json:"data"`
		Total int64          `json:"total"`
	}
)
